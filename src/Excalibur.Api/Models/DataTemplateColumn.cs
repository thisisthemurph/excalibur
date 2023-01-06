using MongoDB.Bson.Serialization.Attributes;
using MongoDB.Bson;

namespace Excalibur.Api.Models;

public class DataTemplateColumn
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public ObjectId Id { get; set; } = ObjectId.GenerateNewId();

    [BsonElement("originalName")]
    public string OriginalName { get; set; }

    [BsonElement("prettyName")]
    public string PrettyName { get; set; }

    [BsonElement("dataType")]
    public string DataType { get; set; }
}
