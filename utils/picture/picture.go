package picture

import (
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//Get 根据关键词获取相关配图地址
func Get(keyword string) string {
	var sources []string
	for word, src := range keywordMap {
		if strings.Contains(word, keyword) {
			sources = src
			break
		}
	} //匹配到关键词，随机返回一个图片地址
	if sources != nil {
		rnd := rand.New(rand.NewSource(time.Now().Unix()))
		index := rnd.Intn(len(sources))
		return sources[index]
	}
	return GetRandomPic("")
}

type pictureCrawler struct {
	request           *http.Request
	client            *http.Client
	mainPageFormat    string
	detailPagePattern string
	targetPattern     string
	search            string
}

func NewPictureCrawler(baseFormat, detailPattern, targetPattern, search string) *pictureCrawler {
	mainUrl := strings.Replace(baseFormat, "{search}", search, 1)
	mainUrl = strings.Replace(mainUrl, "{page}", "1", 1)
	req, _ := http.NewRequest("GET", mainUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	client := &http.Client{}
	return &pictureCrawler{
		req,
		client,
		baseFormat,
		detailPattern,
		targetPattern,
		search,
	}
}

func (c *pictureCrawler) Get(page ...int) ([]string, error) {
	var pictures []string
	if len(page) == 1 {
		mainUrl := strings.Replace(c.mainPageFormat, "{search}", c.search, 1)
		mainUrl = strings.Replace(mainUrl, "{page}", strconv.Itoa(page[0]), 1)
		c.request.URL, _ = url.Parse(mainUrl)
	}
	resp, err := c.client.Do(c.request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	mainPage, err := ioutil.ReadAll(resp.Body)
	root, err := htmlquery.Parse(strings.NewReader(string(mainPage)))
	if err != nil {
		return nil, err
	}
	detailPages := htmlquery.Find(root, c.detailPagePattern)
	for _, detail := range detailPages {
		detailUrl := detail.Attr[0].Val
		c.request.URL, _ = url.Parse(detailUrl)
		resp, err = c.client.Do(c.request)
		detailPage, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		root, err = htmlquery.Parse(strings.NewReader(string(detailPage)))
		if err != nil {
			return nil, err
		}
		pic := htmlquery.FindOne(root, c.targetPattern)
		if pic == nil {
			return nil, err
		}
		pictures = append(pictures, pic.Attr[1].Val)
		_ = resp.Body.Close()
		rnd := rand.New(rand.NewSource(time.Now().Unix()))
		sleep := rnd.Intn(3) + 1
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	return pictures, nil
}

//GetRandomPic 获取随机配图
func GetRandomPic(keyword string) string {
	return ""
}

func UpdatePictureSource(keyword string, page int) ([]string, error) {
	baseFormat := pictureConfig.GetString("pictureSource.source1.urlPattern")
	detailPattern := pictureConfig.GetString("pictureSource.source1.detailPattern")
	targetPattern := pictureConfig.GetString("pictureSource.source1.targetPattern")
	crawler := NewPictureCrawler(baseFormat, detailPattern, targetPattern, keyword)
	pics, err := crawler.Get(page)
	if err != nil {
		return nil, err
	}
	return pics, nil
}
