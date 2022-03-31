package entity

import pb "cellphone/protos/go"

type ApiError = pb.ApiError

type Cellphone = pb.Cellphone

type Provider = pb.Provider

// type ApiError struct {
// 	Error string `json:"error"`
// }

// type Cellphone struct {
// 	Id         int    `json:"id" gorm:"column:ID"`
// 	ProviderId int    `json:"provider_id" gorm:"column:PROVIDER_ID"`
// 	Number     string `json:"number" gorm:"column:NUMBER"`
// }

// type Provider struct {
// 	Id    int    `json:"id" gorm:"column:ID"`
// 	Name  string `json:"name" gorm:"column:NAME"`
// 	Total int    `json:"total" gorm:"column:TOTAL"`
// }
