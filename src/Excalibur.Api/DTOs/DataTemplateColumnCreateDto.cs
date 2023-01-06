using Excalibur.Api.Models;
using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs;

public class DataTemplateColumnCreateDto
{
    [Required]
    [StringLength(1)]
    public string OriginalName { get; set; }

    [Required]
    [StringLength(1)]
    public string PrettyName { get; set; }

    [Required]
    [StringLength(1)]
    public string DataType { get; set; }

    public DataTemplateColumn MapToModel()
    {
        return new DataTemplateColumn
        {
            OriginalName = OriginalName,
            PrettyName = PrettyName,
            DataType = DataType,
        };
    }
}
