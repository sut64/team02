package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

)

func TestOrderDateMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	bod := BookOrder{
		BookTitle:   "Conan",
		Author:      "Aoyama",
		OrderAmount: 15,
		Price:       1523.25,
		OrderDate:   time.Now().Add(24 * time.Hour), //ผิดเพราะเวลาเป็นอนาคต 

	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bod)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("OrderDate must be present"))
}

func TestOrderAmountCanNotBeZeoAndNegative(t *testing.T){
	g := NewGomegaWithT(t)

	bod := BookOrder{
		BookTitle:   "Conan",
		Author:      "Aoyama",
		OrderAmount: 0,			//ผิดเพราะห้ามเป็น 0
		Price:       1523.25,
		OrderDate:   time.Now(), 

	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bod)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("OrderAmount cannot be zero and negative"))
}

func TestPriceCanNotBeNegativeAndZero(t *testing.T){
	g := NewGomegaWithT(t)

	bod := BookOrder{
		BookTitle:   "Conan",
		Author:      "Aoyama",
		OrderAmount: 15,			
		Price:       -1500.50,	//ผิด เพราะเป็นค่าลบ
		OrderDate:   time.Now(), 
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bod)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Price cannot be negative and zero"))
}
 
func TestBookOrderPass(t *testing.T){
	g := NewGomegaWithT(t)

	//ข้อมูลถูกทั้งหมด
	bod := BookOrder{
		BookTitle:   "Conan",
		Author:      "Aoyama",
		OrderAmount: 15,			
		Price:       1522.25,	
		OrderDate:   time.Now(), 
	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bod)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())
}