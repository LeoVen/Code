import { NestFactory } from '@nestjs/core';
import { AppModule } from './api/app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ListenerModule } from './listener/listener.module';

async function bootstrap() {
    const listener = await NestFactory.createMicroservice<MicroserviceOptions>(
        ListenerModule,
        {
            transport: Transport.KAFKA,
            options: {
                client: {
                    clientId: 'kafka-listener',
                    brokers: ['localhost:9092'],
                },
                consumer: {
                    groupId: 'test-group'
                }
            }
        },
    )

    const app = await NestFactory.create(AppModule)

    console.log('Starting kafka listener')
    console.log('Starting app')

    await Promise.all([listener.listen(), app.listen(3000)])
}
bootstrap();
