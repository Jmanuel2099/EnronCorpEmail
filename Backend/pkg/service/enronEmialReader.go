package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

var dbContent []domain.Email

// getDataBaseContent gets the content of the files in a folder
func getUserFolderContent(path string) ([]domain.Email, error) {
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error occurred when opening a folder %s", path)
		return nil, err
	}

	for _, directory := range directories {
		newPath := fmt.Sprintf("%s/%s", path, directory.Name())
		if directory.IsDir() {
			getUserFolderContent(newPath)
		} else {
			email := readEmailFileContent(newPath)
			dbContent = append(dbContent, *email)
		}
	}
	return dbContent, nil
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

// GetEnronUsers gets the names of Enron users' folders
func GetEnronUsers(path string) ([]string, error) {
	var users []string
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error occurred when opening a folder: %s", path)
		return nil, err
	}

	for _, user := range directories {
		if user.IsDir() {
			users = append(users, user.Name())
		}
	}
	return users, nil
}

// GetEmailsByUser gets the emails that a user has
func GetEmailsByUser(path, userName string) ([]domain.Email, error) {
	emails, err := getUserFolderContent(fmt.Sprintf("%s/%s", path, userName))
	if err != nil {
		return nil, err
	}
	return emails, nil
}
