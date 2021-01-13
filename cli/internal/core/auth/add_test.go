package auth

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	nodeFakes "github.com/opctl/opctl/sdks/go/node/fakes"
)

var _ = Context("Adder", func() {
	Context("Invoke", func() {
		It("should call apiClient.Invoke w/ expected args", func() {
			/* arrange */
			fakeCore := new(nodeFakes.FakeOpNode)

			providedCtx := context.TODO()

			expectedCtx := providedCtx
			expectedReq := model.AddAuthReq{
				Resources: "Resources",
				Creds: model.Creds{
					Username: "username",
					Password: "password",
				},
			}

			objectUnderTest := _adder{
				core: fakeCore,
			}

			/* act */
			err := objectUnderTest.Add(
				expectedCtx,
				expectedReq.Resources,
				expectedReq.Username,
				expectedReq.Password,
			)

			/* assert */
			actualCtx, actualReq := fakeCore.AddAuthArgsForCall(0)
			Expect(err).To(BeNil())
			Expect(actualCtx).To(Equal(expectedCtx))
			Expect(actualReq).To(BeEquivalentTo(expectedReq))
		})
		Context("apiClient.Invoke errors", func() {
			It("should return expected error", func() {
				/* arrange */
				fakeCore := new(nodeFakes.FakeOpNode)
				expectedError := errors.New("dummyError")
				fakeCore.AddAuthReturns(expectedError)

				objectUnderTest := _adder{
					core: fakeCore,
				}

				/* act */
				err := objectUnderTest.Add(context.TODO(), "", "", "")

				/* assert */
				Expect(err).To(MatchError(expectedError))
			})
		})
	})
})
