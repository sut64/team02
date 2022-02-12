package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestDateToBorrowMustBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	bd := BorrowDetail{
		DateToBorrow:   time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
		Tel:            "0123456789",
		BorrowDuration: 5,
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bd)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("DateToBorrow must be in the future and present"))
}

func TestTelLengthAndCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"01234",
		"1234567890",
	}

	for _, fixture := range fixtures {
		bd := BorrowDetail{
			DateToBorrow:   time.Date(2025, 10, 5, 0, 0, 0, 0, time.UTC),
			Tel:            fixture,
			BorrowDuration: 5,
		}

		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(bd)

		//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Tel not match"))
	}
}

func TestBorrowDurationRange(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []uint{
		31,
		35,
	}

	for _, fixture := range fixtures {
		bd := BorrowDetail{
			DateToBorrow:   time.Date(2025, 10, 5, 0, 0, 0, 0, time.UTC),
			Tel:            "0123456789",
			BorrowDuration: fixture,
		}
		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(bd)

		//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf(`BorrowDuration: %d does not validate as range(1|30)`, fixture)))
	}
}

func TestBorrowDetailPass(t *testing.T) {
	g := NewGomegaWithT(t)

	//ข้อมูลถูกทั้งหมด
	bd := BorrowDetail{
		DateToBorrow:   time.Date(2025, 10, 5, 0, 0, 0, 0, time.UTC),
		Tel:            "0123456789",
		BorrowDuration: 5,
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bd)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())
}
