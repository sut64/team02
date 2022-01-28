import { ServicePlacesInterface } from "./IPlace";
import { InfoInterface } from "./IInfo";
import { MembersInterface } from "./IMember";
import { StatusInterface } from "./IStatus";

export interface BorrowDetailInterface {
    ID: number,
    MemberID:   number,
    Member:     MembersInterface,
    InfoID:     number,
    Info:       InfoInterface,
    PlaceID:    number,
    Place:      ServicePlacesInterface,
    StatusID:   number,
    Status:     StatusInterface,
    DateToBorrow: Date,
    BorrowDuration: number,
    Tel:        string,
}