using Excalibur.Api.Models;

namespace Excalibur.Api.DTOs;

public class DataTemplateColumnDto
{
    public string? Id { get; set; }
    public string OriginalName { get; set; }
    public string PrettyName { get; set; }
    public string DataType { get; set; }
}
