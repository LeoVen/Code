export class MessageDto {
    readonly title: string
    readonly author: string

    constructor(partial: Partial<MessageDto>) {
        Object.assign(this, partial)
    }
}
