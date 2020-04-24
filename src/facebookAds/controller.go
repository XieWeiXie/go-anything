package facebookAds

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/wuxiaoxiaoshen/go-anything/model"
	chromedp_helper "github.com/wuxiaoxiaoshen/go-anything/pkg/chromedp"
	"log"
	"strings"
)

type ControllerFaceBookAds struct {
	Url string `json:"url"`
}

func NewControllerFaceBookAds(brandId string) *ControllerFaceBookAds {
	return &ControllerFaceBookAds{
		Url: fmt.Sprintf(facebookUrl, brandId),
	}
}

func (F ControllerFaceBookAds) Do() []model.FaceBookAds {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var res string
	res = chromedp_helper.GetPageSource(ctx, F.Url)
	//log.Println(strings.TrimSpace(res))
	//f, _ := os.Open("ad.html")
	//b, _ := ioutil.ReadAll(f)
	//doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(b))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(res))
	var result []model.FaceBookAds

	doc.Find("div._7jj- ._7owt").Each(func(i int, selection *goquery.Selection) {
		var ad model.FaceBookAds
		version := selection.Find("a")
		ad.Version = version.First().Text()

		partOne := selection.Find("div div._7jvy div._8k-- div").Eq(0)
		//log.Println(partOne.Html())

		ad.Status = partOne.Find("div._7jv-").Text()
		ad.RunningTime = partOne.Find("div._7jwu").Text()
		ad.Number = partOne.Find("div._8jox").Text()

		partTwo := selection.Find("div div._7jwy div._7jyg._7jyh")
		//log.Println(partTwo.Html())
		ad.BrandName = partTwo.Find("div").Eq(0).Text()
		ad.BrandName = ad.BrandName[:strings.Index(ad.BrandName, "Sponsored")]
		ad.BrandContent = partTwo.Find("div._7jyr div").First().Text()

		links := partTwo.Find("a._231w._231z._4yee")
		//log.Println(links.Html())
		ad.ShoppingUrl, _ = links.Attr("href")

		footer := links.Find("div>div").First()
		ad.AdContents = footer.Find("div").First().Text()
		ad.ShoppingHost = footer.Find("div").Last().Text()

		if partTwo.Children().Size() == 4 {
			ad.AdVideoUrl, _ = partTwo.Find("div._8o0a._8o0b").Find("video").Attr("src")
		} else {
			ad.AdImageUrl, _ = links.Find("img").Attr("src")
		}

		log.Println(ad)
		result = append(result, ad)
	})
	return nil
}
