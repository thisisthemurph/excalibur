using MongoDB.Bson.Serialization.Attributes;
using MongoDB.Bson;
using Excalibur.Domain.Enums;

namespace Excalibur.Domain.Entities;

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
    public ColumnDataType DataType { get; set; }
}
