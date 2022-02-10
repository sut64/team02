import { BorrowDetailInterface } from "./IBorrowDetail"
import { MembersInterface } from "./IMember"
import { ServicePlacesInterface } from "./IServicePlace"
import { StatusInterface } from "./IStatus"

export interface BookReturnsInterface{
   
    ID: Number,
    Damage: number,
    Tel: string,
    DateReturn: Date,
    MemberID: string,
    Member: MembersInterface,
    BorrowDetailID: number,
    BorrowDetail: BorrowDetailInterface
    ServicePlaceID: number,
    ServicePlace: ServicePlacesInterface,
    StatusID: number,
    Status: StatusInterface,


    
}