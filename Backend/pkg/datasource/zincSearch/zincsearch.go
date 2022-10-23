package zincsearch

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

const (
	zincSearchHost = "http://localhost:4080/"
	// headerContentType     = "Content-Type"
	// headerApplicationJSON = "application/json"
)

func BulkDocument(indexName string, emalRecords []domain.Email) (*bulkDocumentResponse, error) {
	bodyResponse := &bulkDocumentResponse{}
	url := zincSearchHost + "api/_bulkv2"
	bodyRequest := bulkDocumentsRequest{
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

	client := http.Client{}
	response, err := client.Do(req)
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
