using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Diagnostics;
using System.Threading;
using System.Threading.Tasks;

namespace Weather.Job.Abstract
{
    public abstract class WeatherBackgroundService : BackgroundService
    {
        protected readonly ILogger logger;
        private readonly IServiceProvider services;
        private Stopwatch watch;

        public abstract string ServiceName { get; }
        public abstract Func<IServiceProvider, CancellationToken, Task> Work { get; }
        public WeatherBackgroundService(ILogger logger, IServiceProvider services)
        {
            this.logger = logger;
            this.services = services;
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            logger.LogInformation($"[{ServiceName}] Started ExecuteAsync");
            watch = new Stopwatch();
            watch.Start();
            using (var scope = services.CreateScope())
                await Work(scope.ServiceProvider, stoppingToken);

            //TODO: add delay from settings
            await Task.Delay(10000, stoppingToken);
            watch.Stop();
            logger.LogInformation($"[{ServiceName}] Finish ExecuteAsync, Elapsed: {watch.ElapsedMilliseconds} ms");
        }

        public override Task StartAsync(CancellationToken cancellationToken)
        {
            logger.LogInformation($"[{ServiceName}] Started Async");
            return base.StartAsync(cancellationToken);
        }

        public override Task StopAsync(CancellationToken cancellationToken)
        {
            logger.LogInformation($"[{ServiceName}] Stopped Async");
            return base.StopAsync(cancellationToken);
        }
    }
}
