package healthcheck

// StatusResponse data transfer object
type StatusResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
