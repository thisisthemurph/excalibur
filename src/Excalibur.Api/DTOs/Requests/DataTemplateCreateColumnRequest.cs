using Excalibur.Api.Models;
using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs.Requests;

public class DataTemplateCreateColumnRequest
{
    [Required]
    [MinLength(1, ErrorMessage = "The original name field must be at least 1 character in length.")]
    public string OriginalName { get; set; }

    [Required]
    [MinLength(1, ErrorMessage = "The pretty name field must be at least 1 character in length.")]
    public string PrettyName { get; set; }

    [Required]
    public string DataType { get; set; }
}
