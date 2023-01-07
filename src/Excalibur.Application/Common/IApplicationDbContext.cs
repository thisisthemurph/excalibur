using Excalibur.Domain.Entities;
using MongoDB.Driver;

namespace Excalibur.Application.Common;

public interface IApplicationDbContext
{
    IMongoCollection<DataTemplate> DataTemplateCollection { get; }
}
