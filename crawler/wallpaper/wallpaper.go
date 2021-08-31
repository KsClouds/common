package wallpaper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func GetList() (rsp []string) {
	url := "https://pic.netbian.com/"

	c := colly.NewCollector()

	c.OnHTML("img", func(h *colly.HTMLElement) {
		src := h.Attr("src")
		fmt.Println(src)
		if strings.EqualFold(src[:8], "/uploads") {
			rsp = append(rsp, src)
		}
	})
	c.Visit(url)

	c.Wait()
	return
}
