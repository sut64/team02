import { BookTypeInterface } from "./IBookType";
import { CompaniesInterface } from "./ICompany";
import { OrderStatusesInterface } from "./IOrderStatus";

export interface BookOrderInterface {
    ID: number,
    BookTitle: string;
    Author: string;
    OrderAmount: number;
    Price: Float32Array;
    OrderDate: Date;
    BookTypeID: number,
    BookType: BookTypeInterface,
    CompanyID: number,
    Company: CompaniesInterface,
    OrderStatusID: number,
    OrderStatus: OrderStatusesInterface
   }
   