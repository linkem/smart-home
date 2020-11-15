using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Weather.Domain;
using Weather.Infrastructure;
using Weather.Job.Abstract;
using Weather.Job.OpenWeather;

namespace Weather.Job
{
    public class CurrentWeatherService : WeatherBackgroundService
    {
        public override Func<IServiceProvider, CancellationToken, Task> Work => async (prov, ct) =>
        {
            var openWeatherApi = prov.GetRequiredService<IOpenWeatherApi>();
            var weatherContext = prov.GetRequiredService<WeatherContext>();

            OpenWeatherModel currentWeather = null;
            try
            {
                logger.LogInformation($"[{ServiceName}] Fetching Weather");
                currentWeather = await openWeatherApi.GetWeather(stoppingToken: ct);
            }
            catch (Exception ex)
            {
                logger.LogError(ex.Message);
                throw;
            }

            var weatherEntity = new WeatherEntity
            {
                CityId = currentWeather.id,
                CityName = currentWeather.name,
                Conditions = currentWeather.weather.Select(s => new WeatherConditionEntity
                {
                    Description = s.description,
                    Icon = s.icon,
                    Id = s.id,
                    Main = s.main
                }),
                Clouds = currentWeather.clouds.all,
                DateTimeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.dt).UtcDateTime,
                Humidity = currentWeather.main.humidity,
                Temp = currentWeather.main.temp,
                SunriseDateTImeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.sys.sunrise).UtcDateTime,
                SunsetDateTImeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.sys.sunset).UtcDateTime,
            };
            logger.LogInformation($"[{ServiceName}] Inserting Current Weather");
            await weatherContext.Weather.InsertOneAsync(weatherEntity, cancellationToken: ct);
        };
        public CurrentWeatherService(ILogger<CurrentWeatherServiceOld> logger, IServiceProvider serviceProvider) : base(logger, serviceProvider)
        {
        }

        public override string ServiceName => nameof(CurrentWeatherServiceOld);
    }
}
