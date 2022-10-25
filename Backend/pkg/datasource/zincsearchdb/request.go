package zincsearchdb

import "github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"

type BulkDocumentsRequest struct {
	Index   string         `json:"index"`
	Records []domain.Email `json:"records"`
}

type SearchDocumentsRequest struct {
	Sort         []string                    `json:"sort"`
	From         int                         `json:"from"`
	Size         int                         `json:"size"`
	Query_string SearchDocumentsRequestQuery `json:"query_string"`
}

type SearchDocumentsRequestQuery struct {
	Query string `json:"query"`
}
