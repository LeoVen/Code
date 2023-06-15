import { Injectable } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';
import { BookDto } from 'src/api/books/book.dto';

@Injectable()
export class DatabaseService {
    private prisma: PrismaClient

    constructor() {
        this.prisma = new PrismaClient({ log: ['query'] })
    }

    async createBook(book: BookDto): Promise<BookDto> {
        return await this.prisma.book.create({
            data: {
                author: book.author,
                title: book.title,
            }
        })
    }

    async listBooks(): Promise<BookDto[]> {
        let books = await this.prisma.book.findMany()
        return books.map(book => {
            return {
                author: book.author,
                title: book.title,
            }
        })
    }
}
