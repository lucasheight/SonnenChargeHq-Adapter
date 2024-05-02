package main

type SonnenStatus struct {
	StatusCode                float32
	Status                    string
	Apparent_output           float32
	BackupBuffer              string
	BatteryCharging           bool
	BatteryDischarging        bool
	Consumption_Avg           float32
	Consumption_W             float32
	Fac                       float32
	FlowConsumptionBattery    bool
	FlowConsumptionGrid       bool
	FlowConsumptionProduction bool
	FlowGridBattery           bool
	FlowProductionBattery     bool
	FlowProductionGrid        bool
	GridFeedIn_W              float32
	IsSystemInstalled         float32
	OperatingMode             string
	Pac_total_W               float32
	Production_W              float32
	RSOC                      float32
	RemainingCapacity_Wh      float32
	Sac1                      float32
	Sac2                      float32
	Sac3                      float32
	SystemStatus              string
	Timestamp                 string
	USOC                      float32
	Uac                       float32
	Ubat                      float32
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
