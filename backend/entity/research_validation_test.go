package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestResearchPass(t *testing.T) {
	g := NewGomegaWithT(t)

	rs := Research{
		NameResearch:      "ABCD01",
		YearOfPublication: 2020,
		RecordingDate:     time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rs)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestNameResearchMustBealphanum(t *testing.T) {
	g := NewGomegaWithT(t)

	rs := Research{
		NameResearch:      "A@BCD10",
		YearOfPublication: 2020,
		RecordingDate:     time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rs)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("NameResearch: does not validate as alphanum"))
}

func TestYearOfPublicationValidateAsRange(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []uint{
		2033,
		1100,
	}

	for _, fixture := range fixtures {
		rs := Research{
			NameResearch:      "ABCD01",
			YearOfPublication: fixture,
			RecordingDate:     time.Now(),
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(rs)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf(`YearOfPublication: %d does not validate as range(1900|2022)`, fixture)))
	}
}

func TestResearchMustBeNow(t *testing.T) {
	g := NewGomegaWithT(t)

	rs := Research{
		NameResearch:      "ABCD01",
		YearOfPublication: 2020,
		RecordingDate:     time.Now().Add(24 * time.Hour),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rs)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("RecordingDate: must be in the now"))
}
