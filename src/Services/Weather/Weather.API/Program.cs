using HealthChecks.UI.Client;
using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Diagnostics.HealthChecks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Routing;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using MongoDB.Bson;
using MongoDB.Driver;
using System.Threading.Tasks;
using Weather.Infrastructure;


namespace Weather.API
{
    class Program
    {
        static void Main(string[] args)
        {
            WebHost.CreateDefaultBuilder()
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
                })
                .Configure((context, app) =>
                    {
                        if (context.HostingEnvironment.IsDevelopment())
                        {
                            app.UseDeveloperExceptionPage();
                        }
                        app.UseRouting();
                        app.UseEndpoints(e =>
                        {
                            e.MapHealthChecks("/hc", new HealthCheckOptions()
                            {
                                Predicate = _ => true,
                                ResponseWriter = UIResponseWriter.WriteHealthCheckUIResponse
                            });
                            e.MapGet("/current", async c =>
                            {
                                var weatherContext = c.RequestServices.GetRequiredService<WeatherContext>();
                                var currentWeather = await weatherContext.Weather.Find(new BsonDocument()).FirstOrDefaultAsync();
                                await c.Response.WriteAsJsonAsync(currentWeather);
                            });
                        });
                    })
                .Build()
                .Run();

        }
    }
}