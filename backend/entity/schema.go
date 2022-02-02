package entity

import (
	"gorm.io/gorm"

	"time"
)

type Member struct {
	gorm.Model

	Name string

	Email string `gorm:"uniqueIndex"`

	Password string

	BookInformation []BookInformation `gorm:"foreignKey:MemberID"`

	BookOrder []BookOrder `gorm:"foreignKey:MemberID"`

	BorrowDetails []BorrowDetail `gorm:"foreignKey:MemberID"`

	BookReturns []BookReturn `gorm:"foreinkey:MemberID"`
}

type BookType struct {
	gorm.Model

	Type string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookTypeID"`

	BookOrder []BookOrder `gorm:"foreignKey:BookTypeID"`
}

type BookCategory struct {
	gorm.Model

	Category string 						`gorm:"uniqueIndex"`

	BookInformation []BookInformation 		`gorm:"foreignKey:BookCategoryID"`
}

type BookLocation struct {
	gorm.Model

	Location string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookLocationID"`
}

type BookInformation struct {
	gorm.Model

	Date            time.Time
	CallNumber      string
	YearPublication uint
	// BookOrderID ทำหน้าที่เป็น FK
	BookOrderID *uint
	BookOrder   BookOrder `valid:"-"`

	// BookLocationID ทำหน้าที่เป็น FK
	BookLocationID *uint
	BookLocation   BookLocation `valid:"-"`

	// BookTypeID ทำหน้าที่เป็น FK
	BookTypeID *uint
	BookType   BookType `valid:"-"`

	// MemberID ทำหน้าที่เป็น FK
	MemberID *uint
	Member   Member `valid:"-"`
}

type Company struct {
	gorm.Model
	NameThai    string
	NameEng     string
	Address     string
	PhoneNumber string
	Email       string
	Website     string

	BookOrder []BookOrder `gorm:"foreignKey:CompanyID"`
}

type OrderStatus struct {
	gorm.Model
	Status string

	BookOrder []BookOrder `gorm:"foreignKey:OrderStatusID"`
}

type BookOrder struct {
	gorm.Model
	BookTitle   string
	Author      string
	OrderAmount uint
	Price       float32
	OrderDate   time.Time

	//Company ทำหน้าที่เป็น FK
	CompanyID *uint
	Company   Company

	//Company ทำหน้าที่เป็น FK
	OrderStatusID *uint
	OrderStatus   OrderStatus

	//BookType ทำหน้าที่เป็น FK
	BookTypeID *uint
	BookType   BookType

	//Librarian
	MemberID *uint
	Member   Member

	BorrowDetail []BorrowDetail `gorm:"foreignKey:BookOrderID"`
}

type BorrowDetail struct {
	gorm.Model

	DateToBorrow   time.Time //`valid:"future~DateToBorrow must be in the future"`
	Tel            string    //`valid:"matches(^[0]{1}[0-9]{9})"`
	BorrowDuration uint      // `valid:"range(1|30)"`

	MemberID *uint
	Member   Member `gorm:"references:id"`

	ServicePlaceID *uint
	ServicePlace   ServicePlace `gorm:"references:id"`

	BookOrderID *uint
	BookOrder   BookOrder `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	BookReturns []BookReturn `gorm:"foreignKey:BorrowDetailID"`
}

type ServicePlace struct {
	gorm.Model
	Name          string
	BorrowDetails []BorrowDetail `gorm:"foreignKey:ServicePlaceID"`
	BookReturns   []BookReturn   `gorm:"foreinkey:ServicePlaceID"`
}

type Status struct {
	gorm.Model
	Name          string
	BorrowDetails []BorrowDetail `gorm:"foreignKey:StatusID"`
	BookReturns   []BookReturn   `gorm:"foreinkey:StatusID"`
}
type DeviceList struct {
	gorm.Model
	DeviceCode   string
	DeviceBorrow []DeviceBorrow `gorm:"foreignKey:DeviceListID"`
}

type DeviceType struct {
	gorm.Model
	Type         string
	DeviceBorrow []DeviceBorrow `gorm:"foreignKey:DeviceTypeID"`
}

type DeviceBorrow struct {
	gorm.Model
	DeviceName string `gorm:"not null"`
	BorrowCode string
	//`valid:"matches(^[B]+[D]\\d{4}$)~BorrowCode: %s does not validate as matches"`
	Amount int
	//`valid:"range(0|9)~Amount must be in negative"`
	Date time.Time
	//	`valid:"present~Date must be in the present"`

	//MemberID ทำหน้าที่เป็น FK
	MemberID *uint
	Member   Member

	//DeviceListID ทำหน้าที่เป็น FK
	DeviceListID *uint
	DeviceList   DeviceList

	//DeviceTypeID ทำหน้าที่เป็น FK
	DeviceTypeID *uint
	DeviceType   DeviceType
}

type BookReturn struct {
	gorm.Model
	DateReturn time.Time
	Damage     int
	Tel        string

	MemberID *uint
	Member   Member `gorm:"references:id" `

	BorrowDetailID *uint
	BorrowDetail   BorrowDetail `gorm:"references:id"`

	ServicePlaceID *uint
	ServicePlace   ServicePlace `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`
}

type BookingRoom struct {
	gorm.Model

	MemberID *uint
	Member   Member `gorm:"references:ID"`

	RoomAndTimeID *uint
	RoomAndTime   RoomAndTime `gorm:"references:ID"`

	RoomTypeID *uint
	RoomType   RoomType `gorm:"references:ID"`

	RoomObjectiveID *uint
	RoomObjective   RoomObjective `gorm:"references:ID"`

	QuantityMember uint      // `valid:"range(1|10),required"`
	PhoneBooker    string    //`valid:"matches(^[0]\\d{9}$)"`
	BookingRoomAt  time.Time //`valid:"future~BookingRoomAt must not be in the past"`

}

type RoomAndTime struct {
	gorm.Model

	Name string

	BookingRoom []BookingRoom `gorm:"foreignkey:RoomAndTimeID"`
}

type RoomType struct {
	gorm.Model

	Name string

	BookingRoom []BookingRoom `gorm:"foreignkey:RoomTypeID"`
}

type RoomObjective struct {
	gorm.Model

	Name string

	BookingRoom []BookingRoom `gorm:"foreignkey:RoomObjectiveID"`
}
