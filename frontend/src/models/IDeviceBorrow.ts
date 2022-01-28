import { MembersInterface } from "./IMember";
import { DeviceListsInterface } from "./IDeviceList";
import { DeviceTypesInterface } from "./IDeviceType";

export interface DeviceBorrowInterface {
  ID:                 number,
  DeviceName:         string,

  BorrowCode:         string,

  Amount:             number,

  Date:               Date,

  MemberID:        number,
  Member:          MembersInterface,
  
  DeviceListID:       number,
  DeviceList:         DeviceListsInterface,
  
  DeviceTypeID:       number,
  
  DeviceType:         DeviceTypesInterface,
}
