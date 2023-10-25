import { Kafka } from "kafkajs";

async function main(): Promise<void> {
    console.log('Producer: Main: START')

    const kafka = new Kafka({
        clientId: 'producer',
        brokers: ['0.0.0.0:9092'],
        requestTimeout: 10000,
    })

    const producer = kafka.producer()

    await producer.connect()

    await producer.send({
        topic: 'test-topic',
        messages: [
            { value: 'Hello KafkaJS user!' },
        ],
    })

    await producer.disconnect()

    console.log('Producer: Main: END')
}

main()
