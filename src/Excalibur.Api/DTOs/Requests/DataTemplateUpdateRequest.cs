using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs.Requests;

public class DataTemplateUpdateRequest
{
    [MinLength(1)]
    public string Name { get; set; }
}
