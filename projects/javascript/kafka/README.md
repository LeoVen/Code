# Kafka Example with TypeScript

* Kafka
* Zookeeper

Topic name: `test-topic`
GroupId used: `test-group`

## How to use

* Run `docker compose up` to spin up kafka and zookeeper
* Run `npm run admin` to setup a topic
  * Can also be reused to list topics and groups
* Run `npm run consumer` to spin up the consumer
* Run `npm run producer` to send messages
