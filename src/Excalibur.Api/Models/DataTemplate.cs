namespace Excalibur.Api.Models;

using Excalibur.Api.DTOs.Responses;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

public class DataTemplate
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string? Id { get; set; }

    [BsonElement("name")]
    public string Name { get; set; } = string.Empty;

    [BsonElement("columns")]
    public List<DataTemplateColumn>? Columns { get; set; }

    [BsonElement("files")]
    public List<DataTemplateUploadedFileMetadata> Files { get; set; }
}

