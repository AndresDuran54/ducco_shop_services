package config

var AppInfo = appInfo{}

type appInfo struct {
	Version string
}

func init() {
	AppInfo = appInfo{
		Version: "v1",
	}
}
