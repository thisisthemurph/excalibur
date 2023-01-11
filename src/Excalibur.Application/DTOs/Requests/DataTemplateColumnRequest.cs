using Excalibur.Domain.Enums;
using System.ComponentModel.DataAnnotations;

namespace Excalibur.Application.DTOs.Requests;

public class DataTemplateColumnRequest
{
    [Required]
    [MinLength(1, ErrorMessage = "The original name field must be at least 1 character in length.")]
    public string OriginalName { get; set; } = string.Empty;

    [Required]
    [MinLength(1, ErrorMessage = "The pretty name field must be at least 1 character in length.")]
    public string PrettyName { get; set; } = string.Empty;

    [Required]
    public ColumnDataType DataType { get; set; }
}
