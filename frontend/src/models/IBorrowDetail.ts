import { ServicePlacesInterface } from "./IServicePlace";
import { BookOrderInterface } from "./IBookOrder";
import { MembersInterface } from "./IMember";
import { StatusInterface } from "./IStatus";


export interface BorrowDetailInterface {
    ID: number,

    MemberID:   number,
    Member:     MembersInterface,

    BookOrderID:     number,
    BookOrder:       BookOrderInterface,

    ServicePlaceID:    number,
    SevicePlace:      ServicePlacesInterface,

    StatusID:   number,
    Status:     StatusInterface,

    DateToBorrow: Date,

    BorrowDuration: number,
    
    Tel:        string,
}