using service;
IHost host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((context, services) =>
    {
        services.AddHostedService<Worker>();
        services.AddHttpClient("Sonnen", httpClient =>
        {
            var url = context.Configuration["Sonnen:BaseUrl"]?.ToString();
            var apiKey = context.Configuration["Sonnen:ApiKey"]?.ToString();
            httpClient.BaseAddress = new Uri($"{url}api/v2/");
        });
        services.AddHttpClient("ChargeHq", httpClient =>
        {
            var url = context.Configuration["ChargeHq:BaseUrl"]?.ToString();
            httpClient.BaseAddress = new Uri(url);
        });

    }).Build();

host.Run();