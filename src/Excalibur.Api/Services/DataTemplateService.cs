namespace Excalibur.Api.Services;

using Microsoft.Extensions.Options;
using MongoDB.Driver;
using MongoDB.Bson;
using Excalibur.Api.Models;
using MongoDB.Driver.Linq;

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

	public async Task<DataTemplate> CreateAsync(DataTemplate dataTemplate)
	{
		await _dataTemplateCollection.InsertOneAsync(dataTemplate);
		return dataTemplate;
	}

	public async Task<bool> AddColumnAsync(string id, DataTemplateColumn column, CancellationToken cancellationToken = default)
	{
		var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
		var update = Builders<DataTemplate>.Update.AddToSet<DataTemplateColumn>("Columns", column);
		
		var result = await _dataTemplateCollection
			.UpdateOneAsync(filter, update, new UpdateOptions() { IsUpsert = true }, cancellationToken);

		return result.ModifiedCount == 1;
	}

	public async Task<DataTemplate> UpdateAsync(string id, string dateTemplateName, CancellationToken cancellationToken = default)
	{
		var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
		var update = Builders<DataTemplate>.Update
			.Set(t => t.Name, dateTemplateName);
		
		return await _dataTemplateCollection.FindOneAndUpdateAsync(
			filter, 
			update, 
			cancellationToken: cancellationToken);
	}

	public async Task<bool> DeleteAsync(string id, CancellationToken cancellationToken = default)
	{
		var filter = Builders<DataTemplate>.Filter.Eq("Id", id);
		var result = await _dataTemplateCollection.DeleteOneAsync(filter, cancellationToken);

		return result.DeletedCount == 1;
	}
}
