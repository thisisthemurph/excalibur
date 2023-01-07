using Excalibur.Api.Models;

namespace Excalibur.Api.DTOs.Requests;

public class DataTemplateColumnRequest
{
    public string OriginalName { get; set; }
    public string PrettyName { get; set; }
    public string DataType { get; set; }
}
