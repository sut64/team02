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


}