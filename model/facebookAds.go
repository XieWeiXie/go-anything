package model

type FaceBookAds struct {
	BrandName string `json:"brand_name"` // 品牌名称
	//BrandId        string `json:"brand_id"`         // 品牌ID
	//Like           int    `json:"like"`             // 主页获赞
	//BrandCrateTime string `json:"brand_crate_time"` // 主页创建时间
	Status       string `json:"status"`        // 投放状态
	Version      string `json:"version"`       // 版本
	Number       string `json:"number"`        // 编号
	RunningTime  string `json:"running_time"`  // 广告投放时间
	BrandContent string `json:"brand_content"` // 文案内容
	AdVideoUrl   string `json:"ad_video_url"`  // 广告视频地址
	AdImageUrl   string `json:"ad_image_url"`  // 广告封面地址
	AdContents   string `json:"ad_contents"`   // 广告宣语
	ShoppingUrl  string `json:"shopping_url"`  // 购买时跳转的链接
	ShoppingHost string `json:"shopping_host"` // 购买链接域名
}
