export class Customer {
    constructor(
        public id: number = 0,
        public full_name: string = '',
        public gender: string = '',
        public birth_date: string = '',
        public birth_place: string = '',
        public image: string = ''
    ) {
    }
}