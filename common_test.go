package common

import (
	"fmt"
	"testing"

	"github.com/ksclouds/common/crawler/wallpaper"
)

func TestWeather(t *testing.T) {
	rsp := wallpaper.GetList()
	fmt.Println(rsp)
}
