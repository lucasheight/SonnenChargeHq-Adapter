package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	info     *log.Logger
	warn     *log.Logger
	err      *log.Logger
	useOsEnv bool = false
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
	go worker()
	interval := time.Tick(time.Duration(ms) * time.Millisecond)
	for _ = range interval {
		go worker()
	}
}
func worker() {
	sonnenData, e, statusCode := readSonnen()
	if e != nil {
		err.Fatalf("%s \n", e.Error())
	}
	if statusCode != 200 {
		warn.Printf("%d: %+v \n", statusCode, sonnenData)
	}
	var er, httpEr = publishData(sonnenData, e, statusCode)
	if er != nil {
		err.Fatalf("%s \n", er.Error())
	}
	if httpEr != nil {
		warn.Printf("%s; %d: %s => %s", httpEr.method, httpEr.statusCode, httpEr.status, httpEr.url)
	}
}
func publishData(data SonnenStatus, er error, statusCode int) (error, *httpError) {
	mapped := mapData(data, er)
	charge := map[string]interface{}{
		"apiKey": mapped.apiKey}
	if er != nil {
		charge["error"] = er.Error()
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
	er = json.NewEncoder(&postBuffer).Encode(&charge)
	if err != nil {
		return er, nil
	}
	h := &http.Client{}
	var endpointUrl = ChargeHqBaseUrl + "/api/public/push-solar-data"
	resp, e := h.Post(endpointUrl, "application/json", &postBuffer)
	if e != nil {
		return e, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, &httpError{method: "POST", url: endpointUrl, status: resp.Status, statusCode: resp.StatusCode} // httperr() errors.New(fmt.Sprintf("StatusCode:%d, Status:%s => %s", resp.StatusCode, resp.Status, endpointUrl))
	}
	info.Printf("ChargeHq data sent: %+v \n", mapped)
	return nil, nil
}
func readSonnen() (SonnenStatus, error, int) {
	sonnenClient := &http.Client{}
	var req *http.Response
	var err error
	var theUrl = getEnv(SonnenBaseUrl) + "/api/v2/status/"
	req, err = sonnenClient.Get(theUrl)

	if err != nil {
		return SonnenStatus{}, err, req.StatusCode
	}
	defer req.Body.Close()
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
