package inputs

import (
	"errors"
	"fmt"
	"github.com/golang-interfaces/ios"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"io/ioutil"
)

var _ = Describe("Validate", func() {
	Context("invoked w/ non-nil param.File", func() {
		Context("value.File is empty", func() {
			It("should return expected errors", func() {

				/* arrange */
				providedValue := &model.Value{}
				providedParam := &model.Param{
					File: &model.FileParam{},
				}

				expectedErrors := []error{
					errors.New("File required"),
				}

				objectUnderTest := newValidator()

				/* act */
				actualErrors := objectUnderTest.Validate(providedValue, providedParam)

				/* assert */
				Expect(actualErrors).To(Equal(expectedErrors))

			})
		})
		Context("value nil", func() {
			It("should return expected errors", func() {

				/* arrange */
				providedParam := &model.Param{
					File: &model.FileParam{},
				}

				expectedErrors := []error{
					errors.New("File required"),
				}

				objectUnderTest := newValidator()

				/* act */
				actualErrors := objectUnderTest.Validate(nil, providedParam)

				/* assert */
				Expect(actualErrors).To(Equal(expectedErrors))

			})
		})
		Context("value.File isn't empty", func() {
			It("should call fs.Stat w/ expected args", func() {

				/* arrange */
				providedValueFile := "dummyFile"
				providedValue := &model.Value{
					File: &providedValueFile,
				}
				providedParam := &model.Param{
					File: &model.FileParam{},
				}

				fakeOS := new(ios.Fake)
				// error to trigger immediate return
				fakeOS.StatReturns(nil, errors.New("dummyError"))

				objectUnderTest := _validator{
					os: fakeOS,
				}

				/* act */
				objectUnderTest.Validate(providedValue, providedParam)

				/* assert */
				Expect(fakeOS.StatArgsForCall(0)).To(Equal(*providedValue.File))

			})
			Context("fs.Stat errors", func() {
				It("should return expected errors", func() {

					/* arrange */
					providedValueFile := "dummyFile"
					providedValue := &model.Value{
						File: &providedValueFile,
					}
					providedParam := &model.Param{
						File: &model.FileParam{},
					}

					expectedErrors := []error{
						errors.New("dummyError"),
					}

					fakeOS := new(ios.Fake)
					fakeOS.StatReturns(nil, expectedErrors[0])

					objectUnderTest := _validator{
						os: fakeOS,
					}

					/* act */
					actualErrors := objectUnderTest.Validate(providedValue, providedParam)

					/* assert */
					Expect(actualErrors).To(Equal(expectedErrors))

				})

			})
			Context("fs.Stat doesn't error", func() {
				Context("FileInfo.IsDir returns false", func() {
					It("should return no errors", func() {

						/* arrange */
						// no good way to fake fileinfo
						tmpFile, err := ioutil.TempFile("", "")
						if nil != err {
							panic(err)
						}

						tmpFilePath := tmpFile.Name()

						providedValue := &model.Value{
							File: &tmpFilePath,
						}
						providedParam := &model.Param{
							File: &model.FileParam{},
						}

						expectedErrors := []error{}

						objectUnderTest := newValidator()

						/* act */
						actualErrors := objectUnderTest.Validate(providedValue, providedParam)

						/* assert */
						Expect(actualErrors).To(Equal(expectedErrors))

					})
				})
				Context("FileInfo.IsDir returns true", func() {
					It("should return expected errors", func() {

						/* arrange */
						// no good way to fake fileinfo
						tmpDirPath, err := ioutil.TempDir("", "")
						if nil != err {
							panic(err)
						}

						providedValue := &model.Value{
							File: &tmpDirPath,
						}
						providedParam := &model.Param{
							File: &model.FileParam{},
						}

						expectedErrors := []error{
							fmt.Errorf("%v not a file", tmpDirPath),
						}

						objectUnderTest := newValidator()

						/* act */
						actualErrors := objectUnderTest.Validate(providedValue, providedParam)

						/* assert */
						Expect(actualErrors).To(Equal(expectedErrors))

					})
				})
			})
		})
	})

})
