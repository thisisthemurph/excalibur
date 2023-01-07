using Excalibur.Domain.Models;

namespace Excalibur.Application.DTOs.Responses;

public class FileUploadStatusResponse
{
    public string DataTemplateId { get; set; }

    public string FileId { get; set; }

    public string Status { get; set; } = FileUploadStatus.Unknown;
}
