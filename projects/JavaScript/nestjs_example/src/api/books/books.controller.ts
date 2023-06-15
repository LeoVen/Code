import { Body, Controller, Get, Post } from '@nestjs/common';
import { BookService } from './books.service';
import { BookDto } from './book.dto';

@Controller('books')
export class BooksController {

    constructor(private readonly bookService: BookService) { }

    @Get('/')
    async listBooks(): Promise<BookDto[]> {
        return await this.bookService.getBooks()
    }

    @Post('/')
    async createBook(@Body() body: BookDto): Promise<BookDto> {
        return await this.bookService.createBook(body)
    }
}
