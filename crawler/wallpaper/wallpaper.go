package wallpaper

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// HandleError 错误处理
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// GetHTML 抽取根据url获取内容
func GetHTML(url string) (html string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	html = string(pageBytes)
	return html
}

var (
	// 定义正则匹配图片链接的规则
	reImg = `https://uploadfile.bizhizu.cn[^"]+?(\.jpg)`
)

// 获取当前页图片链接
func getUrls(url string) (urls []string) {
	pageStr := GetHTML(url)
	// 正则匹配网页所有的图片urls
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	// 由于re.FindAllStringSubmatch的特性，返回的不是纯粹的urls，所以重新提取赋值
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return urls
}

// url是传的整页链接
func getImgUrls(url string) (urls []string) {
	urls = getUrls(url)
	return urls
}

func GetList() (rsp []string) {
	randomNum := rand.Intn(2)

	if randomNum == 0 {
		url := "https://pic.netbian.com/index_" + strconv.Itoa(rand.Intn(1275)+1) + ".html"

		c := colly.NewCollector()

		c.OnHTML("img", func(h *colly.HTMLElement) {
			src := h.Attr("src")
			if len(src) > 8 && strings.EqualFold(src[:8], "/uploads") {
				rsp = append(rsp, "https://pic.netbian.com"+src)
			}
		})
		c.Visit(url)

		c.Wait()
		return
	} else {
		i := rand.Intn(18) + 1
		rsp := getImgUrls("https://www.bizhizu.cn/shouji/tag-%E7%BE%8E%E5%A5%B3/" + strconv.Itoa(i) + ".html")
		return rsp[3:]
	}

}
