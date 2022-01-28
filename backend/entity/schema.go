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
}

type BookType struct {
	gorm.Model

	Type string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookTypeID"`

	BookOrder []BookOrder `gorm:"foreignKey:BookTypeID"`
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
}

type BorrowDetail struct {
	gorm.Model

	DateToBorrow   time.Time //`valid:"future~DateToBorrow must be in the future"`
	Tel            string    //`valid:"matches(^[0]{1}[0-9]{9})"`
	BorrowDuration uint      // `valid:"range(1|30)"`

	MemberID *uint
	Member   Member `gorm:"references:id"`

	PlaceID *uint
	Place   ServicePlace `gorm:"references:id"`

	InfoID *uint
	Info   BookInformation `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`
}

type ServicePlace struct {
	gorm.Model
	Name          string
	BorrowDetails []BorrowDetail `gorm:"foreignKey:PlaceID"`
}

type Status struct {
	gorm.Model
	Name          string
	BorrowDetails []BorrowDetail `gorm:"foreignKey:StatusID"`
}
