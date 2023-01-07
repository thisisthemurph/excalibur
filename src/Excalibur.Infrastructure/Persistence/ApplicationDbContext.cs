using Excalibur.Application.Common;
using Excalibur.Domain.Entities;
using Microsoft.Extensions.Options;
using MongoDB.Driver;
using MongoDB.Driver.Linq;

namespace Excalibur.Infrastructure.Persistence;

public class ApplicationDbContext : IApplicationDbContext
{
    private readonly IMongoCollection<DataTemplate> _dataTemplateCollection;

    public ApplicationDbContext(IOptions<MongoDbSettings> settings)
    {
        var mongoSettings = MongoClientSettings.FromConnectionString(settings.Value.ConnectionURI);
        mongoSettings.LinqProvider = LinqProvider.V3;

        var mongoClient = new MongoClient(mongoSettings);
        var database = mongoClient.GetDatabase(settings.Value.DatabaseName);

        _dataTemplateCollection = database.GetCollection<DataTemplate>(settings.Value.CollectionName);
    }

    public IMongoCollection<DataTemplate> DataTemplateCollection { get { return _dataTemplateCollection; } }
}
