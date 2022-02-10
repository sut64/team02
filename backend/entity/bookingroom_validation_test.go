package entity

import (
	"testing"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookingRoomPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	bookingroom := BookingRoom{
		PhoneBooker: "0946027710",
		QuantityMember: 9,
		BookingRoomAt: time.Now().Add(time.Hour*24),
		
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bookingroom)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

 func TestQuantityMemberOneToTen(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []uint{
		12,15,	
	}

	for _, fixture := range fixtures {

	bookingroom := BookingRoom{
		PhoneBooker: "0946027710",
		QuantityMember: uint(fixture),
		BookingRoomAt: time.Now().Add(time.Hour*24),
	}

	ok, err := govalidator.ValidateStruct(bookingroom)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal(fmt.Sprintf(`QuantityMember: %d does not validate as range(1|10)`, fixture)))
	}
} 


//  ตรวจสอบเบอร์โทรแล้วต้องเจอ Error
func TestPhoneBookerMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{

		"1234567890",
		"X123456789", //ขึ้นต้นด้วยตัวอักษร \d 9 ตัว
		"1111111111", //ขึ้นต้นเลขอื่นที่ไม่ใช่ 0 \d 9 ตัว
		"094602771",  //ขึ้นต้นด้วย 0  \d 8 ตัว

	}

	for _, fixture := range fixtures {
		bookingroom := BookingRoom{
			PhoneBooker: fixture,
			QuantityMember: 9,
			BookingRoomAt: time.Now().Add(24*time.Hour),
		}

		ok, err := govalidator.ValidateStruct(bookingroom)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf(`PhoneBooker: %s does not validate as matches(^[0]\d{9}$)`, fixture)))
	}

}



func TestBookingAtMustBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

		bookingroom := BookingRoom{
			PhoneBooker: "0946027710",
			QuantityMember: 9,
			BookingRoomAt: time.Now().Add(24 / time.Hour),
		}

		ok, err := govalidator.ValidateStruct(bookingroom)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("BookingRoomAt must not be in the past"))

}
