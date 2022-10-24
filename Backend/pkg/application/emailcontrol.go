package application

import (
	"net/http"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
	"github.com/Jmanuel2099/EnronCorpEmail/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	dataBasePath     = "../enron_mail_20110402/maildir"
	dataBasePathTest = "../dbTest/maildir"
	indexName        = "enron_email"
)

type EmailControl struct {
	zincSearchDb ZincSearchInterface
}

func NewEmailControl(z ZincSearchInterface) *EmailControl {
	return &EmailControl{
		zincSearchDb: z,
	}
}

// getEmailUsers is the response for the GetEmailUsers method
type getEmailUsers struct {
	Users []string
}

// GetEmailUsers searches the names of ENRON users
func (e EmailControl) GetEmailUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetEnronUsers(dataBasePath)
	if err != nil {
		return
	}
	response := &getEmailUsers{
		Users: users,
	}
	render.JSON(w, r, response)
}

// getEmailsByUser is the response for the GetEmailsByUser method
type getEmailsByUser struct {
	Emails []domain.Email
}

// GetEmailsByUser searches the contents of an enron user's e-mails
func (e EmailControl) GetEmailsByUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user-name")
	emails, err := service.GetEmailsByUser(dataBasePath, userId)
	if err != nil {
		return
	}
	response := &getEmailsByUser{
		Emails: emails,
	}
	render.JSON(w, r, response)
}
