package zincsearchdb

import "github.com/Jmanuel2099/EnronCorpEmail/pkg/domain"

type BulkDocumentsRequest struct {
	Index   string         `json:"index"`
	Records []domain.Email `json:"records"`
}

type SearchDocumentsRequest struct {
	SearchType string                      `json:"search_type"`
	SortFields []string                    `json:"sort_fields"`
	From       int                         `json:"from"`
	MaxResults int                         `json:"max_results"`
	Query      SearchDocumentsRequestQuery `json:"query"`
	Source     map[string]interface{}      `json:"_source"`
}

type SearchDocumentsRequestQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
