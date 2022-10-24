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
	zincSearchHost = "http://localhost:4080/"
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
	url := fmt.Sprintf("%sapi/_bulkv2", zincSearchHost)
	bodyRequest := BulkDocumentsRequest{
		Index:   indexName,
		Records: emalRecords,
	}
	response, err := makeRequest(http.MethodPost, url, bodyRequest, *c.Client)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(response.Body).Decode(bodyResponse)

	return bodyResponse, nil
}

// SearchDocuments finds the documents that I have a match with a filter in ZincSearch
func (c *ZincSearchClient) SearchDocuments(indexName string, bodyRequest SearchDocumentsRequest) (*SearchDocumentsResponse, error) {
	bodyResponse := &SearchDocumentsResponse{}
	url := fmt.Sprintf("%s/api/%s/_search", zincSearchHost, indexName)

	response, err := makeRequest(http.MethodPost, url, bodyRequest, *c.Client)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(response.Body).Decode(bodyResponse)

	return bodyResponse, nil
}

func makeRequest(verbHttp, url string, body interface{}, client http.Client) (http.Response, error) {
	req, err := http.NewRequest(verbHttp, url, adapterBodyRequest(body))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	return *response, nil
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
