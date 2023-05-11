mkdir mongodb && cd $_

npm init --yes
npm i -D typescript ts-node
npm i ts-node-dev --save-dev

./node_modules/.bin/tsc --init
mkdir src
echo 'console.log("Hello, World!")' > ./src/index.ts

npm i -D prisma
npm i @prisma/client
npx prisma init

# redo whenever necessary, to generate the new prisma client
npx prisma generate

# head over to https://github.com/minhhungit/mongodb-cluster-docker-compose
# setup mongodb cluster
docker-compose exec configsvr01 sh -c "mongosh < /scripts/init-configserver.js"

docker-compose exec shard01-a sh -c "mongosh < /scripts/init-shard01.js"
docker-compose exec shard02-a sh -c "mongosh < /scripts/init-shard02.js"
docker-compose exec shard03-a sh -c "mongosh < /scripts/init-shard03.js"

docker-compose exec router01 sh -c "mongosh < /scripts/init-router.js"

docker-compose exec router01 mongosh --port 27017

# inside the mongosh
# sh.enableSharding("MyDatabase")
# db.adminCommand( { shardCollection: "MyDatabase.MyCollection", key: { oemNumber: "hashed", zipCode: 1, supplierId: 1 } } )

# todo
docker run --name mongo-express -p 8081:8081 -e ME_CONFIG_MONGODB_URL="mongodb://127.0.0.1:27117,127.0.0.1:27118/MyDatabase?authSource=admin?replicaSet=rs-shard-01" -d mongo-express
