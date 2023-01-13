namespace service;
public struct status

{
    public int Apparent_output { get; set; }// 4597,

    public string BackupBuffer { get; set; }// "0",

    public bool BatteryCharging { get; set; }// false,

    public bool BatteryDischarging { get; set; }// false,

    public int Consumption_Avg { get; set; }// 3077,

    public int Consumption_W { get; set; }// 3096,

    public decimal Fac { get; set; }// 49.99100112915039,

    public bool FlowConsumptionBattery { get; set; }// false,

    public bool FlowConsumptionGrid { get; set; }// false,

    public bool FlowConsumptionProduction { get; set; }// true,

    public bool FlowGridBattery { get; set; }// false,

    public bool FlowProductionBattery { get; set; }// false,

    public bool FlowProductionGrid { get; set; }// true,

    public int GridFeedIn_W { get; set; }// 1506,

    public int IsSystemInstalled { get; set; }// 1,

    public string OperatingMode { get; set; }// "2",

    public int Pac_total_W { get; set; }// -21,

    public int Production_W { get; set; }// 4636,

    public int RSOC { get; set; }// 100,

    public int RemainingCapacity_Wh { get; set; }// 5112,

    public int? Sac1 { get; set; }// 4597,

    public int? Sac2 { get; set; }// null,

    public int? Sac3 { get; set; }// null,

    public string SystemStatus { get; set; }// "OnGrid",

    public string Timestamp { get; set; }// "2022-12-27 03 {get; set;}//37 {get; set;}//46",

    public int USOC { get; set; }// 100,

    public int Uac { get; set; }// 244,

    public int Ubat { get; set; }// 51,

    public bool dischargeNotAllowed { get; set; }// false,

    public bool generator_autostart { get; set; }// false


}