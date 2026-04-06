using GkhCalculator.Services;
using GkhCalculator.Models;
using GkhCalculator.Controllers;

var builder = WebApplication.CreateBuilder(args);

// Swagger / OpenAPI
builder.Services.AddOpenApi();

builder.Services.Configure<GkhTariffsOptions>(
    builder.Configuration.GetSection(GkhTariffsOptions.SectionName));

// 👉 Регистрируем сервис
builder.Services.AddScoped<GkhCalculatorService>();

builder.Services.AddCors(options =>
{
    options.AddDefaultPolicy(policy =>
    {
        policy.AllowAnyOrigin()
              .AllowAnyHeader()
              .AllowAnyMethod();
    });
});

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.UseHttpsRedirection();
}

app.UseCors();


app.MapPost("/api/gkh/calculate", (GkhRequest request, GkhCalculatorService service) =>
{
    var result = service.Calculate(request);
    return Results.Ok(result);
});


var summaries = new[]
{
    "Freezing", "Bracing", "Chilly", "Cool", "Mild"
};

app.MapGet("/weatherforecast", () =>
{
    return Enumerable.Range(1, 3).Select(index =>
        new
        {
            Date = DateTime.Now.AddDays(index),
            Temp = Random.Shared.Next(-10, 30)
        });
});

app.UseDefaultFiles();
app.UseStaticFiles();

app.Run();