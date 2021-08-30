package weather

import (
	"fmt"
	"testing"
)

func TestWeather(t *testing.T) {
	rsp := getWeather("南京")
	fmt.Println(rsp)
}
