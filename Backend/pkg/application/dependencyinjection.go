package application

import (
	"github.com/Jmanuel2099/EnronCorpEmail/pkg/datasource/zincsearchdb"
	"github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"
)

// ZincSearchInterface is the interface that must implement any database to be used in the project.
type ZincSearchInterface interface {
	BulkDocument(indexName string, emalRecords []domain.Email) (*zincsearchdb.BulkDocumentResponse, error)
	SearchDocuments(indexName string, bodyRequest zincsearchdb.SearchDocumentsRequest) (*zincsearchdb.SearchDocumentsResponse, error)
}
