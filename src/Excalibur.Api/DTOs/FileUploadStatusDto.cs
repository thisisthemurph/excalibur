using Excalibur.Api.Models;

namespace Excalibur.Api.DTOs;

public class FileUploadStatusDto
{
    public string DataTemplateId { get; set; }

    public string FileId { get; set; }

    public string Status { get; set; } = FileUploadStatus.Unknown;
}
