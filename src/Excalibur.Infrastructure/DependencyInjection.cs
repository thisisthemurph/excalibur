using AutoMapper;
using Excalibur.Application.Common;
using Excalibur.Application.Repositories;
using Excalibur.Application.Services;
using Excalibur.Infrastructure.Persistence;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Excalibur.Infrastructure;
public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddAutoMapper(typeof(Profile));

        // Configure the database context

        services.Configure<MongoDbSettings>(configuration.GetSection("MongoDb"));
        services.AddSingleton<IApplicationDbContext, ApplicationDbContext>();

        // Configure Application Services

        services.AddScoped<IDataTemplateRepo, DataTemplateRepo>();
        services.AddScoped<IFileUploadService, FileUploadService>();

        return services;
    }
}
