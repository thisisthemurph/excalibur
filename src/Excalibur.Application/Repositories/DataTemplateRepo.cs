using MongoDB.Driver;
using MongoDB.Bson;
using MongoDB.Driver.Linq;
using Excalibur.Application.Common;
using Excalibur.Domain.Entities;
using Excalibur.Application.DTOs.Requests;

namespace Excalibur.Application.Repositories;

public class DataTemplateRepo : IDataTemplateRepo
{
    private readonly IApplicationDbContext _dbContext;
    private readonly IMongoCollection<DataTemplate> _dataTemplateCollection;

    public DataTemplateRepo(IApplicationDbContext dbContext)
    {
        _dbContext = dbContext ?? throw new ArgumentNullException(nameof(dbContext));
        _dataTemplateCollection = _dbContext.DataTemplateCollection;
    }

    public async Task<IEnumerable<DataTemplate>> GetAsync()
    {
        return await _dataTemplateCollection.AsQueryable().ToListAsync();
    }

    public async Task<DataTemplate> GetByIdAsync(string id, CancellationToken cancellationToken = default)
    {
        return await _dataTemplateCollection.AsQueryable()
            .Where(t => t.Id == id)
            .SingleOrDefaultAsync(cancellationToken);
    }

    public async Task<DataTemplate> CreateAsync(DataTemplate dataTemplate)
    {
        var exists = await ExistsWithNameAsync(dataTemplate.Name);
        if (exists)
        {
            throw new ArgumentException($"A data template with the name '{dataTemplate.Name}' already exists.");
        }

        await _dataTemplateCollection.InsertOneAsync(dataTemplate);
        return dataTemplate;
    }

    public async Task<bool> AddColumnAsync(string id, DataTemplateCreateColumnRequest column, CancellationToken cancellationToken = default)
    {
        var entity = new DataTemplateColumn
        {
            OriginalName = column.OriginalName,
            PrettyName = column.PrettyName,
            DataType = column.DataType,
        };

        var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
        var update = Builders<DataTemplate>.Update.AddToSet("Columns", entity);

        var result = await _dataTemplateCollection
            .UpdateOneAsync(filter, update, new UpdateOptions() { IsUpsert = true }, cancellationToken);

        return result.ModifiedCount == 1;
    }

    public async Task<DataTemplate> UpdateAsync(string id, string dataTemplateName, CancellationToken cancellationToken = default)
    {
        var exists = await ExistsWithNameAsync(dataTemplateName);
        if (exists)
        {
            throw new ArgumentException($"A data template with the name '{dataTemplateName}' already exists.");
        }

        var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
        var update = Builders<DataTemplate>.Update
            .Set(t => t.Name, dataTemplateName);

        var options = new FindOneAndUpdateOptions<DataTemplate>
        {
            ReturnDocument = ReturnDocument.After,
        };

        return await _dataTemplateCollection.FindOneAndUpdateAsync(
            filter,
            update,
            options,
            cancellationToken: cancellationToken);
    }

    public async Task<bool> DeleteAsync(string id, CancellationToken cancellationToken = default)
    {
        var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
        var result = await _dataTemplateCollection.DeleteOneAsync(filter, cancellationToken);

        return result.DeletedCount == 1;
    }

    public async Task<DataTemplate> AddFileMetadata(
        string dataTemplateId,
        DataTemplateAddFileMetadataRequest metadata,
        CancellationToken cancellationToken = default)
    {
        var entity = new DataTemplateUploadedFileMetadata
        {
            Id = ObjectId.GenerateNewId().ToString(),
            Name = metadata.Name,
            StoredName = metadata.StoredName,
            Status = metadata.Status,
        };

        var filter = Builders<DataTemplate>.Filter.Eq("Id", dataTemplateId);
        var update = Builders<DataTemplate>.Update
            .Push(t => t.Files, entity);

        var options = new FindOneAndUpdateOptions<DataTemplate>
        {
            ReturnDocument = ReturnDocument.After,
        };

        var updateResult = await _dataTemplateCollection.FindOneAndUpdateAsync(filter, update, options, cancellationToken: cancellationToken);
        return updateResult;
    }

    private async Task<bool> ExistsWithNameAsync(string dataTemplateName)
    {
        var numberOfResults = await _dataTemplateCollection
            .AsQueryable()
            .Where(x => x.Name == dataTemplateName)
            .CountAsync();

        return numberOfResults > 0;
    }
}

public interface IDataTemplateService
{
}