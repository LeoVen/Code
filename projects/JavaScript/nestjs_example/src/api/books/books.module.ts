import { Module } from '@nestjs/common';
import { BooksController } from './books.controller';
import { BookService } from './books.service';
import { DatabaseModule } from 'src/database/database.module';

@Module({
    imports: [DatabaseModule],
    controllers: [BooksController],
    providers: [BookService],
    exports: [BookService],
})
export class BooksModule { }
