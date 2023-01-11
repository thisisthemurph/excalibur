using Ardalis.SmartEnum.JsonNet;
using Excalibur.Domain.Enums;
using System.Text.Json.Serialization;

namespace Excalibur.Application.DTOs.Responses;

public class DataTemplateColumnResponse
{
    public string? Id { get; set; }
    public string OriginalName { get; set; } = string.Empty;
    public string PrettyName { get; set; } = string.Empty;
    public ColumnDataType DataType { get; set; }
    public string DataTypeValue { get; set; } = string.Empty;
}
