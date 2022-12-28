using service;
IHost host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((context, services) =>
    {
        services.AddHostedService<Worker>();
        services.AddHttpClient("Sonnen", httpClient =>
        {
            Uri baseUri = new Uri(context.Configuration["Sonnen:BaseUrl"] ?? "http://localhost");
            httpClient.BaseAddress = new Uri(baseUri, "api/v2/");
        });
        services.AddHttpClient("ChargeHq", httpClient =>
        {
            var baseUri = new Uri(context.Configuration["ChargeHq:BaseUrl"]?.ToString());
            httpClient.BaseAddress = new Uri(baseUri, "api/public/");
        });

    }).Build();

host.Run();