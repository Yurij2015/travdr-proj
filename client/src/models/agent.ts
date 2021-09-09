import {Organization} from "@/models/organization";

export class Agent {
    constructor(
        public id: number = 0,
        public full_name: string = '',
        public phone_number: string = '',
        public organization: Organization = new Organization(),
        public image: string = '',
    ) {
    }
}