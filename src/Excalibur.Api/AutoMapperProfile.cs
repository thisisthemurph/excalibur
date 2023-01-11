using AutoMapper;
using AutoMapper.Extensions.EnumMapping;
using Excalibur.Application.DTOs.Requests;
using Excalibur.Application.DTOs.Responses;
using Excalibur.Domain.Entities;
using Excalibur.Domain.Enums;
using Excalibur.Domain.ExtensionMethods;

namespace Excalibur.Api;

public class AutoMapperProfile : Profile
{
	public AutoMapperProfile()
	{
        // DTO to entity model

        CreateMap<DataTemplate, DataTemplateResponse>();

        CreateMap<DataTemplateColumn, DataTemplateColumnResponse>()
            .ForMember
            (
                dest => dest.DataTypeValue,
                opt => opt.MapFrom
                (
                    src => src.DataType.GetDescription()
                )
            );

        CreateMap<DataTemplateUploadedFileMetadata, DataTemplateUploadedFileMetadataResponse>();

        // Entity model to DTO

        CreateMap<DataTemplateCreateRequest, DataTemplate>();

        CreateMap<DataTemplateColumnRequest, DataTemplateColumn>();
    }
}
