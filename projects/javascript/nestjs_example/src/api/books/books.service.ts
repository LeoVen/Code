import { Injectable } from '@nestjs/common';
import { BookDto } from './book.dto';
import { DatabaseService } from 'src/database/database.service';

@Injectable()
export class BookService {
    constructor(private readonly database: DatabaseService) { }

    async createBook(book: BookDto): Promise<BookDto> {
        // to-do create a validation service
        // if book.title is empty/null or if book.author is empty/null
        return await this.database.createBook(book)
    }

    async getBooks(): Promise<BookDto[]> {
        return await this.database.listBooks()
    }
}
