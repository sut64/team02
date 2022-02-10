package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"

	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("se-64.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Member{},
		&BookInformation{},
		&BookOrder{},
		&BookType{},
		&BookLocation{},
		&Company{},
		&OrderStatus{},
		&BookingRoom{},
		&Member{},
		&DeviceList{},
		&DeviceType{},
		&DeviceBorrow{},
		&BookReturn{},
		&RoomType{},
		&RoomAndTime{},
		&RoomObjective{},
		&BorrowDetail{},
		&BookCategory{},
		&TypeResearch{},
		&AuthorName{},
		&InstitutionName{},
		&Research{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	Ploy := Member{
		Name:     "พลอยชมพู วงศ์ฉันท์ทัต",
		Email:    "p@gmail.com",
		Password: string(password),
	}
	db.Model(&Member{}).Create(&Ploy)

	Auy := Member{
		Name:     "บุญฑิตา ปวงสันเทียะ",
		Email:    "b@gmail.com",
		Password: string(password),
	}
	db.Model(&Member{}).Create(&Auy)

	//--BookType Data
	documentary := BookType{
		Type: "สารคดี",
	}
	db.Model(&BookType{}).Create(&documentary)

	fiction := BookType{
		Type: "บันเทิงคดี",
	}
	db.Model(&BookType{}).Create(&fiction)

	publication := BookType{
		Type: "สิ่งพิมพ์",
	}
	db.Model(&BookType{}).Create(&publication)

	textbook := BookType{
		Type: "ตำรา",
	}
	db.Model(&BookType{}).Create(&textbook)

	journal := BookType{
		Type: "วารสาร",
	}
	db.Model(&BookType{}).Create(&journal)

	magazine := BookType{
		Type: "นิตยสาร",
	}
	db.Model(&BookType{}).Create(&magazine)

	//--BookCategory Data
	Category000 := BookCategory{
		Category: "หมวด  000  เบ็ดเตล็ด  ความรู้ทั่วไป  บรรณารักษศาสตร์",
	}
	db.Model(&BookCategory{}).Create(&Category000)

	Category100 := BookCategory{
		Category: " หมวด  100  ปรัชญา  จิตวิทยา",
	}
	db.Model(&BookCategory{}).Create(&Category100)

	Category200 := BookCategory{
		Category: "หมวด  200  ศาสนา",
	}
	db.Model(&BookCategory{}).Create(&Category200)

	Category300 := BookCategory{
		Category: "หมวด  300  สังคมศาสตร์",
	}
	db.Model(&BookCategory{}).Create(&Category300)

	Category400 := BookCategory{
		Category: " หมวด  400  ภาษาศาสตร์",
	}
	db.Model(&BookCategory{}).Create(&Category400)

	Category500 := BookCategory{
		Category: "หมวด  500  วิทยาศาสตร์  คณิตศาสตร์",
	}
	db.Model(&BookCategory{}).Create(&Category500)

	Category600 := BookCategory{
		Category: "หมวด  600  เทคโนโลยี  หรือวิทยาศาสตร์ประยุกต์",
	}
	db.Model(&BookCategory{}).Create(&Category600)

	Category700 := BookCategory{
		Category: " หมวด  700  ศิลปกรรม  และนันทนาการ",
	}
	db.Model(&BookCategory{}).Create(&Category700)

	Category800 := BookCategory{
		Category: "หมวด  800  วรรณคดี",
	}
	db.Model(&BookCategory{}).Create(&Category800)

	Category900 := BookCategory{
		Category: "หมวด  900  ภูมิศาสตร์และประวัติศาสตร์",
	}
	db.Model(&BookCategory{}).Create(&Category900)

	//--BookLocation Data
	floor1 := BookLocation{
		Location: "ชั้น 1",
	}
	db.Model(&BookLocation{}).Create(&floor1)

	floor2 := BookLocation{
		Location: "ชั้น 2",
	}
	db.Model(&BookLocation{}).Create(&floor2)

	floor3 := BookLocation{
		Location: "ชั้น 3",
	}
	db.Model(&BookLocation{}).Create(&floor3)

	done := OrderStatus{
		Status: "สั่งซื้อสำเร็จ",
	}
	db.Model(&OrderStatus{}).Create(&done)

	request := OrderStatus{
		Status: "เสนอสั่งซื้อ",
	}
	db.Model(&OrderStatus{}).Create(&request)

	approval := OrderStatus{
		Status: "อนุมัติ",
	}
	db.Model(&OrderStatus{}).Create(&approval)

	A := Company{
		NameThai:    "บริษัท เอ จำกัด",
		NameEng:     "A Company Limited",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "a4586sa@mail.com",
		Website:     "A Co.Ltd.com",
	}
	db.Model(&Company{}).Create(&A)

	B := Company{
		NameThai:    "บริษัท สมชายการพิมพ์ จำกัด",
		NameEng:     "Somchai Company Limited",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "somchai_typing@mail.com",
		Website:     "SomChai.com",
	}
	db.Model(&Company{}).Create(&B)

	C := Company{
		NameThai:    "สำนักพิมพ์ดีต่อใจ",
		NameEng:     "deetorjaibooks",
		Address:     "12 ถนน สาม ตำบอลสังคม อำเภอเมือง จังหวัดสมุทรปราการ 10270",
		PhoneNumber: "0123456789",
		Email:       "deetorjaibooks@mail.com",
		Website:     "deetorjaibooks.com",
	}
	db.Model(&Company{}).Create(&C)

	order1 := BookOrder{
		BookTitle:   "คณิตศาสตร์",
		Author:      "หญิงสาม ใจดี",
		BookType:    fiction,
		Company:     B,
		OrderAmount: 5,
		Price:       2052.50,
		OrderStatus: request,
		OrderDate: time.Now(),
	}
	db.Model(&BookOrder{}).Create(&order1)

	order2 := BookOrder{
		BookTitle:   "ชีวิตของสิงโต",
		Author:      "สมาสองหนึ่ง",
		BookType:    documentary,
		Company:     C,
		OrderAmount: 12,
		Price:       5200.25,
		OrderStatus: done,
		OrderDate: time.Now(),
	}
	db.Model(&BookOrder{}).Create(&order2)

	db.Model(&Member{}).Create(&Member{
		Name:     "ณัฐรินทร์ เนื้อทอง",
		Email:    "nattarin@gmail.com",
		Password: string(password),
	})
	db.Model(&Member{}).Create(&Member{
		Name:     "บุญฑิตา ปวงสันเทียะ",
		Email:    "boontita@gmail.com",
		Password: string(password),
	})

	var natarin Member
	db.Raw("SELECT * FROM members WHERE name = ?", "ณัฐรินทร์ เนื้อทอง").Scan(&natarin)

	// DeviceList Data
	code1 := DeviceList{
		DeviceCode: "D00000",
	}
	db.Model(&DeviceList{}).Create(&code1)

	// DeviceList Data
	code2 := DeviceList{
		DeviceCode: "D00001",
	}
	db.Model(&DeviceList{}).Create(&code2)

	// DeviceType Data
	type1 := DeviceType{
		Type: "อุปกรณ์ไฟฟ้า",
	}
	db.Model(&DeviceType{}).Create(&type1)

	// DeviceType Data
	type2 := DeviceType{
		Type: "อุปกรณ์อิเล็กทรอนิกส์",
	}
	db.Model(&DeviceType{}).Create(&type2)

	// DeviceBorrow 1
	db.Model(&DeviceBorrow{}).Create(&DeviceBorrow{
		DeviceName: "ปลั๊กไฟ",
		DeviceList: code1,
		DeviceType: type1,
		BorrowCode: "BD0000",
		Amount:     2,
		Date:       time.Now(),
		Member:     natarin,
	})

	Type1 := RoomType{
		Name: "ห้องเดี่ยว",
	}
	db.Model(&RoomType{}).Create(&Type1)

	Type2 := RoomType{
		Name: "ห้องกลุ่ม",
	}
	db.Model(&RoomType{}).Create(&Type2)

	Objective1 := RoomObjective{
		Name: "อ่านหนังสือ",
	}
	db.Model(&RoomObjective{}).Create(&Objective1)

	Objective2 := RoomObjective{
		Name: "ทำกิจกรรม",
	}
	db.Model(&RoomObjective{}).Create(&Objective2)

	Objective3 := RoomObjective{
		Name: "ติวหนังสือเป็นกลุ่ม",
	}
	db.Model(&RoomObjective{}).Create(&Objective3)

	Objective4 := RoomObjective{
		Name: "ทำ Project",
	}
	db.Model(&RoomObjective{}).Create(&Objective4)

	Objective5 := RoomObjective{
		Name: "อื่นๆ",
	}
	db.Model(&RoomObjective{}).Create(&Objective5)

	RoomAndTime1 := RoomAndTime{
		Name: "17:00-18:00",
	}
	db.Model(&RoomAndTime{}).Create(&RoomAndTime1)

	RoomAndTime2 := RoomAndTime{
		Name: "18:00-19:00",
	}
	db.Model(&RoomAndTime{}).Create(&RoomAndTime2)

	RoomAndTime3 := RoomAndTime{
		Name: "19:00-20:00",
	}
	db.Model(&RoomAndTime{}).Create(&RoomAndTime3)

	RoomAndTime4 := RoomAndTime{
		Name: "20:00-21:00",
	}
	db.Model(&RoomAndTime{}).Create(&RoomAndTime4)

	Borrowed := Status{
		Name: "Borrowed",
	}
	db.Model(&Status{}).Create(&Borrowed)

	Returned := Status{
		Name: "Returned",
	}
	db.Model(&Status{}).Create(&Returned)

	Library := ServicePlace{
		Name: "บรรณาสาร",
	}
	db.Model(&ServicePlace{}).Create((&Library))

	// TypeResearch data
	t1 := TypeResearch{
		Value: "แบ่งตามจุดมุ่งหมายของการวิจัย",
	}
	db.Model(&TypeResearch{}).Create(&t1)

	t2 := TypeResearch{
		Value: "แบ่งตามประโยชน์ของการวิจัย",
	}
	db.Model(&TypeResearch{}).Create(&t2)

	t3 := TypeResearch{
		Value: "แบ่งตามวิธีการเก็บรวบรวมข้อมูล",
	}
	db.Model(&TypeResearch{}).Create(&t3)

	t4 := TypeResearch{
		Value: "แบ่งตามลักษณะการวิเคราะห์ข้อมูล",
	}
	db.Model(&TypeResearch{}).Create(&t4)

	t5 := TypeResearch{
		Value: "แบ่งตามลักษณะวิชาหรือศาสตร์",
	}
	db.Model(&TypeResearch{}).Create(&t5)

	t6 := TypeResearch{
		Value: "แบ่งตามระเบียบวิธีวิจัย",
	}
	db.Model(&TypeResearch{}).Create(&t6)

	// AuthorName data
	AuthorName1 := AuthorName{
		AuthorName: "Chatchai Asara",
	}
	db.Model(&AuthorName{}).Create(&AuthorName1)

	AuthorName2 := AuthorName{
		AuthorName: "Edward Lee",
	}
	db.Model(&AuthorName{}).Create(&AuthorName2)

	AuthorName3 := AuthorName{
		AuthorName: "วิศนุ ทองหลาง",
	}
	db.Model(&AuthorName{}).Create(&AuthorName3)

	// InstitutionName data
	InstitutionName1 := InstitutionName{
		InstitutionName: "มหาวิทยาลัยเทคโนโลยีสุรนารี",
	}
	db.Model(&InstitutionName{}).Create(&InstitutionName1)

	InstitutionName2 := InstitutionName{
		InstitutionName: "มหาวิทยาลัยเทคโนโลยีราชมงคลอีสาน",
	}
	db.Model(&InstitutionName{}).Create(&InstitutionName2)

	InstitutionName3 := InstitutionName{
		InstitutionName: "มหาวิทยาลัยเชียงใหม่",
	}
	db.Model(&InstitutionName{}).Create(&InstitutionName3)

	// Research data
	research1 := Research{
		NameResearch:      "ResearchInDekThai2019",
		YearOfPublication: 2019,
		TypeResearchID:    &t1.ID,
		AuthorNameID:      &AuthorName3.ID,
		InstitutionNameID: &InstitutionName3.ID,
		RecordingDate:     time.Now(),
	}
	db.Model(&Research{}).Create(&research1)

	research2 := Research{
		NameResearch:      "ResearchMCPU",
		YearOfPublication: 2017,
		TypeResearchID:    &t5.ID,
		AuthorNameID:      &AuthorName2.ID,
		InstitutionNameID: &InstitutionName1.ID,
		RecordingDate:     time.Now(),
	}
	db.Model(&Research{}).Create(&research2)
}
