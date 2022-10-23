package zincsearch

type bulkDocumentResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}
