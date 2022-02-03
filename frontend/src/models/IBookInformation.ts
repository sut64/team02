import { BookOrderInterface } from "./IBookOrder";
import { BookLocationInterface } from "./IBookLocation";
import { BookCategoryInterface } from "./IBookCategory";
import { MembersInterface } from "./IMember";

export interface BookInformationInterface {
  ID: number,
  Date: Date,
  YearPublication: number,
  CallNumber: String,

  BookOrderID: number,
  BookOrder: BookOrderInterface,

  BookCategoryID: number,
  BookCategory:  BookCategoryInterface,

  BookLocationID: number,
  BookLocation: BookLocationInterface,

  MemberID: number,
  Member: MembersInterface,
  
}