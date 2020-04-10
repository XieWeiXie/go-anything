package MakeMoney

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
	"strings"
	"testing"
)

func TestRealMakeMoneyAction_Do(t *testing.T) {

	mm := NewRealMakeMoneyAction(topicURL)
	content := mm.Do()

	for _, i := range content {
		ts := template.New("doc")
		box := packr.NewBox("../../doc/money/template")
		templateText, e := box.FindString("template.md")
		tem, e := ts.Parse(templateText)
		if e != nil {
			log.Println(e)
			return
		}
		var byt bytes.Buffer
		e = tem.Execute(&byt, i)
		if e != nil {
			log.Println(e)
			return
		}
		toSave(i.Title, byt.Bytes())

	}

}

func toSave(title string, content []byte) error {
	replacer := strings.NewReplacer("，", "", "\"", "", " ", "",
		"/", "")
	title = replacer.Replace(title)
	f, e := os.Create(fmt.Sprintf("../../doc/money/%s.md", title))
	if e != nil {
		log.Println(e)
		return e
	}
	f.Write(content)
	defer f.Close()
	return nil
}

func TestUrl(t *testing.T) {
	text := "<e type=\"hashtag\" hid=\"552815125124\" title=\"%23%E6%A1%88%E4%BE%8B%E5%88%86%E4%BA%AB%20%23\" />思路分享  一个卖得了几百万豪车，也卖得了几十块生蚝的90后的生意经。\n今天，我聊聊一个认识的90后的生意经，他既卖几百万的豪车，也卖几十块的生蚝，而且把两个跨界的生意做得都红红火火，豪车收入没问，120平米左右的生蚝餐饮店最高峰时月营业额400多万，平常时也能做到200-300万月营业额，做过餐饮的朋友可以算算这坪效有多惊人。这个90后，是老纪，差不多是一年前这个时候，通过我老板介绍，我第一次见到他，这一年中多次听他分享他做生意的经验，每次都很受启发，附上我的一些理解，一并分享给生财有术的朋友们。\n\n听他说，他十四岁时，随全家人一起从江西搬去上海，跟着家人一起开始做豆芽菜生意，后来在2009年，从上海搬到杭州，在一个4S店做销售。这个过程中积攒了一些买车的顾客资源和4S店资源，后来他出来自己开了一个公司卖汽车，从卖大众价格的车开始，现在转型为销售豪车，他说他是从最早的8个汽车客户开始做起的，卖给他们的都是几万块的车。卖车、卖豪车这件事本身稀松平常，没太多值得说的，老纪这个生意值得说的是，他是整合资源做平台的，公司没有任何现货车，没有仓库，客户要买车，他报价，客户先下定金，他从全国各地调车，给客户送货上门。所以，他可以根据客户的实际情况，进行品牌和车型的推荐，而不像4S店里，只能拼命推销自家品牌的车型。最绝的还不在于没有库存车，而是他和很多客户从来没见过面，通过手机聊天就把交易给谈成了，从服务最早的8个客户开始，一个推荐一个，建立了强大的信任纽带，几十几百万的一笔交易，不见面，也不看货，就通过QQ、微信（在没有微信时，他主要依赖QQ维护客户）聊天，就把生意给做成了。（刀姐这篇私域流量的文章中谈到第三种私人伙伴的玩法，正是老纪实践了多年的方法：<e type=\"web\" href=\"https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FVUkiwU2a2HcqCMeQholqhA\" title=\"2019%E7%88%86%E7%81%AB%E9%BB%91%E8%AF%9D%E3%80%8C%E7%A7%81%E5%9F%9F%E6%B5%81%E9%87%8F%E3%80%8D%E7%9A%84%E6%9C%AC%E8%B4%A8%E5%92%8C%E7%8E%A9%E6%B3%95\" cache=\"\" />）。\n\n老纪因为掌握着全国汽车价格，"
	fmt.Println(formatText(text))

	file := "https://files.zsxq.com/lt9IJQoZ_WJw-daLgCNKQP3SWsob?attname=%E5%85%9A%E5%BB%BA%E4%BF%A1%E6%81%AF%E5%B9%B3%E5%8F%B0.pptx&e=1586620524&token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:YVRuaYDFVgMxgWPUzvzzvdiDn7E="
	fmt.Println(formatText(file))
}
