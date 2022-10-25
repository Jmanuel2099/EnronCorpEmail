package application

import (
	"encoding/json"
	"net/http"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/datasource/zincsearchdb"
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
	Users []string `json:"users"`
}

// GetEmailUsers searches the names of ENRON users
func (e EmailControl) GetEmailUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	users, err := service.GetEnronUsers(dataBasePath)
	if err != nil {
		return
	}
	response := &getEmailUsers{
		Users: users,
	}
	render.JSON(w, r, response)
}

// bulkDocuments is the response for the BulkDocuments method
type BulkDocuments struct {
	Emails zincsearchdb.BulkDocumentResponse `json:"emails"`
}

// BulkDocuments indexes the emails in ZincSerach
func (e EmailControl) BulkDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	userId := chi.URLParam(r, "user-name")
	emails, err := getEmailsByUser(userId)
	if err != nil {
		return
	}

	bulkResponse, err := e.zincSearchDb.BulkDocument(indexName, emails)
	if err != nil {
		return
	}

	response := &BulkDocuments{
		Emails: *bulkResponse,
	}
	render.JSON(w, r, response)
}

// searchDocumentsByfilter is the response for the SearchDocumentsByfilter method
type searchDocumentsByfilter struct {
	Emails []domain.Email `json:"emails"`
}

// SearchDocumentsByfilter searches for the content of emails that have the filter in their content.
func (e EmailControl) SearchDocumentsByfilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	temp := chi.URLParam(r, "temp")

	documentsResponse, err := e.zincSearchDb.SearchDocuments(indexName, temp)
	if err != nil {
		return
	}
	emails, err := mapEmailsResponse(*documentsResponse)
	if err != nil {
		return
	}
	response := &searchDocumentsByfilter{
		Emails: emails,
	}
	render.JSON(w, r, response)
}

// GetEmailsByUser searches the database for the contents of an enron user's emails.
func getEmailsByUser(user string) ([]domain.Email, error) {
	emails, err := service.GetEmailsByUser(dataBasePath, user)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func mapEmailsResponse(response zincsearchdb.SearchDocumentsResponse) ([]domain.Email, error) {
	var emails []domain.Email
	for _, hit := range response.Hits.Hits {
		var email domain.Email

		bytes, err := json.Marshal(hit.Source)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, &email)
		if err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	print("before -> ", emails)
	return emails, nil
}
