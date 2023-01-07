using AutoMapper;
using Excalibur.Application.DTOs.Requests;
using Excalibur.Application.DTOs.Responses;
using Excalibur.Domain.Entities;

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
