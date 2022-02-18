package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookReturn(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	br := BookReturn{
		Damage:     5,
		Tel:        "0123467891",
		DateReturn: time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(br)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestDamagenNotOverLimit(t *testing.T) {
	g := NewGomegaWithT(t)

	br := BookReturn{
		Damage:     120,
		Tel:        "0123467891",
		DateReturn: time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(br)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Damage: 120 does not validate as range(1|50)"))
}

func TestDateReturnMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	wv := BookReturn{
		DateReturn: time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC), // อนาคต, fail
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(wv)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("DateReturn: must be in the now"))
}

func TestTelMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"1002578963",  //ตัวเเรกไม่ขึ้นต้นด้วย 0
		"050000",      //ไม่ครบจำนวน
		"01234567897", //เกินจำนวน

	}

	for _, fixture := range fixtures {
		br := BookReturn{
			Damage:     0,
			Tel:        fixture, //ผิด
			DateReturn: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(br)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Tel not match"))
	}
}
