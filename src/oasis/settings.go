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
- annotation
	/v1/annotation/search

- app 版本:https://ota.weibo.cn
	/api/app/version/upgrade
- 图片上传: https://picupload.weibo.com // https://picupload.service.weibo.com
	/interface/upload.php
https://picupload.weibo.com/interface/upload.php?appid=3157701306&cs=3430395255&ent=xiancheng&file_source=53&gsid=_2A9Y6hnV0DIVOw5wYhtwUazlgUavKDdCBNrCioHznUhyKVZN6hIl6sD-ykK3SLmiw9wUZaarhL63iWsw84mgxWgid&ori=1&uid=1948244870
curl -H 'Host: oasis.weibo.cn' -H 'gsid: rWf/skaVKcSmS1v+WmGQmFFI5rCwUfhG7kx9FomaSbnUzxZZw6ViLbReM6RRMZRY5d8XhiG7PlDBK+wW6ApbU2dxl1gxx9OX41NSAbEqKIONpCIgJM9FBYKAypZZjwxQ' -H 'authorization: Bearer' -H 'x-log-uid: 1948244870' -H 'x-sessionid: 1756e10c-1a3d-419f-bf89-ed55daee013a' -H 'aid: 01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.' -H 'user-agent: smartisan-OD103__oasis__1.4.17__Android__Android7.1.1' -H 'content-type: text/plain; charset=utf-8' --data-binary '{"text":"图图","medias":[{"pid":"741fdb86gy3g8exyys66wj20u00u0tpf","width":1080,"height":1080}]}' --compressed 'https://oasis.weibo.cn/v1/status/create?ua=smartisan-OD103__oasis__1.4.17__Android__Android7.1.1&wm=9009_90025&aid=01A0QclqWHjYDiPccIVGsKkq8VuZpCUIaUQOAW6PCht1I4BUo.&cuid=1948244870&sign=f772f7c942753c3e7e5f1d75878ac042&type=1&timestamp=1572322112022&cfrom=2899595010&poiid=8008631000000000000&title=%E5%9B%BE%E5%9B%BE&weibo=true&version=1.4.17&noncestr=8977xV620e15Bka96584gk89UZ2841&platform=ANDROID'
*/
