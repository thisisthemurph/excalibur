using Excalibur.Api.Models;
using Microsoft.Extensions.Options;
using MongoDB.Driver;
using MongoDB.Driver.Linq;

namespace Excalibur.Api;

public class DatabaseContext
{
    private readonly IMongoCollection<DataTemplate> _dataTemplateCollection;

	public DatabaseContext(IOptions<MongoDBSettings> settings)
	{
        var mongoSettings = MongoClientSettings.FromConnectionString(settings.Value.ConnectionURI);
        mongoSettings.LinqProvider = LinqProvider.V3;
        
        var mongoClient = new MongoClient(mongoSettings);
        var database = mongoClient.GetDatabase(settings.Value.DatabaseName);

        _dataTemplateCollection = database.GetCollection<DataTemplate>(settings.Value.CollectionName);
    }

    public IMongoCollection<DataTemplate> DataTemplateCollection { get { return _dataTemplateCollection; } }
}
