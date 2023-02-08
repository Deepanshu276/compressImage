package cmd

import (
	"fmt"
	"image"
	"image/jpeg"

	"image/png"

	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var quality int
var input string
var output string

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress image with specified quality",
	Long:  `Compress image with specified quality`,
	Run: func(cmd *cobra.Command, args []string) {
		err := CompressImage(input, output, quality)
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println("Image compression successful")
			fmt.Println("Compressed image stored at", output)

		}

	},
}

func CompressImage(input string, output string, quality int) error {
	img, format, err := openImage(input)
	if err != nil {
		return err
	}
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()
	err = encodeImage(img, format, f, quality)
	if err != nil {
		return err
	}
	return nil
}

func openImage(input string) (image.Image, string, error) {
	if input[:4] == "http" {
		response, err := http.Get(input)
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
	f, err := os.Open(input)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

func encodeImage(img image.Image, format string, output io.Writer, quality int) error {
	var opt jpeg.Options
	opt.Quality = quality
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

func init() {
	compressCmd.Flags().IntVarP(&quality, "quality", "q", 0, "image compression ranging from 0-100")
	compressCmd.Flags().StringVarP(&input, "input", "i", "", "input file or url")
	compressCmd.Flags().StringVarP(&output, "output", "o", "", "compressed image path where you want to store the output image")
	compressCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(compressCmd)
}
