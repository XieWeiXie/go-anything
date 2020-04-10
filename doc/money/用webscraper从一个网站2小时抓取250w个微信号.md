<h1 align="center">用webscraper从一个网站2小时抓取250w个微信号</h1>




<p align="center">
    <a>明白 || 2018-04-08 08:01:45 &#43;0000 UTC</a>
</p>

<div align="center">
    <img src="https://images.zsxq.com/FjQmHspasTB6fS9i4Psn1Vi_tfMe?e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:G1ott7_PptRi18eZpfiMpU6xbzw=" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>




## 正文

<div>
#思路分享# 
前几天，我用webscraper从一个网站，用了2小时，抓了250w个微信号。

经过去重，去掉无效微信号，剩下9w个微信号。

随机抽样测试过，都是真实有效的用户，不是那种不用审核就通过的微信号，当然从签名也可以看出不是微商，也可能有微商。见下图。

有人问我，从哪个网站抓的这些微信号，抱歉，这个真不能告诉你。这是一个朋友告诉我的，说要保密。也是巧，知道这个网站的不知道webscraper，知道webscraper的不知道这个网站，刚好我两个都知道了，于是就抓到了这些数据。

不过我能分享一下，另一个抓取其他微信号的方式，而且是精准粉。不过这种方式就需要你会webscraper或者编程。

下面是具体方式：
————————
在微博界面，搜索“留下微信号”，会发现很多相关帖子，帖子下面的评论都是微信号。如下图：

这个抓取下来后，根据帖子内容，得到的都是相关行业，人群的精准粉。

我来分享下抓取思路：
1、微博搜索关键字得到的网页，作为一级页面。
2、每篇帖子，作为二级页面。
3、在二级页面，抓取下面所以的评论

抓取下来后，可能会有一些无效评论，这个需要你处理一下。

我大致看了下流程，webscraper是可以做到的，有些细节，比如，搜索结果里，需要用鼠标点击一下才能展开留言，这个我都有讲过，没印象的去我那个免费的webscraper星球翻翻。

思路在这里，可能过程中还会有一些具体问题，再说吧。有兴趣的朋友可以试试。

当然，可能用webscraper不能一次性完美抓取，那就分开吧。

先把各个帖子的链接抓下来，放到一个网页，再到各个帖子抓评论。

这个思路也可以用到其他地方抓不同类型的数据，比如邮箱
————————
咦，最后是不是又要推广爬虫课呢？🙅🏻‍♂️不，有兴趣的按图索骥，可以找到哒！

如果有大佬对这些微信号感兴趣，可以加我微信，我提供一部分供测试。
</div>

### 文章配图

<div class="image" align="center">

</div>


## 评论

<div align="left">
<div>

<blockquote >
<span> <strong>jesse : 不是号多就马上能用的 </strong></span>
</blockquote>

<blockquote >
<span> <strong>鸣略 : 棒棒哒 </strong></span>
</blockquote>

<blockquote >
<span> <strong>ZK🍃 : 赞👍🏻 </strong></span>
</blockquote>

<blockquote >
<span> <strong>陆怀唐 : 软文推销自己的爬虫课水平越来越高 </strong></span>
</blockquote>

</div>
</div>