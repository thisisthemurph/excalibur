namespace Excalibur.Api.Services;

using MongoDB.Driver;
using MongoDB.Bson;
using Excalibur.Api.Models;
using MongoDB.Driver.Linq;
using Excalibur.Api.DTOs.Requests;

public class DataTemplateService
{
	private readonly DatabaseContext _dbContext;
    private readonly IMongoCollection<DataTemplate> _dataTemplateCollection;

	public DataTemplateService(DatabaseContext dbContext)
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

	public async Task<DataTemplate> CreateAsync(DataTemplateCreateRequest dataTemplate)
	{
		var entity = new DataTemplate
		{
			Name = dataTemplate.Name,
			Columns = dataTemplate.Columns.Select(
				c => new DataTemplateColumn
				{
					OriginalName = c.OriginalName,
					PrettyName = c.PrettyName,
					DataType = c.DataType,
				}).ToList(),
		};

		await _dataTemplateCollection.InsertOneAsync(entity);
		return entity;
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
		var update = Builders<DataTemplate>.Update.AddToSet<DataTemplateColumn>("Columns", entity);
		
		var result = await _dataTemplateCollection
			.UpdateOneAsync(filter, update, new UpdateOptions() { IsUpsert = true }, cancellationToken);

		return result.ModifiedCount == 1;
	}

	public async Task<DataTemplate> UpdateAsync(string id, string dateTemplateName, CancellationToken cancellationToken = default)
	{
		var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
		var update = Builders<DataTemplate>.Update
			.Set(t => t.Name, dateTemplateName);

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

	public async Task<DataTemplate> AddFileMetadata(string id, DataTemplateUploadedFileMetadata metadata, CancellationToken cancellationToken = default)
	{
		metadata.Id = ObjectId.GenerateNewId().ToString();

		var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
		var update = Builders<DataTemplate>.Update
			.Push<DataTemplateUploadedFileMetadata>(t => t.Files, metadata);

		var options = new FindOneAndUpdateOptions<DataTemplate>
		{
			ReturnDocument = ReturnDocument.After,
		};

		var updateResult = await _dataTemplateCollection.FindOneAndUpdateAsync(filter, update, options, cancellationToken: cancellationToken);
		return updateResult;
	}
}
