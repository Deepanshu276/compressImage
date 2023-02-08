package cmd

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/disintegration/imaging"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompressCmd", func() {
	It("compresses an image from the web", func() {
		url := "https://fileinfo.com/img/ss/xl/jpeg_43.png"
		quality := 75

		// Get the image from the URL
		response, err := http.Get(url)
		Expect(err).To(BeNil(), "Error fetching image")
		defer response.Body.Close()

		// Decode the image
		img, err := imaging.Decode(response.Body)
		Expect(err).To(BeNil(), "Error decoding image")

		// Resize the image
		img = imaging.Resize(img, 0, 800, imaging.Lanczos)

		// Save the compressed image with specified quality
		err = imaging.Save(img, "compressed.jpeg", imaging.JPEGQuality(quality))
		Expect(err).To(BeNil(), "Error saving compressed image")

		fmt.Println("Compressed image saved as compressed.jpeg")
	})
})

func TestCompressCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CompressCmd Suite")
}
