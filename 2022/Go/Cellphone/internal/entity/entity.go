package entity

type Cellphone struct {
	Id         int    `json:"id,omitempty"`
	ProviderId int    `json:"provider_id,omitempty"`
	Total      int    `json:"total,omitempty"`
	Number     string `json:"number,omitempty"`
}

type Provider struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
