package main

type SonnenStatus struct {
	Apparent_output           int
	BackupBuffer              string
	BatteryCharging           bool
	BatteryDischarging        bool
	Consumption_Avg           int
	Consumption_W             int
	Fac                       float32
	FlowConsumptionBattery    bool
	FlowConsumptionGrid       bool
	FlowConsumptionProduction bool
	FlowGridBattery           bool
	FlowProductionBattery     bool
	FlowProductionGrid        bool
	GridFeedIn_W              int
	IsSystemInstalled         int
	OperatingMode             string
	Pac_total_W               int
	Production_W              int
	RSOC                      int
	RemainingCapacity_Wh      int
	Sac1                      int
	Sac2                      int
	Sac3                      int
	SystemStatus              string
	Timestamp                 string
	USOC                      int
	Uac                       int
	Ubat                      int
	dischargeNotAllowed       bool
	generator_autostart       bool
}

type ChargeHq struct {
	apiKey string
	// [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	siteMeters SiteMeters

	// timestamp of meter data (milliseconds since epoch)
	// if the meter data is delayed and has a reliable timestamp then this field should
	// be provided if, otherwise it should be left unset
	// [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	tsms float32

	// set this field only if there was an error obtaining the meter data
	// [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	error string
}
type SiteMeters struct {
	// if a consumption meter is present, the following fields should be set
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	consumption_kw float32
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	net_import_kw float32 //grid import, negative means export

	// if solar is present, provide the following field
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	production_kw float32

	// if accumulated import/export energy is available, set the following fields

	// [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	exported_kwh float32
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	imported_kwh float32
	// if a battery is present, provide the following fields
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	battery_discharge_kw float32 //negative mean charging
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	battery_soc float32 //eg 0.5 = 50%
	//[JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
	battery_energy_kwh float32 //amount of energy in the battery (optional)
}
type httpError struct {
	method     string
	url        string
	status     string
	statusCode int
}
