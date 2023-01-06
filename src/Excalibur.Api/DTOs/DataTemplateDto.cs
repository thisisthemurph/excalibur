using Excalibur.Api.Models;

namespace Excalibur.Api.DTOs;

public class DataTemplateDto
{
    public string? Id { get; set; }
    public string Name { get; set; }
    public List<DataTemplateColumnDto> Columns { get; set; }

    public DataTemplate MapToModel()
    {
        return new DataTemplate
        {
            Id = Id,
            Name = Name,
            Columns = Columns.Select(c => new DataTemplateColumn 
            { 
                OriginalName = c.OriginalName, 
                PrettyName = c.PrettyName,
                DataType = c.DataType
            }).ToList(),
        };
    }
}
