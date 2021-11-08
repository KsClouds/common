package common

import (
	"fmt"
	"github.com/ksclouds/common/crawler/dujitang"
	"testing"

	"github.com/ksclouds/common/crawler/wallpaper"
)

func TestWeather(t *testing.T) {
	rsp := wallpaper.GetList()
	fmt.Println(rsp)
}

func TestDuJiTang(t *testing.T) {
	rsp := dujitang.Get()
	fmt.Println(rsp)
}

func TestGetWallpaper(t *testing.T) {
	rsp := wallpaper.GetList()
	fmt.Println(rsp)
}
