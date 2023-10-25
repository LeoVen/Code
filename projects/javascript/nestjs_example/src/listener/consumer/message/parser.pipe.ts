import { Injectable, ArgumentMetadata, PipeTransform } from '@nestjs/common';
import { MessageDto } from './message.dto';

@Injectable()
export class ParseMessagePipe implements PipeTransform<any, MessageDto> {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    transform(rawMessage: any, metadata: ArgumentMetadata): MessageDto {
        console.log('RAW MESSAGE: ', rawMessage)

        const { author, title } = rawMessage;

        const parsedMessage = new MessageDto({ author, title });

        return parsedMessage;
    }
}
