import { Kafka } from "kafkajs";

async function main(): Promise<void> {
    console.log('Consumer: Main: START')

    const kafka = new Kafka({
        clientId: 'consumer',
        brokers: ['0.0.0.0:9092'],
        requestTimeout: 60000,
    })

    const consumer = kafka.consumer({ groupId: 'test-group' })

    await consumer.connect()

    await consumer.subscribe({ topic: 'test-topic', fromBeginning: true })

    await consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
            if (message.value !== null) {
                console.log({
                    value: message.value.toString(),
                })
            }
        },
    })

    console.log('Consumer: Main: END')
}

main()
