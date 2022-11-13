package models

type Configuration struct {
	Secure        bool   `json:"secure"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	AllowRun      bool   `json:"allowRun"`
	AllowShutdown bool   `json:"allowShutdown"`
	AllowReboot   bool   `json:"allowReboot"`
}
