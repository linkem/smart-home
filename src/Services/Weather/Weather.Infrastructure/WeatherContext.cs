using MongoDB.Driver;
using Weather.Domain;

namespace Weather.Infrastructure
{
    public class WeatherContext
    {
        private readonly MongoClient _mongoClient;
        private readonly IMongoDatabase _database;

        public WeatherContext(MongoDbOptions options)
        {
            _mongoClient = new MongoClient(options.ConnectionString);
            _database = _mongoClient.GetDatabase(options.DatabaseName);

        }
        public IMongoCollection<WeatherEntity> Weather => _database.GetCollection<WeatherEntity>("weather");
        public IMongoCollection<WeatherEntity> ForecastWeather => _database.GetCollection<WeatherEntity>("forecast");

    }
}
