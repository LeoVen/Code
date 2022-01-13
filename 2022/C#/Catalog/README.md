# .NET Core 5

* .NET Core 5 API
* MongoDB
* Docker
* Secret Manager

Start Mongo Docker

```
docker run -d --rm --name mongo -p 27017:27017 -v mongodbdata:/data/db -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=Pass#word1 mongo
```
