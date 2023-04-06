package helpers

import (
	"context"
	"fmt"
	"os"

	"io/ioutil"

	//"time"

	"mime/multipart"

	"path/filepath"

	"github.com/bivek/fmt_backend/constants"
)

func FileUpload(ctx context.Context, fileHeader *multipart.FileHeader, folderName constants.Folder) (string, error) {
	fileExtension := filepath.Ext(fileHeader.Filename)
	fmt.Printf("file extension %v", fileExtension)
	file, err := fileHeader.Open()

	if err != nil {
		fmt.Printf("Failed to open file")
		return "", err
	}

	defer file.Close()

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	formatedFile := fmt.Sprintf("clientimg*%s", fileExtension)
	tempFile, err := ioutil.TempFile("/files", formatedFile)
	filename := tempFile.Name()
	fmt.Printf("temp file name %v", filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Uploaded File")
	fmt.Println("Uploaded File Name:", filename)
	return filename, nil
}

func DeleteFileUpload(filepath string) {
	fmt.Println("filepath", filepath)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Error on Deleting file", err)
	}
	fmt.Printf("Image deleted from directory")
}
