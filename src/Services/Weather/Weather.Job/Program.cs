using System;
using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using System.Net.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.AspNetCore.Builder;
using RestEase;
using Microsoft.AspNetCore.Diagnostics.HealthChecks;
using HealthChecks.UI.Client;
using Microsoft.AspNetCore.Http;
using Weather.Infrastructure;
using MongoDB.Bson;
using MongoDB.Driver;

namespace Weather.Job
{
    class Program
    {
        static void Main(string[] args)
        {
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.Configure((context, app) =>
                    {
                        if (context.HostingEnvironment.IsDevelopment())
                        {
                            app.UseDeveloperExceptionPage();
                        }
                        app.UseRouting();
                        app.UseEndpoints(endpoints =>
                        {
                            endpoints.MapHealthChecks("/hc", new HealthCheckOptions()
                            {
                                Predicate = _ => true,
                                ResponseWriter = UIResponseWriter.WriteHealthCheckUIResponse
                            });
                            endpoints.MapGet("/", async context =>
                            {
                                
                                //await context.Response.WriteAsync("Hello World!");
                            });
                        });
                    });
                })
                .ConfigureAppConfiguration((context, config) =>
                {
                    config.Sources.Clear();

                    var env = context.HostingEnvironment;
                    config.AddJsonFile("appsettings.json", optional: true, reloadOnChange: true)
                          .AddJsonFile($"appsettings.{env.EnvironmentName}.json",
                                         optional: true, reloadOnChange: true);
                    config.AddEnvironmentVariables();
                })
                .ConfigureServices(services =>
                {
                    services.AddHostedService<CurrentWeatherService>();
                    services.AddSingleton<CurrentWeatherServiceOptions>();
                    services.AddTransient(prov =>
                    {
                        var conf = prov.GetRequiredService<IConfiguration>();
                        var options = new MongoDbOptions
                        {
                            ConnectionString = conf["MongoDb:ConnectionString"],
                            DatabaseName = conf["MongoDb:DatabaseName"]
                        };
                        return new WeatherContext(options);
                    });

                    services.AddHealthChecks();
                    services.AddHttpClient("weather", (provider, client) =>
                    {
                        var configuration = provider.GetRequiredService<IConfiguration>();
                        var weatherApiUrl = configuration["CurrentWeatherService:weatherapi:url"];
                        client.BaseAddress = new Uri(weatherApiUrl);
                    });
                    services.AddTransient(c => new RestClient(c.GetService<IHttpClientFactory>().CreateClient("weather"))
                    {
                    }.For<IOpenWeatherApi>());
                })
                .ConfigureLogging(configureLogging =>
                {

                })
                .Build().Run();
        }
    }
}