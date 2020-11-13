#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:5.0-buster-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:5.0-buster-slim AS build
WORKDIR /src
COPY ["Weather.Job/Weather.Job.csproj", "Weather.Job/"]
COPY ["Weather.Shared/Weather.Shared.csproj", "Weather.Shared/"]
COPY ["Weather.Domain/Weather.Domain.csproj", "Weather.Domain/"]
RUN dotnet restore "Weather.Job/Weather.Job.csproj"
COPY . .
WORKDIR "/src/Weather.Job"
RUN dotnet build "Weather.Job.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Weather.Job.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Weather.Job.dll"]