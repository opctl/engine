package colorer

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "fmt"
)

var _ = Describe("colorer", func() {
  Context("New", func() {
    It("should return Colorer", func() {
      /* arrange/act/assert */
      Expect(New()).Should(Not(BeNil()))
    })
  })
  objectUnderTest := New()
  Context("Attention", func() {
    It("should return expected string", func() {
      /* arrange */
      providedFormatString := "%v"
      providedValue1 := "dummyString"
      expectedString := fmt.Sprintf("\x1b[93;1m%s\x1b[0m", fmt.Sprintf(providedFormatString, providedValue1))

      /* act */
      actualString := objectUnderTest.Attention(providedFormatString, providedValue1)

      /* assert */
      Expect(fmt.Sprintf("%+q", actualString)).To(Equal(fmt.Sprintf("%+q", expectedString)))
    })
  })
  Context("Error", func() {
    It("should return expected string", func() {
      /* arrange */
      providedFormatString := "%v"
      providedValue1 := "dummyString"
      expectedString := fmt.Sprintf("\x1b[91;1m%s\x1b[0m", fmt.Sprintf(providedFormatString, providedValue1))

      /* act */
      actualString := objectUnderTest.Error(providedFormatString, providedValue1)

      /* assert */
      Expect(fmt.Sprintf("%+q", actualString)).To(Equal(fmt.Sprintf("%+q", expectedString)))
    })
  })
  Context("Info", func() {
    It("should return expected string", func() {
      /* arrange */
      providedFormatString := "%v"
      providedValue1 := "dummyString"
      expectedString := fmt.Sprintf("\x1b[96;1m%s\x1b[0m", fmt.Sprintf(providedFormatString, providedValue1))

      /* act */
      actualString := objectUnderTest.Info(providedFormatString, providedValue1)

      /* assert */
      Expect(fmt.Sprintf("%+q", actualString)).To(Equal(fmt.Sprintf("%+q", expectedString)))
    })
  })
  Context("Success", func() {
    It("should return expected string", func() {
      /* arrange */
      providedFormatString := "%v"
      providedValue1 := "dummyString"
      expectedString := fmt.Sprintf("\x1b[92;1m%s\x1b[0m", fmt.Sprintf(providedFormatString, providedValue1))

      /* act */
      actualString := objectUnderTest.Success(providedFormatString, providedValue1)

      /* assert */
      Expect(fmt.Sprintf("%+q", actualString)).To(Equal(fmt.Sprintf("%+q", expectedString)))
    })
  })
})