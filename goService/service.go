package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
)

var (
	info     *log.Logger
	warn     *log.Logger
	err      *log.Logger
	useOsEnv bool   = false
	apiKey   string = ""
)

func init() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	err = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	apiKey = getEnv(ChargeHqApiKey)
}
func main() {
	info.Printf("%s has started \n", ServiceName)
	poll()
}
func poll() {
	ms, err := strconv.Atoi(getEnv(ChargeHqRefreshMs))
	if err != nil {
		ms = ChargeHqDefaultRefreshMs
	}
	worker()
	interval := time.Tick(time.Duration(ms) * time.Millisecond)
	for range interval {
		worker()
	}
}
func worker() {
	read := make(chan SonnenStatus)
	go get(read, "/api/v2/status/")
	select {
	case res := <-read:
		if res.StatusCode != 200 {
			warn.Printf("GET-%d: %+v \n", res.StatusCode, res.Status)
		} else {
			info.Printf("Sonnen data read: %+v \n", res)
			//now post it
			go post(res)
		}
	case <-time.After(time.Second * 5):
		info.Println("GET timeout received.")
	}
}
func post(data SonnenStatus) {
	mapped := func() ChargeHq {
		ch := ChargeHq{}
		ch.apiKey = apiKey
		ch.siteMeters = SiteMeters{
			consumption_kw:       float32(data.Consumption_W) / 1000,
			production_kw:        float32(data.Production_W) / 1000,
			net_import_kw:        float32(-data.GridFeedIn_W) / 1000,
			battery_soc:          float32(data.USOC) / 100,
			battery_discharge_kw: float32(data.Pac_total_W) / 1000}
		return ch
	}()
	charge := map[string]interface{}{
		"apiKey": mapped.apiKey
	}
	meters := map[string]interface{}{
		"consumption_kw":       mapped.siteMeters.consumption_kw,
		"production_kw":        mapped.siteMeters.production_kw,
		"net_import_kw":        mapped.siteMeters.net_import_kw,
		"battery_soc":          mapped.siteMeters.battery_soc,
		"battery_discharge_kw": mapped.siteMeters.battery_discharge_kw
	}

	charge["siteMeters"] = meters
	var postBuffer bytes.Buffer
	er := json.NewEncoder(&postBuffer).Encode(charge)
	if er != nil {
		warn.Printf("POST-Json encoding: %s \n", er.Error())
		return
	}
	h := http.Client{}
	var endpointUrl = ChargeHqBaseUrl + "/api/public/push-solar-data"
	resp, e := h.Post(endpointUrl, "application/json", &postBuffer)
	if e != nil {
		warn.Printf("POST-Endpoint: %s \n", er.Error())
		return
	}
	if resp.StatusCode != 200 {
		warn.Printf("POST-%d %s. : %s \n", resp.StatusCode, resp.Status, endpointUrl)
		return
	}
	defer resp.Body.Close()
	info.Printf("ChargeHq data sent: %+v \n", mapped)
}
func get(read chan<- SonnenStatus, url string) {
	sonnenClient := http.Client{}
	var req *http.Response
	var err error
	var theUrl = getEnv(SonnenBaseUrl) + url 
	req, err = sonnenClient.Get(theUrl)
	sonnenData := SonnenStatus{}
	if err != nil {
		sonnenData.StatusCode = 500
		sonnenData.Status = fmt.Sprintf("Server Error: %s \n", err.Error())
		read <- sonnenData
		return
	}
	sonnenData.StatusCode = req.StatusCode
	sonnenData.Status = req.Status
	defer req.Body.Close()
	dec := json.NewDecoder(req.Body)
	decError := dec.Decode(&sonnenData)
	if decError != nil {
		warn.Printf("Serialisation Error: %s \n", decError.Error())
	}
	read <- sonnenData
}

func getEnv(key string) string {
	if !useOsEnv {
		e := godotenv.Load(".env")
		if e != nil {
			warn.Println("Unable to load .env file, will try OS environment.")
			useOsEnv = true
		}
	}
	val, hasVal := os.LookupEnv(key)
	if !hasVal {
		defer err.Fatal(MissingEnv)
	}
	return val
}
