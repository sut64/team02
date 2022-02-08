package entity

import (
	"testing"
	//"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookInformationPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องทุก field
	bookInformation := BookInformation{
		Date:	time.Now(),
		CallNumber:	"BI.100",
		YearPublication: 2000,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bookInformation)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}
func TestDateMustBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลวันที่ผิด
	bookInformation := BookInformation{
		Date:            time.Now().Add(24 * time.Hour), // อนาคต, fail
		CallNumber:      "BI.100",
		YearPublication: 2022,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bookInformation)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Date must be in the now"))
}

func TestYearPublicationMustbetween(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลปีที่พิมพ์ผิด
	bookInformation := BookInformation{
		Date:            time.Now(),
		CallNumber:      "BI.100",
		YearPublication: 1898,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bookInformation)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Year Publication must be between 1900 - 2022"))
}

func TestCallNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"B.456",
		"BI.90",  
		"BO456",   
		"AZ*456", 
	}
	// ข้อมูล callnnumber ผิด
	for _, fixture := range fixtures {
		bookInformation := BookInformation{
			Date:            time.Now(),
			CallNumber:      fixture, //ผิด
			YearPublication: 2022,
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bookInformation)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal((`CallNumber: does not validate as matches`)))
}
}