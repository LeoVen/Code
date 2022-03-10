package entity

// API entities

type ApiError struct {
	Error string `json:"error"`
}

type GetProviderByIdResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}

type GetProviderByNameResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}
