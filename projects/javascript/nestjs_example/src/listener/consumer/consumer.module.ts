import { Module } from '@nestjs/common';
import { ConsumerService } from './consumer.service';
import { ConsumerController } from './consumer.controller';
import { BooksModule } from 'src/api/books/books.module';

@Module({
    imports: [BooksModule],
    providers: [ConsumerService],
    controllers: [ConsumerController]
})
export class ConsumerModule { }
