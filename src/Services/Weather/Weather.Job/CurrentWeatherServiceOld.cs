using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using MongoDB.Driver;
using System;
using System.Diagnostics;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Weather.Domain;
using Weather.Infrastructure;
using Weather.Job.OpenWeather;

namespace Weather.Job
{
    public class CurrentWeatherServiceOld : BackgroundService, ISmartHomeBackgroundService
    {
        private Stopwatch _watch;
        private readonly string ServiceName;
        private readonly ILogger<CurrentWeatherServiceOld> _logger;
        private readonly IOpenWeatherApi _openWeatherApi;
        private readonly WeatherContext _weatherContext;
        private readonly CurrentWeatherServiceOptions _options;
        private IMongoCollection<WeatherEntity> _weather;

        public CurrentWeatherServiceOld(ILogger<CurrentWeatherServiceOld> logger, 
            IOpenWeatherApi openWeatherApi, 
            WeatherContext weatherContext, 
            CurrentWeatherServiceOptions options)
        {
            ServiceName = nameof(CurrentWeatherServiceOld);
            _logger = logger;
            _openWeatherApi = openWeatherApi;
            _weatherContext = weatherContext;
            _options = options;
            _weather = _weatherContext.Weather;
            //var client = new MongoClient("mongodb://mlinke:Ny1y5SuGBeNAh8duSBz3x8HSMlVrSvTFcjLXjMuMMV8jxN4G2bMDZfCH0SGpZuN67OZ3AQjXs4CRWISrxqUcKw==@mlinke.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@mlinke@");
            //_weather = database.GetCollection<WeatherEntity>("weather");
        }
        public Task StartAsync(CancellationToken cancellationToken)
        {
            _logger.LogInformation($"[{ServiceName}] Started Async");
            return Task.CompletedTask;
        }
        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            await DoWork(stoppingToken);
            await Task.Delay(10000, stoppingToken);
        }

        //Change to UoW method in future
        private async Task DoWork(CancellationToken stoppingToken)
        {
            _logger.LogInformation($"[{ServiceName}] Started Work");
            _watch = new Stopwatch();
            _watch.Start();
            _logger.LogInformation($"[{ServiceName}] Start Fetch Weather");
            OpenWeatherModel currentWeather = null;
            try
            {
                currentWeather = await _openWeatherApi.GetWeather(stoppingToken: stoppingToken);
            }
            catch(Exception ex)
            {
                _logger.LogError(ex.Message);
                throw;
            }
            _logger.LogInformation($"[{ServiceName}] Finish Fetch Weather");
            _logger.LogInformation($"[{ServiceName}] Start Insert");
            var weatherEntity = new WeatherEntity
            {
                CityId = currentWeather.id,
                CityName = currentWeather.name,
                Conditions = currentWeather.weather.Select(s => new WeatherConditionEntity
                {
                    Description = s.description,
                    Icon = s.icon,
                    Id = s.id,
                    Main =s.main
                }),
                Clouds = currentWeather.clouds.all,
                DateTimeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.dt).UtcDateTime,
                Humidity = currentWeather.main.humidity,
                Temp = currentWeather.main.temp,
                SunriseDateTImeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.sys.sunrise).UtcDateTime,
                SunsetDateTImeUtc = DateTimeOffset.FromUnixTimeSeconds(currentWeather.sys.sunset).UtcDateTime,
            };
            await _weather.InsertOneAsync(weatherEntity, cancellationToken: stoppingToken);
            _watch.Stop();
            
            _logger.LogInformation($"[{ServiceName}] Finish Insert Elapsed: {_watch.ElapsedMilliseconds} ms");
        }

        public override Task StopAsync(CancellationToken cancellationToken)
        {
            _logger.LogInformation("Stop Async");
            return Task.CompletedTask;
        }


    }
}
