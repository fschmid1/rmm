package models

type Configuration struct {
	Secure              bool   `json:"secure"`
	Host                string `json:"host"`
	Port                string `json:"port"`
	AllowRun            bool   `json:"allowRun"`
	AllowShutdown       bool   `json:"allowShutdown"`
	AllowReboot         bool   `json:"allowReboot"`
	AllowProcessList    bool   `json:"allowProcessList`
	AllowServiceList    bool   `json:"allowServiceList`
	AllowServiceLogs    bool   `json:"allowServiceLogs`
	AllowServiceStop    bool   `json:"allowServiceStop`
	AllowServiceStart   bool   `json:"allowServiceStart`
	AllowServiceStatus  bool   `json:"allowServiceStatus`
	AllowServiceRestart bool   `json:"allowServiceRestart`
	AllowKill           bool   `json:"allowKill`
}
