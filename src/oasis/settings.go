package oasis

/*
微博清爽社交平台：绿洲 接口
*/

var (
	WOW      = "https://oasis.weibo.cn/v1/wow/list" // 用户信息
	PROFILE  = "https://oasis.weibo.cn/v1/user/profile"
	TIMELINE = "https://oasis.weibo.cn/v1/timeline/user"
	DISCOVER = "https://oasis.weibo.cn/v1/timeline/discovery"
	Upload   = "https://picupload.weibo.com/interface/upload.php?appid=3157701306&cs=3588630082&ent=xiancheng&file_source=53&gsid=_2A9Y6hVryDIVOw5wYhtwUazlgUavKDdCBN9YkoHznUhyKSZN6hPZZsD-ykK3SLhoo4M4sybT2OFwCyfIMAIPJgxAu&ori=1&uid=1948244870"
	Create   = "https://oasis.weibo.cn/v1/status/create"
)

var (
	Star    = 16 // 明星
	Face    = 11 // 颜值
	Photo   = 18 // 摄影
	Fashion = 6  // 时尚
)

/*
绿洲 API
- status:
	/v1/status/destroy
	/v1/status/create
	/v1/status/detail
- at
	/v1/status/at/statuses
- like
	/v1/status/like/statuses
- timeline
	/v1/timeline/user
	/v1/timeline/following
	/v1/timeline/discovery
- report
	/v1/report/stats
- h5
	/v1/h5/share
- user
	/v1/user/profile
	/v1/user/star
	/v1/user/following_recommend
	/v1/user/logout
	/v1/user/captcha
	/v1/user/login
- wow
	/v1/wow/list
- sso
	/v1/sso/mfp
- message
	/v1/message/list
- search
	/v1/search/home
- friend_chain
	/v1/friend_chain/feed_user_recommends
- channel
	/v1/channel/list
- location
	/v1/location/nearby
- comment
	/v1/comment/create

- app 版本:https://ota.weibo.cn
	/api/app/version/upgrade
- 图片上传: https://picupload.weibo.com // https://picupload.service.weibo.com
	/interface/upload.php

*/
