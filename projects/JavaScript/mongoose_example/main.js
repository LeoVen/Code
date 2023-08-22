const { Schema } = (mongoose = require('mongoose'));

mongoose.Promise = global.Promise;
mongoose.set('debug', true);

const uri = 'mongodb://localhost:27017/manydemo',
    options = { useNewUrlParser: true };

const itemSchema = new Schema({
    name: String,
    stores: [{ type: Schema.Types.ObjectId, ref: 'Store' }],
});

const storeSchema = new Schema({
    name: String,
    items: [{ type: Schema.Types.ObjectId, ref: 'Item' }],
});

const Item = mongoose.model('Item', itemSchema);
const Store = mongoose.model('Store', storeSchema);

const log = (data) => console.log(JSON.stringify(data, undefined, 2));

(async function () {
    try {
        const conn = await mongoose.connect(uri, options);

        // Clean data
        await Promise.all(Object.entries(conn.models).map(([k, m]) => m.deleteMany()));

        // Create some instances
        let [toothpaste, brush] = ['toothpaste', 'brush'].map((name) => new Item({ name }));

        let [billsStore, tedsStore] = ['Bills', 'Teds'].map((name) => new Store({ name }));

        // Add items to stores
        [billsStore, tedsStore].forEach((store) => {
            store.items.push(toothpaste); // add toothpaste to store
            toothpaste.stores.push(store); // add store to toothpaste
        });

        // Brush is only in billsStore
        billsStore.items.push(brush);
        brush.stores.push(billsStore);

        // Save everything
        await Promise.all([toothpaste, brush, billsStore, tedsStore].map((m) => m.save()));

        // Show stores
        let stores = await Store.find().populate('items', '-stores');
        log(stores);

        // Show items
        let items = await Item.find().populate('stores', '-items');
        log(items);
    } catch (e) {
        console.error(e);
    } finally {
        mongoose.disconnect();
    }
})();
