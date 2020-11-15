using RestEase;
using System.Threading.Tasks;

namespace Weather.Job.OpenWeather
{
    //4c5d8f38918e66518b856052767d1ac2
    public interface IOpenWeatherApi
    {
        [Query("appId")]
        public string AppId { get; set; }
        [Query("units")]
        public string Units { get; set; }

        [Get("/weather")]
        Task<OpenWeatherModel> GetWeather([Query("q")]string cityName = "poznan", System.Threading.CancellationToken stoppingToken = default);

        [Get("/forecast")]
        Task<string> GetForecastWeather([Query("q")] string cityName = "poznan", System.Threading.CancellationToken stoppingToken = default);
    }
}
