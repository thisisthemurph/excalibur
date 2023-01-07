using System.ComponentModel.DataAnnotations;

namespace Excalibur.Application.DTOs.Requests;

public class DataTemplateCreateRequest
{
    [MinLength(6, ErrorMessage = "The name of the data template must be at least '6' characters in length.")]
    public string Name { get; set; }

    public List<DataTemplateColumnRequest> Columns { get; set; }
}
