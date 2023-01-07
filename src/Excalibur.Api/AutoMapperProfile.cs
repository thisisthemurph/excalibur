using AutoMapper;
using Excalibur.Api.DTOs.Requests;
using Excalibur.Api.DTOs.Responses;
using Excalibur.Api.Models;

namespace Excalibur.Api;

public class AutoMapperProfile : Profile
{
	public AutoMapperProfile()
	{
        // DTO to entity model

        CreateMap<DataTemplate, DataTemplateResponse>();
        CreateMap<DataTemplateColumn, DataTemplateColumnResponse>();
        CreateMap<DataTemplateUploadedFileMetadata, DataTemplateUploadedFileMetadataResponse>();

        // Entity model to DTO

        CreateMap<DataTemplateCreateRequest, DataTemplate>();
        CreateMap<DataTemplateColumnRequest, DataTemplateColumn>();
    }
}
