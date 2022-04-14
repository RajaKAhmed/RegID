package models

type Registration struct {
	AppID   uint   `json:"AppID"`
	AppName string `json:"AppName"`
	LOBID   string `json:"LOBID"`
	AppEnv  string `json:"AppEnv"`
	AppLoc  string `json:"AppLoc"`
}
