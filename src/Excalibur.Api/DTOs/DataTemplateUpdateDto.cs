using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs;

public class DataTemplateUpdateDto
{
    [MinLength(1)]
    public string Name { get; set; }
}
