namespace Excalibur.Application.DTOs.Responses;

public class DataTemplateResponse
{
    public string? Id { get; set; }

    public string Name { get; set; }

    public List<DataTemplateColumnResponse> Columns { get; set; }

    public List<DataTemplateUploadedFileMetadataResponse> Files { get; set; }
}

