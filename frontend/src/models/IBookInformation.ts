import { BookOrderInterface } from "./IBookOrder";
import { BookTypeInterface } from "./IBookType";
import { BookLocationInterface } from "./IBookLocation";

export interface BookInformationInterface {
  ID: number,
  Date: Date,
  YearPublication: number,
  CallNumber: String,

  BookOrderID: number,
  BookOrder: BookOrderInterface,

  BookTypeID: number,
  BookType:  BookTypeInterface,

  BookLocationID: number,
  BookLocation: BookLocationInterface,
  
}