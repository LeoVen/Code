import { Controller } from '@nestjs/common';
import { ConsumerService } from './consumer.service';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { ParseMessagePipe } from './message/parser.pipe';

@Controller()
export class ConsumerController {
    constructor(private readonly consumerService: ConsumerService) { }

    @MessagePattern('test-topic')
    getMessage(@Payload(new ParseMessagePipe()) message): void {
        return this.consumerService.createData(message)
    }
}
