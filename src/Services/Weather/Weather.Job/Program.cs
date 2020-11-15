using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Configuration;
using Microsoft.AspNetCore.Builder;
using RestEase;
using Microsoft.AspNetCore.Diagnostics.HealthChecks;
using HealthChecks.UI.Client;
using Microsoft.AspNetCore.Http;
using Weather.Infrastructure;
using Weather.Job.OpenWeather;
using System;

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
                            endpoints.MapGet("/fc", async context =>
                            {
                                //TODO: to remove
                                var api = context.RequestServices.GetRequiredService<IOpenWeatherApi>();
                                var response = await api.GetForecastWeather();

                                await context.Response.WriteAsync(response);
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
                    services.AddSingleton(prov => {
                        var conf = prov.GetRequiredService<IConfiguration>();
                        var sampling = conf["CurrentWeatherService:SamplingInMinutes"];
                        if(!int.TryParse(sampling, out int samplingInt))
                            throw new ArgumentException("Value must be int.");
                        
                        return new CurrentWeatherServiceOptions { SamplingInMinutes = samplingInt };
                    });
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
                    services.AddTransient(provider => 
                    {
                        var configuration = provider.GetRequiredService<IConfiguration>();
                        var weatherApiUrl = configuration["weatherapi:url"];
                        var weatherApiKey = configuration["weatherapi:apikey"];
                        var weatherApi = RestClient.For<IOpenWeatherApi>(weatherApiUrl);
                        weatherApi.Units = "metric";
                        weatherApi.AppId = weatherApiKey;
                        return weatherApi;
                    });
                })
                .ConfigureLogging(configureLogging =>
                {

                })
                .Build().Run();
        }
    }
}