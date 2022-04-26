package models

type Registration struct {
	IntID   int    `json:"IntID"`
	AppID   string `json:"AppID"`
	AppName string `json:"AppName"`
	LOBID   string `json:"LOBID"`
	AppEnv  string `json:"AppEnv"`
	AppLoc  string `json:"AppLoc"`
}
