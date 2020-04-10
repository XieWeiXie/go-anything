<h1 align="center">如何获取到张小龙微信ID</h1>
<p align="center">
    <a>Alexander.* / 2017-03-23 11:21:52 &#43;0000 UTC</a>
</p>

<div align="center">
    <img src="https://images.zsxq.com/Frk84GenLwn4hcEItVn69VPvkO2O?e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:9OZo0rMdW7LxXBjDeujLq8AkJBU=" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>

## 正文

<div>
这个分享偏技术向。

我两年半前发过一篇简单科普文，介绍了微信公众平台图文消息的链接组成

&lt;e type=&#34;web&#34; title=&#34;解读微信公众平台图文消息的链接组成&#34; href=&#34;https://mp.weixin.qq.com/s/laVLXBVG_qNWuCdoDk8AUg&#34; /&gt;

在这篇文章里，我没说的是：__biz这个参数，其实就是base64编码，经过base64解码之后能得到公众平台的原始ID，
那么这篇文章里还说了当时获取二维码的办法：
&lt;e type=&#34;web&#34; title=&#34;http://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;s...&#34; href=&#34;https://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;size=102&amp;__biz=MzA5Njg3MjAzOA==&#34; /&gt;

改变size即可改变二维码的大小。

这个二维码获取的办法一直到一个月以前都是有效的，但是现在失效了，我认为跟我之前的分享有关系。

在之前，经过分析知道公众号的原始ID实际上就是数字，而微信个人账号的id为uin，经过base64解码之后原始ID也是数字，
在最开始，微信个人账号和公众账号的原始ID没有分开，那么这个用来获取公众号二维码的办法也可以同样用来获取微信个人账号二维码。

再经过简单的几次尝试(用猜测的原始ID去base64编码得到biz，尝试用这个biz生成的二维码是否可以正常扫码)可以知道：微信的个人账号的ID是从100001开始的，6位数，而qq是5位数，从10001开始。

尝试的结果：
100002: 张文瑞
100003: 刘乐君 （估计弃用了）
100004: Franklin 人在杭州
100005: 微信团队 ID：weixin
100006 ～ 1000011 位置都是广州的
1000011 是张小龙的微信
下面的就没有再尝试了。

我发现这个推导过程是在一年前，直到上个月的一次分享才顺带介绍了这件事，然后放出了通过这种办法推导出的张小龙的个人微信号二维码，
现在如果再去扫码会出现下面图片里的提示。


而且&lt;e type=&#34;web&#34; title=&#34;http://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;s...&#34; href=&#34;https://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;size=102&amp;__biz=MzA5Njg3MjAzOA==&#34; /&gt; 
这种办法获取微信公众号二维码也获取不到了，返回：
{
ret: -2,
errmsg: &#34;get qrcode failed&#34;,
cookie_count: 1
}
需要加上公众平台文章参数才可以获取：
&lt;e type=&#34;web&#34; title=&#34;https://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;...&#34; href=&#34;https://mp.weixin.qq.com/mp/qrcode?scene=10000004&amp;size=102&amp;__biz=MzA5Njg3MjAzOA==&amp;mid=201308885&amp;idx=1&amp;sn=b4509bb0be258a0280a19fdd0b692ed6&amp;send_time=&#34; /&gt;


另外如果你用张文瑞去Google搜索，会搜到几篇不错的微信相关架构的分享文章，建议阅读。
</div>

## 配图
<div class="image" align="center">

<img src="https://images.zsxq.com/Fg6mf0qw4vDrq5_SmPZZT0AtbFHN?imageMogr2/auto-orient/thumbnail/800x/format/jpg/blur/1x0/quality/75&amp;e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:mP15Ead8kwZ5VOXaD2A2pEL4BGo=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

</div>

## 评论

<div align="left">
<div>

<blockquote >
<span> <strong>亦仁 : 这是一篇分享如何获取到张小龙微信ID的神帖 </strong></span>
</blockquote>

<blockquote >
<span> <strong>刘泽君 : 请问圈子里的内容可以截图分享到朋友圈么？ </strong></span>
</blockquote>

<blockquote >
<span> <strong>亦仁 : 直接有分享链接功能，不要截图 </strong></span>
</blockquote>

<blockquote >
<span> <strong>刘泽君 : OK，明白。 </strong></span>
</blockquote>

</div>
</div>