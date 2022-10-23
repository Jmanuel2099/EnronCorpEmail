package zincsearch

import "github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"

type bulkDocumentsRequest struct {
	Index   string         `json:"index"`
	Records []domain.Email `json:"records"`
}
