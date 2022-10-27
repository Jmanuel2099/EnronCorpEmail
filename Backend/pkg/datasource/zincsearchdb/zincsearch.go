package zincsearchdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

const (
	zincSearchHost = "http://localhost:4080"
)

type ZincSearchClient struct {
	Client *http.Client
}

func NewZincSearchClient(c *http.Client) *ZincSearchClient {
	return &ZincSearchClient{
		Client: c,
	}
}

// BulkDocument upload the content of multiple emails into ZincSearch
func (c *ZincSearchClient) BulkDocument(indexName string, emalRecords []domain.Email) (*BulkDocumentResponse, error) {
	bodyResponse := &BulkDocumentResponse{}
	url := fmt.Sprintf("%s/api/_bulkv2", zincSearchHost)
	bodyRequest := BulkDocumentsRequest{
		Index:   indexName,
		Records: emalRecords,
	}

	req, err := http.NewRequest(http.MethodPost, url, adapterBodyRequest(bodyRequest))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	response, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(bodyResponse)

	return bodyResponse, nil
}

// SearchDocuments finds the documents that I have a match with a filter in ZincSearch
func (c *ZincSearchClient) SearchDocuments(indexName, term string) (*SearchDocumentsResponse, error) {
	bodyResponse := &SearchDocumentsResponse{}

	url := fmt.Sprintf("%s/es/%s/_search", zincSearchHost, indexName)
	bodyRequest := SearchDocumentsRequest{
		Query_string: SearchDocumentsRequestQuery{
			Query: term,
		},
		Sort: []string{"-@timestamp"},
		From: 0,
		Size: 12,
	}

	req, err := http.NewRequest(http.MethodPost, url, adapterBodyRequest(bodyRequest))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	response, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(bodyResponse)
	return bodyResponse, nil
}

func adapterBodyRequest(bodyRequest interface{}) io.Reader {
	if bodyRequest == nil {
		return nil
	}
	body, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(body)
}
