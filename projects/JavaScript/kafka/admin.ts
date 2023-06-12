import { Kafka } from "kafkajs";

async function main(): Promise<void> {
    console.log('Create Topic: Main: START')

    const kafka = new Kafka({
        clientId: 'admin',
        brokers: ['0.0.0.0:9092'],
        requestTimeout: 10000,
    })

    const admin = kafka.admin()

    await admin.connect()

    try {
        let topics = await admin.listTopics()
        if (!topics.find(x => x === 'test-topic')) {
            console.log('Creating topic: test-topic')
            await admin.createTopics({ topics: [{ topic: 'test-topic', replicationFactor: 1 }] })
        } else {
            console.log('Topics that already exist: ', topics)
        }
        let overview = await admin.listGroups()
        console.log('Groups that already exist: ', overview.groups)
    } catch (e) {
        console.error(e)
    } finally {
        admin.disconnect()
    }

    console.log('Create Topic: Main: END')
}

main()
