package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestDeviceBorrowPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	deviceborrow := DeviceBorrow{
		BorrowCode: "BD0000",
		Amount:     2,
		Date:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(deviceborrow)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestBorrowCodeMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"XX0000",
		"BD00000", // BD ตามด้วย \d 5 ตัว
		"BD000",   // BD ตามด้วย \d 3 ตัว
		"BA0000",  // BA ตามด้วย \d 4 ตัว
	}

	for _, fixture := range fixtures {
		deviceborrow := DeviceBorrow{
			BorrowCode: fixture, //ผิด
			Amount:     2,
			Date:       time.Now(),
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(deviceborrow)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf(`BorrowCode: %s does not validate as matches`, fixture)))
	}
}

//ตรวจสอบจำนวนห้ามติดลบ
func TestAmount(t *testing.T) {
	g := NewGomegaWithT(t)

	deviceborrow := DeviceBorrow{
		BorrowCode: "BD0000",
		Amount:     -2, //ข้อมูล amount ผิด
		Date:       time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(deviceborrow)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err.Error()).To(Equal("Amount must be in negative"))
}

//ตรวจสอบวันที่เป็นปัจจุบัน
func TestDateMustBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	deviceborrow := DeviceBorrow{

		Date: time.Now().Add(24 * time.Hour), // อนาคต, fail
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(deviceborrow)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Date must be in the present"))
}
