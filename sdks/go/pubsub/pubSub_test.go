package pubsub

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("pubSub", func() {
	tempDirPath, err := ioutil.TempDir("", "")
	if nil != err {
		panic(err)
	}
	tempEventStore := NewBadgerDBEventStore(tempDirPath)

	Context("New", func() {
		It("should return PubSub", func() {
			/* arrange/act/assert */

			Expect(New(tempEventStore)).To(Not(BeNil()))
		})
	})
	Context("Publish", func() {
		Context("subscription exist", func() {
			Context("is subscribed", func() {
				It("receives event", func() {
					/* arrange */
					expectedEvent := model.Event{
						CallStarted: &model.CallStarted{
							RootCallID: "dummyRootCallID",
							CallID:     "dummyCallID",
							OpRef:      "dummyOpRef",
						},
					}

					objectUnderTest := New(tempEventStore)

					eventChannel, _ := objectUnderTest.Subscribe(context.TODO(), model.EventFilter{})

					/* act */
					objectUnderTest.Publish(expectedEvent)

					/* assert */
					var actualEvent model.Event
					Eventually(eventChannel).Should(Receive(&actualEvent))
					Expect(actualEvent).To(Equal(expectedEvent))
				})
			})
			Context("isn't subscribed", func() {
				It("doesn't receive event", func() {
					/* arrange */
					subscriberEventFilter := model.EventFilter{Roots: []string{"notPublishedRootCallID"}}

					publishedEvent := model.Event{
						CallStarted: &model.CallStarted{
							RootCallID: "dummyRootCallID",
							CallID:     "dummyCallID",
							OpRef:      "dummyOpRef",
						},
					}

					objectUnderTest := New(tempEventStore)

					eventChannel, _ := objectUnderTest.Subscribe(context.TODO(), subscriberEventFilter)

					/* act */
					objectUnderTest.Publish(publishedEvent)

					/* assert */
					Consistently(eventChannel).ShouldNot(Receive())
				})
			})
		})
	})
	Context("Subscribe", func() {
		Context("one publish has occurred", func() {
			Context("no filter", func() {
				It("should receive published event", func() {
					/* arrange */
					expectedEvent := model.Event{
						ContainerStarted: &model.ContainerStarted{
							RootCallID:  "dummyRootCallID",
							ContainerID: "dummyContainerID",
							OpRef:       "dummyOpRef",
						},
					}

					objectUnderTest := New(tempEventStore)
					objectUnderTest.Publish(expectedEvent)

					/* act */
					eventChannel, _ := objectUnderTest.Subscribe(context.TODO(), model.EventFilter{})

					/* assert */
					var actualEvent model.Event
					Eventually(eventChannel).Should(Receive(&actualEvent))
					// ignore timestamp
					actualEvent.Timestamp = expectedEvent.Timestamp
					Expect(actualEvent).To(Equal(expectedEvent))
				})
			})
			Context("filter allows previous publish", func() {
				It("should receive published event", func() {
					/* arrange */
					expectedEvent := model.Event{
						CallStarted: &model.CallStarted{
							RootCallID: "dummyRootCallID",
							CallID:     "dummyCallID",
							OpRef:      "dummyOpRef",
						},
					}

					providedFilter := model.EventFilter{
						Roots: []string{
							expectedEvent.CallStarted.RootCallID,
						},
					}

					objectUnderTest := New(tempEventStore)
					objectUnderTest.Publish(expectedEvent)

					/* act */
					eventChannel, _ := objectUnderTest.Subscribe(context.TODO(), providedFilter)

					/* assert */
					var actualEvent model.Event
					Eventually(eventChannel).Should(Receive(&actualEvent))
					// ignore timestamp
					actualEvent.Timestamp = expectedEvent.Timestamp
					Expect(actualEvent).To(Equal(expectedEvent))
				})
			})
		})
		Context("two publishes have occurred", func() {
			Context("no filter", func() {
				It("should receive published events", func() {
					/* arrange */
					expectedEvent1 := model.Event{
						ContainerStarted: &model.ContainerStarted{
							RootCallID:  "dummyRootCallID",
							ContainerID: "dummyContainerID",
							OpRef:       "dummyOpRef",
						},
						Timestamp: time.Now(),
					}

					expectedEvent2 := model.Event{
						CallStarted: &model.CallStarted{
							RootCallID: "dummyRootCallID",
							CallID:     "dummyCallID",
							OpRef:      "dummyOpRef",
						},
						Timestamp: time.Now().Add(time.Second),
					}

					tempDirPath := os.TempDir()
					tempEventStore := NewBadgerDBEventStore(tempDirPath)

					objectUnderTest := New(tempEventStore)
					objectUnderTest.Publish(expectedEvent1)
					objectUnderTest.Publish(expectedEvent2)

					/* act */
					eventChannel, _ := objectUnderTest.Subscribe(context.TODO(), model.EventFilter{})

					/* assert */
					var actualEvent1 model.Event
					Eventually(eventChannel).Should(Receive(&actualEvent1))
					// ignore timestamp
					actualEvent1.Timestamp = expectedEvent1.Timestamp
					Expect(actualEvent1).To(Equal(expectedEvent1))

					var actualEvent2 model.Event
					Eventually(eventChannel).Should(Receive(&actualEvent2))
					// ignore timestamp
					actualEvent2.Timestamp = expectedEvent2.Timestamp
					Expect(actualEvent2).To(Equal(expectedEvent2))
				})
			})
		})
	})
})
