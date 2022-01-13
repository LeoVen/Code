# .NET Core 5

* .NET Core 5 API
* MongoDB
* Docker
* Secret Manager

## Start Mongo Docker

```
docker run -d --rm --name mongo -p 27017:27017 -v mongodbdata:/data/db -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=Pass#word1 mongo
```

## Build docker image from project

```
docker build -t catalog:v1 .
```

## Create a docker network

```
docker network create net5tutorial
```

## Run network

* Database

```
docker run -d --rm --name mongo -p 27017:27017 -v mongodbdata:/data/db -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=Pass#word1 --network=net5tutorial mongo
```

* Api

```
docker run --rm -p 8080:80 -e MongoDbSettings:Host=mongo -e MongoDbSettings:Password=Pass#word1 --network=net5tutorial catalog_v1
```
