namespace service;
using System.Text.Json;
using static System.Net.Mime.MediaTypeNames;
public class Worker : BackgroundService
{
    private readonly ILogger<Worker> _logger;
    private readonly IConfiguration _config;
    private readonly IHttpClientFactory _httpClientFactory;
    private readonly int _refreshMs = 60000;

    public Worker(ILogger<Worker> logger, IConfiguration configuration, IHttpClientFactory httpClientFactory)
    {
        _logger = logger;
        _config = configuration;
        _httpClientFactory = httpClientFactory;
        if (int.TryParse(configuration["ChargeHq:RefreshMs"], out int result))
        {
            _refreshMs = result;
        };
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {

        _logger.LogInformation("sonnen-chargehq-service started at: {time}", DateTimeOffset.Now);
        while (!stoppingToken.IsCancellationRequested)
        {

            var client = _httpClientFactory.CreateClient("Sonnen");
            var message = await client.GetAsync("status");
            var output = new chargeHq() { apiKey = _config["ChargeHq:ApiKey"].ToString() };
            if (message.IsSuccessStatusCode)
            {
                var res = await message.Content.ReadAsStringAsync();

                using var content = await message.Content.ReadAsStreamAsync();
                status stat = await JsonSerializer.DeserializeAsync<status>(content);
                output.siteMeters = new siteMeters
                {
                    consumption_kw = Convert.ToDecimal(stat.Consumption_W) / 1000,
                    production_kw = Convert.ToDecimal(stat.Production_W) / 1000,
                    net_import_kw = Convert.ToDecimal(-stat.GridFeedIn_W) / 1000,
                    battery_soc = Convert.ToDecimal(stat.RSOC) / 100,
                    battery_discharge_kw = Convert.ToDecimal(stat.Pac_total_W) / 1000
                };


            }
            else
            {
                output.error = $"Solar source data feed error: {message.ReasonPhrase}";
                _logger.LogError(message.ReasonPhrase);
            }
            var poster = _httpClientFactory.CreateClient("ChargeHq");
            var postContent = new StringContent(JsonSerializer.Serialize(output), System.Text.Encoding.UTF8, Application.Json);
            using var postResult = await poster.PostAsync("push-solar-data", postContent);
            try
            {
                _logger.LogInformation(JsonSerializer.Serialize(output).ToString());
                postResult.EnsureSuccessStatusCode();
            }
            catch (System.Exception ex)
            {

                _logger.LogError(ex.Message);
            }



            await Task.Delay(_refreshMs, stoppingToken);
        }
    }
    public override async Task StopAsync(CancellationToken cancellationToken)
    {
        _logger.LogInformation("sonnen-chargehq-service is stopping: {time}", DateTimeOffset.Now);
        await base.StopAsync(cancellationToken);
    }

}
