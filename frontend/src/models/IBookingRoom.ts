import { RoomAndTimeInterface } from "./IRoomAndTime";
import { RoomTypeInterface } from "./IRoomType";
import { RoomObjectiveInterface } from "./IRoomObjective";
import { MembersInterface } from "./IMember";

export interface BookingRoomInterface {

    ID:                 number,
    RoomAndTimeID:      number,
    RoomAndTime:        RoomAndTimeInterface,
    RoomTypeID:         number,
    RoomType:           RoomTypeInterface,      
    RoomObjectiveID:     number,
    RoomObjective:       RoomObjectiveInterface,
    PhoneBooker:        string,       
    QuantityMember:     number,
    BookingRoomAt:      Date,    
    MemberID:           number,  
    Member:             MembersInterface,  
    	
}