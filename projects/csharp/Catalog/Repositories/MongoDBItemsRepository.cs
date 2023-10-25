using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Catalog.Entitities;
using MongoDB.Bson;
using MongoDB.Driver;

namespace Catalog.Repositories
{
    public class MongoDBItemsRepository : IItemsRepository
    {
        private const string DatabaseName = "Catalog";
        private const string CollectionName = nameof(Item);
        private readonly IMongoCollection<Item> ItemsCollection;
        private readonly FilterDefinitionBuilder<Item> FilterBuilder = Builders<Item>.Filter;

        public MongoDBItemsRepository(IMongoClient mongoClient)
        {
            var db = mongoClient.GetDatabase(DatabaseName);
            ItemsCollection = db.GetCollection<Item>(CollectionName);
        }

        public async Task CreateItemAsync(Item item)
        {
            await ItemsCollection.InsertOneAsync(item);
        }

        public async Task DeleteItemAsync(Guid id)
        {
            await ItemsCollection.DeleteOneAsync(item => item.Id == id);
        }

        public async Task<Item> GetItemAsync(Guid id)
        {
            var filter = FilterBuilder.Eq(item => item.Id, id);
            return await ItemsCollection.Find(filter).SingleOrDefaultAsync();
        }

        public async Task<IEnumerable<Item>> GetItemsAsync()
        {
            return await ItemsCollection.Find(new BsonDocument()).ToListAsync();
        }

        public async Task UpdateItemAsync(Item item)
        {
            var filter = FilterBuilder.Eq(item => item.Id, item.Id);
            await ItemsCollection.ReplaceOneAsync(filter, item);
        }
    }
}
