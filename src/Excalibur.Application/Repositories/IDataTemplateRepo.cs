using Excalibur.Application.DTOs.Requests;
using Excalibur.Domain.Entities;

namespace Excalibur.Application.Repositories;
public interface IDataTemplateRepo
{
    Task<bool> AddColumnAsync(string id, DataTemplateCreateColumnRequest column, CancellationToken cancellationToken = default);
    Task<DataTemplate> AddFileMetadata(string dataTemplateId, DataTemplateAddFileMetadataRequest metadata, CancellationToken cancellationToken = default);
    Task<DataTemplate> CreateAsync(DataTemplate dataTemplate);
    Task<bool> DeleteAsync(string id, CancellationToken cancellationToken = default);
    Task<IEnumerable<DataTemplate>> GetAsync();
    Task<DataTemplate> GetByIdAsync(string id, CancellationToken cancellationToken = default);
    Task<DataTemplate> UpdateAsync(string id, string dataTemplateName, CancellationToken cancellationToken = default);
}