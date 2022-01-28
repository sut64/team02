package entity

import (
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
  database, err := gorm.Open(sqlite.Open("bookorder.db"), &gorm.Config{})
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
    &BookOrder{}, &BookType{}, &Company{}, &OrderStatus{}, 
  )


  db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

  Ploy := Member{
    Name: "ploy",
		Email:   	"p@gmail.com",
		Password: 		string(password),
	}
	db.Model(&Member{}).Create(&Ploy)
  
	Auy := Member{
    Name: "auy",
		Email:   	"b@gmail.com",
		Password: 		string(password),
		
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

  //BookOrder Data
  maths := BookOrder{
    BookTitle: "คณิตศาสตร์",
  }
  db.Model(&BookOrder{}).Create(&maths)

  business := BookOrder{
    BookTitle: "ธุรกิจก้าวหน้า",
  }
  db.Model(&BookOrder{}).Create(&business)


	db.Model(&Member{}).Create(&Member{
    Name: "บุญฑิตา ปวงสันเทียะ",
    Email: "boontita@mail.com",
    Password: string(password),
	})

	done := OrderStatus{
		Status: "สั่งซื้อสำเร็จ",
	}
	db.Model(&OrderStatus{}).Create(&done)

	waitforapproval := OrderStatus{
		Status: "รอการอนุมัติ",
	}
	db.Model(&OrderStatus{}).Create(&waitforapproval)

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

	novel := BookType{
		Type: "นิยาย",
	}
	db.Model(&BookType{}).Create(&novel)

	documentary := BookType{
		Type: "สารคดี",
	}
	db.Model(&BookType{}).Create(&documentary)

	chidren := BookType{
		Type: "หนังสือสำหรับเด็กและเยาวชน",
	}
	db.Model(&BookType{}).Create(&chidren)

	treatise := BookType{
		Type: "ตำรา",
	}
	db.Model(&BookType{}).Create(&treatise)

	order1 := BookOrder{
		BookTitle:   "คณิตศาสตร์",
		Author:      "หญิงสาม ใจดี",
		BookType:    treatise,
		Company:     B,
		OrderAmount: 5,
		Price:       2052.50,
		OrderStatus: request,
	}
	db.Model(&BookOrder{}).Create(&order1)
}
}