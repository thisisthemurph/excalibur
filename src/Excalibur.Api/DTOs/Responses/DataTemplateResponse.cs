using Excalibur.Api.Models;

namespace Excalibur.Api.DTOs.Responses;

public class DataTemplateResponse
{
    public string? Id { get; set; }

    public string Name { get; set; }

    public List<DataTemplateColumnDto> Columns { get; set; }

    public List<DataTemplateUploadedFileMetadata> Files { get; set; }
}

