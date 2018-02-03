package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Images", func() {
	var (
		client *fakes.ImagesClient
		logger *fakes.Logger

		images compute.Images
	)

	BeforeEach(func() {
		client = &fakes.ImagesClient{}
		logger = &fakes.Logger{}

		images = compute.NewImages(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListImagesCall.Returns.Output = &gcpcompute.ImageList{
				Items: []*gcpcompute.Image{{
					Name: "banana-image",
				}},
			}
			filter = "banana"
		})

		It("lists, filters, and prompts for images to delete", func() {
			list, err := images.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListImagesCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete image banana-image?"))

			Expect(list).To(HaveLen(1))
			Expect(list).To(HaveKeyWithValue("banana-image", ""))
		})

		Context("when the client fails to list images", func() {
			BeforeEach(func() {
				client.ListImagesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := images.List(filter)
				Expect(err).To(MatchError("Listing images: some error"))
			})
		})

		Context("when the image name does not contain the filter", func() {
			It("does not add it to the list", func() {
				list, err := images.List("grape")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ListImagesCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))

				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not add it to the list", func() {
				list, err := images.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})
	})

	Describe("Delete", func() {
		var list map[string]string

		BeforeEach(func() {
			list = map[string]string{"banana-image": ""}
		})

		It("deletes images", func() {
			images.Delete(list)

			Expect(client.DeleteImageCall.CallCount).To(Equal(1))
			Expect(client.DeleteImageCall.Receives.Image).To(Equal("banana-image"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting image banana-image\n"}))
		})

		Context("when the client fails to delete the image", func() {
			BeforeEach(func() {
				client.DeleteImageCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				images.Delete(list)

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting image banana-image: some error\n"}))
			})
		})
	})
})
