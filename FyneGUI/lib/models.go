package lib

import "log"

type AppInfo struct {
	Id     int
	Name   string
	Source string
	Type   string
}

func (appInfo AppInfo) Run() bool {
	log.Println(appInfo.Id, appInfo.Name)
	return true
}

func (appInfo AppInfo) GetDetail() string {
	return appInfo.Name
}

type App interface {
	Run() bool
	GetDetail() string
}
