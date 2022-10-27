package service

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

var dbContent []domain.Email

// GetDataBaseContent gets the content of all emails from the enron database
func GetDataBaseContent(path string) ([]domain.Email, error) {
	directories, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error occurred when opening a folder %s", path)
		return nil, err
	}

	// wg := sync.WaitGroup{}
	// mute := sync.Mutex{}
	for _, userDirectory := range directories {
		if userDirectory.IsDir() {
			// wg.Add(1)
			userPath := fmt.Sprintf("%s/%s", path, userDirectory.Name())
			getUserFolderContent(userPath)
		}
	}
	// wg.Wait()
	return dbContent, err
}

// getUserFolderContent get the emails of every enron user that is in the database
func getUserFolderContent(path string) {
	var emails []domain.Email
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error in folder oro file: %s", path)
		}
		if !info.IsDir() {
			email := readEmailFileContent(path)
			if email == nil {
				return nil
			}
			emails = append(emails, *email)
		}
		return nil
	})
	// mute.Lock()
	dbContent = append(dbContent, emails...)
	// mute.Unlock()
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
