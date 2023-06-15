import { Injectable } from '@nestjs/common';
import { MessageDto } from './message/message.dto';
import { BookService } from 'src/api/books/books.service';

@Injectable()
export class ConsumerService {
    constructor(private readonly bookService: BookService) { }

    createData(message: MessageDto): void {
        this.bookService.createBook({
            author: message.author,
            title: message.title,
        })
    }
}
