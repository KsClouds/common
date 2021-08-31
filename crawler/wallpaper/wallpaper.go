package wallpaper

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetList() (rsp []string) {
	url := "https://pic.netbian.com/index_" + strconv.Itoa(rand.Intn(1275)+1) + ".html"

	c := colly.NewCollector()

	c.OnHTML("img", func(h *colly.HTMLElement) {
		src := h.Attr("src")
		if len(src) > 8 && strings.EqualFold(src[:8], "/uploads") {
			rsp = append(rsp, src)
		}
	})
	c.Visit(url)

	c.Wait()
	return
}
