package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	Name string

	Email string `gorm:"uniqueIndex"`

	Password string

	Role string `gorm:"default:'สมาชิก'"`

	BookInformation []BookInformation `gorm:"foreignKey:MemberID"`

	BookOrder []BookOrder `gorm:"foreignKey:MemberID"`

	BorrowDetails []BorrowDetail `gorm:"foreignKey:MemberID"`

	BookReturns []BookReturn `gorm:"foreinkey:MemberID"`

	Researches []Research `gorm:"foreignKey:MemberID"`
}

type BookType struct {
	gorm.Model

	Type string `gorm:"uniqueIndex"`

	BookOrder []BookOrder `gorm:"foreignKey:BookTypeID"`
}

type BookCategory struct {
	gorm.Model

	Category string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookCategoryID"`
}

type BookLocation struct {
	gorm.Model

	Location string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookLocationID"`
}

type BookInformation struct {
	gorm.Model

	Date            time.Time `valid:"present~Date must be in the present"`
	CallNumber      string    `valid:"matches(^[ก-ฮ]+[.]+[ก-ฮ]+\\d{3}$|^[A-Z]+[.]+[A-Z]+\\d{3}$)~CallNumber: does not validate as matches"`
	YearPublication uint      `valid:"range(1900|2022)~Year Publication must be between 1900 - 2022"`
	// BookOrderID ทำหน้าที่เป็น FK
	BookOrderID *uint
	BookOrder   BookOrder `valid:"-"`

	// BookLocationID ทำหน้าที่เป็น FK
	BookLocationID *uint
	BookLocation   BookLocation `valid:"-"`

	// BookCategoryID ทำหน้าที่เป็น FK
	BookCategoryID *uint
	BookCategory   BookCategory `valid:"-"`

	// MemberID ทำหน้าที่เป็น FK
	MemberID *uint
	Member   Member `valid:"-"`
}

type Company struct {
	gorm.Model
	NameThai    string `gorm:"uniqueIndex"`
	NameEng     string
	Address     string
	PhoneNumber string
	Email       string
	Website     string

	BookOrder []BookOrder `gorm:"foreignKey:CompanyID"`
}

type OrderStatus struct {
	gorm.Model
	Status string `gorm:"uniqueIndex"`

	BookOrder []BookOrder `gorm:"foreignKey:OrderStatusID"`
}

type BookOrder struct {
	gorm.Model
	BookTitle   string
	Author      string
	OrderAmount int       `valid:"required~OrderAmount cannot be zero and negative,morethanzero~OrderAmount cannot be zero and negative"`
	Price       float32   `valid:"required~Price cannot be negative and zero,nonnegative~Price cannot be negative and zero"`
	OrderDate   time.Time `valid:"present~OrderDate must be present"`

	//Company ทำหน้าที่เป็น FK
	CompanyID *uint
	Company   Company `valid:"-"`

	//Company ทำหน้าที่เป็น FK
	OrderStatusID *uint
	OrderStatus   OrderStatus `valid:"-"`

	//BookType ทำหน้าที่เป็น FK
	BookTypeID *uint
	BookType   BookType `valid:"-"`

	//Librarian
	MemberID *uint
	Member   Member

	BorrowDetail    []BorrowDetail    `gorm:"foreignKey:BookOrderID"`
	BookInformation []BookInformation `gorm:"foreignKey:BookOrderID"`
}

type BorrowDetail struct {
	gorm.Model

	DateToBorrow   time.Time `valid:"notpast~DateToBorrow must be in the future and present"`
	Tel            string    `valid:"matches(^[0]{1}[0-9]{9}$)~Tel not match"`
	BorrowDuration uint      `valid:"range(1|30)~BorrowDuration must in range 1-30,required~BorrowDuration must in range 1-30"`

	MemberID *uint
	Member   Member `gorm:"references:id" valid:"-"`

	ServicePlaceID *uint
	ServicePlace   ServicePlace `gorm:"references:id" valid:"-"`

	BookOrderID *uint
	BookOrder   BookOrder `gorm:"references:id" valid:"-"`

	StatusID *uint
	Status   Status `gorm:"references:id" valid:"-"`

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
	DeviceName string    `gorm:"not null"`
	BorrowCode string    `valid:"matches(^[B]+[D]\\d{4}$)~BorrowCode: %s does not validate as matches"`
	Amount     int       `valid:"range(0|9)~Amount must be in negative"`
	Date       time.Time `valid:"present~Date must be in the present"`

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
	DateReturn time.Time `valid:"present~DateReturn: must be in the now"`
	Damage     int       `valid:"range(1|50)"`
	Tel        string    `valid:"matches(^[0]{1}[0-9]{9}$)~Tel not match"`

	MemberID *uint
	Member   Member `gorm:"references:id" valid:"-"`

	BorrowDetailID *uint
	BorrowDetail   BorrowDetail `gorm:"references:id" valid:"-"`

	ServicePlaceID *uint
	ServicePlace   ServicePlace `gorm:"references:id" valid:"-"`

	StatusID *uint
	Status   Status `gorm:"references:id" valid:"-"`
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

	QuantityMember uint      `valid:"range(1|10),required"`
	PhoneBooker    string    `valid:"matches(^[0]\\d{9}$)"`
	BookingRoomAt  time.Time `valid:"future~BookingRoomAt must not be in the past"`
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

type TypeResearch struct {
	gorm.Model
	Value      string
	Researches []Research `gorm:"foreignKey:TypeResearchID"`
}

type AuthorName struct {
	gorm.Model
	AuthorName string
	Researches []Research `gorm:"foreignKey:AuthorNameID"`
}

type InstitutionName struct {
	gorm.Model
	InstitutionName string
	Researches      []Research `gorm:"foreignKey:InstitutionNameID"`
}

type Research struct {
	gorm.Model

	NameResearch string `valid:"alphanum~NameResearch: does not validate as alphanum"`

	YearOfPublication uint `valid:"range(1900|2022),required"`

	RecordingDate time.Time `valid:"present~RecordingDate: must be in the now"`

	TypeResearchID *uint
	TypeResearch   TypeResearch `valid:"-"`

	AuthorNameID *uint
	AuthorName   AuthorName `valid:"-"`

	InstitutionNameID *uint
	InstitutionName   InstitutionName `valid:"-"`

	MemberID *uint
	Member   Member `valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.After(t)
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.Before(time.Time(t))
	})
	govalidator.CustomTypeTagMap.Set("present",
		func(i interface{}, context interface{}) bool {
			t := i.(time.Time)
			if t.Year() == time.Now().Year() {
				if int(t.Month()) == int(time.Now().Month()) {
					if t.Day() == time.Now().Day() {
						return true
					}
				}
			}
			return false
		})
	govalidator.CustomTypeTagMap.Set("notpast", func(i interface{}, o interface{}) bool {
		t := i.(time.Time)
		// ย้อนหลังไม่เกิน 1 วัน
		return t.After(time.Now().AddDate(0, 0, -1))
	})
	govalidator.CustomTypeTagMap.Set("nonnegative",
		func(i interface{}, context interface{}) bool {
			value := i.(float32)
			return value > 0
		})
	govalidator.CustomTypeTagMap.Set("morethanzero",
		func(i interface{}, context interface{}) bool {
			num := i
			return num.(int) > 0
		})
}
