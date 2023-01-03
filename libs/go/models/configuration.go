package models

type Configuration struct {
	Secure bool           `json:"secure"`
	Token  string         `json:"token"`
	Host   string         `json:"host"`
	Path   string         `json:"path"`
	Port   string         `json:"port"`
	Allow  AllowFunctions `json:"allow"`
}

type AllowFunctions struct {
	Run            bool `json:"run"`
	Shutdown       bool `json:"shutdown"`
	Reboot         bool `json:"reboot"`
	ProcessList    bool `json:"processList`
	ServiceList    bool `json:"serviceList`
	ServiceLogs    bool `json:"serviceLogs`
	ServiceStop    bool `json:"serviceStop`
	ServiceStart   bool `json:"serviceStart`
	ServiceStatus  bool `json:"serviceStatus`
	ServiceRestart bool `json:"serviceRestart`
	Kill           bool `json:"kill`
}