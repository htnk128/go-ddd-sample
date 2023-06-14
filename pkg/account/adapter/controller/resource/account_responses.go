package resource

type AccountResponses struct {
	Count   int                `json:"count"`
	HasMore bool               `json:"has_more"`
	Data    []*AccountResponse `json:"data"`
}
