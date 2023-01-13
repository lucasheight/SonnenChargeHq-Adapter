package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
)

func init() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	err = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
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
	for _ = range interval {
		worker()
	}
}
func worker() int {
	sonnenData, e, statusCode := readSonnen()
	if e != nil {
		err.Fatal(err)

	}
	if statusCode != 200 {
		warn.Printf("%d: %+v \n", statusCode, sonnenData)

	}
	//info.Printf("%+v \n", sonnenData)
	publishData(sonnenData, e, statusCode)
	return statusCode
}
func publishData(data SonnenStatus, err error, statusCode int) error {
	mapped := mapData(data, err)
	charge := map[string]interface{}{
		"apiKey": mapped.apiKey}
	if err != nil {
		charge["error"] = err.Error()
	} else {
		meters := map[string]interface{}{
			"consumption_kw":       mapped.siteMeters.consumption_kw,
			"production_kw":        mapped.siteMeters.production_kw,
			"net_import_kw":        mapped.siteMeters.net_import_kw,
			"battery_soc":          mapped.siteMeters.battery_soc,
			"battery_discharge_kw": mapped.siteMeters.battery_discharge_kw}
		charge["siteMeters"] = meters
	}
	var postBuffer bytes.Buffer
	err = json.NewEncoder(&postBuffer).Encode(&charge)
	if err != nil {
		return err
	}
	h := &http.Client{}
	resp, e := h.Post(ChargeHqBaseUrl+"/api/public/push-solar-data", "application/json", &postBuffer)

	if e != nil {
		return e
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}

	info.Printf("ChargeHq data sent: %+v \n", mapped)
	return nil

}
func readSonnen() (SonnenStatus, error, int) {
	sonnenClient := &http.Client{}
	var req *http.Response
	var err error
	req, err = sonnenClient.Get(getEnv(SonnenBaseUrl) + "/api/v2/status/")
	//req.Body = http.MaxBytesReader(w, req.Body, 1048576)
	if err != nil {
		return SonnenStatus{}, err, req.StatusCode
	}

	defer req.Body.Close()
	//body, err := io.ReadAll(req.Body)
	sonnenData := new(SonnenStatus)
	dec := json.NewDecoder(req.Body)
	err = dec.Decode(&sonnenData)
	if err != nil {
		req.StatusCode = 500
	}
	info.Printf("Sonnen data read: %+v \n", *sonnenData)

	return *sonnenData, err, req.StatusCode
}
func mapData(data SonnenStatus, err error) ChargeHq {

	ch := new(ChargeHq)
	ch.apiKey = getEnv(ChargeHqApiKey)
	if err != nil {
		ch.error = err.Error()
	}
	ch.siteMeters = SiteMeters{
		consumption_kw:       float32(data.Consumption_W) / 1000,
		production_kw:        float32(data.Production_W) / 1000,
		net_import_kw:        float32(-data.GridFeedIn_W) / 1000,
		battery_soc:          float32(data.USOC) / 100,
		battery_discharge_kw: float32(data.Pac_total_W) / 1000}
	return *ch

}
func getEnv(key string) string {
	e := godotenv.Load(".env")
	if e != nil {
		warn.Println("Unable to load .env file, will try OS environment.")
	}
	val, hasVal := os.LookupEnv(key)
	if !hasVal {
		defer err.Fatal(MissingEnv)
		//log.Fatalf("Unable to find %s environment variable. Terminating process!.", key)
	}

	return val
}
