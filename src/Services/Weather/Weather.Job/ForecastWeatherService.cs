using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;
using System.Threading;
using System.Threading.Tasks;
using Weather.Infrastructure;
using Weather.Job.Abstract;
using Weather.Job.OpenWeather;

namespace Weather.Job
{
    public class ForecastWeatherService : WeatherBackgroundService
    {
        //private readonly IOpenWeatherApi openWeatherApi;
        //private readonly WeatherContext weatherContext;
        private readonly IServiceProvider services;

        public ForecastWeatherService(ILogger<WeatherBackgroundService> logger, IServiceProvider services) : base(logger, services)
        {
            //this.openWeatherApi = openWeatherApi;
            //this.weatherContext = weatherContext;
            this.services = services;
        }
        public override string ServiceName => nameof(ForecastWeatherService);

        public override Func<IServiceProvider, CancellationToken, ValueTask<bool>> Work => async (prov, ct) =>
        {

            var api = prov.GetRequiredService<IOpenWeatherApi>();
            var context = prov.GetRequiredService<WeatherContext>();




            return true;
        };
    }
}
