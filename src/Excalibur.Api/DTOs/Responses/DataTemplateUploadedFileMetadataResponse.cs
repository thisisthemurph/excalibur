namespace Excalibur.Api.DTOs.Responses;

public class DataTemplateUploadedFileMetadataResponse
{
    public string? Id { get; set; }

    public string Name { get; set; } = string.Empty;

    public string StoredName { get; set; } = string.Empty;

    public string Status { get; set; } = string.Empty;
}
