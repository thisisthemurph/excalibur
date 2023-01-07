using Excalibur.Api.Models;
using System.ComponentModel.DataAnnotations;

namespace Excalibur.Api.DTOs.Requests;

public class DataTemplateCreateColumnRequest
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
}
