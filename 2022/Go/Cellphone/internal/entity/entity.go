package entity

type ApiError struct {
	Error string `json:"error"`
}

type Cellphone struct {
	Id         int    `json:"id"`
	ProviderId int    `json:"provider_id"`
	Number     string `json:"number"`
}

type Provider struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}
