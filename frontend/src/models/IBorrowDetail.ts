import { ServicePlacesInterface } from "./IServicePlace";
import { BookInformationInterface } from "./IBookInformation";
import { MembersInterface } from "./IMember";
import { StatusInterface } from "./IStatus";

export interface BorrowDetailInterface {
    ID: number,
    MemberID:   number,
    Member:     MembersInterface,
    InfoID:     number,
    Info:       BookInformationInterface,
    ServicePlaceID:    number,
    SevicePlace:      ServicePlacesInterface,
    StatusID:   number,
    Status:     StatusInterface,
    DateToBorrow: Date,
    BorrowDuration: number,
    Tel:        string,
}