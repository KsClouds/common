package common

import (
	"github.com/ksclouds/common/crawler/dujitang"
	"github.com/ksclouds/common/crawler/wallpaper"
	"github.com/ksclouds/common/crawler/weather"
)

func GetDujitang() (rsp string) {
	return dujitang.Get()
}

func GetWeather(area string) (rsp string) {
	return weather.Get(area)
}

func GetWallpaper() (rsp []string) {
	return wallpaper.GetList()
}
