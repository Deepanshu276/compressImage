# Image Compression Command Line Tool

This repository contains a command line tool for compressing images. It uses Cobra, a library for creating powerful modern CLI applications, to create a compress command that takes in an input file or URL and compresses the image with a specified quality, then outputs the compressed image to a specified path.

# Content

* Installation
* Usage
* Commands
  * Compress
* Test

# Prerequisites
Go installed in the local computer with version >= 1.18.

In case for upgradation refer this Link [Golang](https://www.golinuxcloud.com/upgrade-go-version/)

# Installation
First clone the repo in your $Path. A common place would be within your $GOPATH

Build and copy ```imgCompress ``` to your $GOPATH/bin:

```
$ go build .
```

#Usage

The tool uses the following flags:

* -q or --quality: Specify the quality of the compressed image.
* -i or --input: Specify the input file or URL.
* -o or --output: Specify the path to store the output image.

Example usage:

```
./imgCompress compress -q 50 -i [url or local image directory] -o [location where output image will be stored]

./imgCompress compress -q 50 -i image.jpg -o compressed.jpg

```

This will compress image.jpg with a quality of 50 and output the result to compressed.jpg.

# Test
To test the cli using different test cases 
1. cd into test folder using ```cd test```.
2. To run all test file all at once use 
```
go test
```
# Supported Image Formats

The tool supports both JPEG and PNG image formats. If the input image format is not supported, the tool will return an error.

# Contributing
If you would like to contribute to the project, please fork the repository and make your changes. Once you have finished your changes, please submit a pull request for review.






