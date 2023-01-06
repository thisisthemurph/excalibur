namespace Excalibur.Api.Models;

using Excalibur.Api.DTOs;
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

    public DataTemplateDto MapToDto()
    {
        return new DataTemplateDto
        {
            Id = Id,
            Name = Name,
            Columns = Columns.Select(c => new DataTemplateColumnDto
            {
                Id = c.Id.ToString(),
                OriginalName = c.OriginalName,
                PrettyName = c.PrettyName,
                DataType = c.DataType
            }).ToList(),
        };
    }
}

