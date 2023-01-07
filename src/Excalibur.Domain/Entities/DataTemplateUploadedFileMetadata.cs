using MongoDB.Bson.Serialization.Attributes;
using MongoDB.Bson;

namespace Excalibur.Domain.Entities;

public class DataTemplateUploadedFileMetadata
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string? Id { get; set; }

    [BsonElement("name")]
    public string Name { get; set; } = string.Empty;

    [BsonElement("storedName")]
    public string StoredName { get; set; } = string.Empty;

    [BsonElement("status")]
    public string Status { get; set; } = string.Empty;
}
