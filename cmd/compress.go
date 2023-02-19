package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var resize int
var input string
var output string

// compressCmd represents the sub command called with the base command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress image with specified quality",
	Long: `This is a command line tool for resizing the images.
			To create a compress command it takes in an input file or URL and compresses the image
			with a specified quality, then outputs the compressed image to a specified path.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := CompressImage(filepath.Join(filepath.Dir(input), filepath.Base(input)), output, resize)
		//err := CompressImage(input, output, resize)

		if err != nil {
			fmt.Println(err)
		} else if output == "" {
			fmt.Println("Compressed image stored at Pictures Directory")

		} else {
			fmt.Println("Compressed image stored at", output)
		}
		os.Exit(0)

	},
}

// Logic to compress the code
func CompressImage(input string, output string, resize int) error {

	img, format, err := OpenImage(input)
	if err != nil {
		return err
	}

	//If output path is not specified
	if output == "" {
		//Join the Pictures folder with the file name
		output = filepath.Join(os.Getenv("HOME"), "Pictures", filepath.Base(input))
	}

	//Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	//Close closes the File, rendering it unusable for I/O.
	defer f.Close()
	err = encodeImage(img, format, f, resize)
	if err != nil {
		return err
	}
	return nil

}

func OpenImage(input string) (image.Image, string, error) {
	if input[:4] == "http" {
		//if request is not served with the provided time we we get the timeout error
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		response, err := client.Get(input)
		if err != nil {
			return nil, "", err
		}
		defer response.Body.Close()
		img, format, err := image.Decode(response.Body)
		if err != nil {
			return nil, "", err
		}
		return img, format, nil
	}
	//Open the named file for reading
	f, err := os.Open(input)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	//Decode decodes an image that has been encoded in a registered format
	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil

}

// encoding the image will only support jpeg and png
func encodeImage(img image.Image, format string, output io.Writer, resize int) error {
	var opt jpeg.Options
	opt.Quality = resize
	switch format {
	case "jpeg":
		err := jpeg.Encode(output, img, &opt)
		if err != nil {
			return err
		}
	case "png":
		err := png.Encode(output, img)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}
	return nil
}

// Function to check if the input is a URL or not

// Here you will define your flags and configuration settings.
// Cobra supports persistent flags, which, if defined here,
// will be global for your application.

func init() {
	compressCmd.Flags().IntVarP(&resize, "resize", "r", 0, "image compression number between [0,100]")
	compressCmd.Flags().StringVarP(&input, "input", "i", "", "input image or url")
	compressCmd.Flags().StringVarP(&output, "output", "o", "", "compressed image path where you want to store the output image")
	compressCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(compressCmd)
}
