import {Organization} from "@/models/organization";
import {Agreement} from "@/models/agreement"
import {Agent} from "@/models/agent";
import {Customer} from "@/models/customer";

export class Contract {
    constructor(
        public id: number = 0,
        public contract_number: string = '',
        public start_trip_date: string = '',
        public end_trip_date: string = '',
        public agreement: Agreement = new Agreement(),
        public agent: Agent = new Agent(),
        public customer: Customer = new Customer(),
        public organization: Organization = new Organization(),
    ) {
    }
}