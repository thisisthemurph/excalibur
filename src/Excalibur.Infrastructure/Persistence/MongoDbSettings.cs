namespace Excalibur.Infrastructure.Persistence;
public class MongoDbSettings
{
    public string? ConnectionURI { get; set; }
    public string? DatabaseName { get; set; }
    public string? CollectionName { get; set; }
};