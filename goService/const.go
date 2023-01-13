package main

import "fmt"

const ServiceName string = "sonnen-chargehq-service"
const SonnenBaseUrl string = "Sonnen__BaseUrl"
const ChargeHqApiKey string = "ChargeHq__ApiKey"
const ChargeHqRefreshMs string = "ChargeHq__RefreshMs" //optional

//defaults
const ChargeHqBaseUrl string = "https://api.chargehq.net"
const ChargeHqDefaultRefreshMs = 120000

var MissingEnv string = fmt.Sprintf(`This service requires the following environment settings to be set:
%s, %s`, SonnenBaseUrl, ChargeHqApiKey)
