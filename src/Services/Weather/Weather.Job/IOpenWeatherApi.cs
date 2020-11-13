using RestEase;
using System.Threading.Tasks;

namespace Weather.Job
{
    public interface IOpenWeatherApi
    {
        [Get("?units=metric&q=poznan&appid=4c5d8f38918e66518b856052767d1ac2")]
        Task<OpenWeatherModel> GetWeather(System.Threading.CancellationToken stoppingToken);
    }
}
