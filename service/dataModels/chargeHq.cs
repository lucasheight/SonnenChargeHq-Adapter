namespace service;
using System.Text.Json.Serialization;
public struct chargeHq
{
    public string apiKey { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public siteMeters siteMeters { get; set; }

    // timestamp of meter data (milliseconds since epoch)
    // if the meter data is delayed and has a reliable timestamp then this field should
    // be provided if, otherwise it should be left unset
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? tsms { get; set; }

    // set this field only if there was an error obtaining the meter data
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public string? error { get; set; }


}
public struct siteMeters
{
    // if a consumption meter is present, the following fields should be set
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? consumption_kw { get; set; } //total site consumption
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? net_import_kw { get; set; } //grid import, negative means export


    // if solar is present, provide the following field
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal production_kw { get; set; }

    // if accumulated import/export energy is available, set the following fields

    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? exported_kwh { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? imported_kwh { get; set; }
    // if a battery is present, provide the following fields
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? battery_discharge_kw { get; set; }//negative mean charging
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? battery_soc { get; set; } //eg 0.5 = 50%
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingDefault)]
    public decimal? battery_energy_kwh { get; set; } //amount of energy in the battery (optional)
}
