package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

var dbContent []domain.Email

// GetDataBaseContent gets the contents of the files in a folder.
func GetDataBaseContent(path string) []domain.Email {
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Print("Error occurred when opening a folder")
		return nil
	}

	for _, directory := range directories {
		newPath := fmt.Sprintf("%s/%s", path, directory.Name())
		if directory.IsDir() {
			GetDataBaseContent(newPath)
		} else {
			email := readEmailFileContent(newPath)
			dbContent = append(dbContent, *email)
		}
	}
	return dbContent
}

// readEmailFileContent gets the contents of an email file.
func readEmailFileContent(pathFile string) *domain.Email {
	file, err := os.ReadFile(pathFile)
	if err != nil {
		fmt.Print("Error occurred when opening a file")
		return nil
	}
	return mapEmailContent(string(file))

}

// mapEmailContent maps a string to an Email struct.
func mapEmailContent(content string) *domain.Email {
	email := &domain.Email{}
	contentFile := strings.SplitN(content, "\r\n\r\n", 2)
	emailDetails := strings.Split(contentFile[0], "\r\n")

	for _, emailDetail := range emailDetails {
		detail := strings.SplitN(emailDetail, ":", 2)

		switch detail[0] {
		case "Message-ID":
			email.MessageId = detail[1]
		case "From":
			email.From = detail[1]
		case "To":
			email.To = detail[1]
		case "Date":
			email.Date = detail[1]
		case "Subject":
			email.Subject = detail[1]
		default:
			continue
		}
	}
	email.Content = contentFile[1]
	return email
}
