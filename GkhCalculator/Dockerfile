# syntax=docker/dockerfile:1

# Используем 9.0-alpine — это самая легкая версия
FROM mcr.microsoft.com/dotnet/sdk:10.0-alpine AS build
WORKDIR /src

# Включаем принудительное копирование всех csproj, если их несколько
COPY *.csproj ./
RUN dotnet restore

COPY . .
# Важно: добавляем флаги для экономии памяти
RUN dotnet publish -c Release -o /app/publish /p:UseAppHost=false /p:PublishReadyToRun=false /p:TieredCompilation=false

FROM mcr.microsoft.com/dotnet/aspnet:10.0-alpine AS final
WORKDIR /app

# На Render лучше использовать порт 80 или 10000
ENV ASPNETCORE_URLS=http://+:80
EXPOSE 80

COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "GkhCalculator.dll"]