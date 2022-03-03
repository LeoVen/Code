package entity

type Cellphone struct {
	Id         int    `json:"id"`
	ProviderId int    `json:"provider_id"`
	Total      int    `json:"total"`
	Number     string `json:"number"`
}

type Provider struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}
