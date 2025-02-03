package model

type CreateRequest struct {
	Type   string `json:"type"`
	Total  int    `json:"total"`
	Action string `json:"action"`
}

type CreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrCreateResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type TestNumberCounter struct {
	Type        string `gorm:"primaryKey"`
	StartNumber int
	LastNumber  *int
}

type TestNumberTransaction struct {
	Number int `gorm:"primaryKey"`
	Action string
}
