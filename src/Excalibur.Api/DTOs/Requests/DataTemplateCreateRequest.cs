using Excalibur.Api.DTOs;
using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs.Requests;

public class DataTemplateCreateRequest
{
    [StringLength(1)]
    public string Name { get; set; }

    public List<DataTemplateColumnDto> Columns { get; set; }
}
