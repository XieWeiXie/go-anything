<h1 align="center">用Excel做一个生财龙珠管理系统</h1>
<p align="center">
    <a>一条枸杞 / 2019-11-27 16:37:46 &#43;0000 UTC</a>
</p>

<div align="center">
    <img src="https://images.zsxq.com/FtjMAMbo87Y-3GRSTIbxsk_jCEby?e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:bxYZwmmCX_ORB7FKBrFDRDSBaA8=" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>

## 正文

<div>
&lt;e type=&#34;hashtag&#34; hid=&#34;5481558844&#34; title=&#34;#随便聊聊#&#34; /&gt; 
《 小产品尝试 | 用Excel做的生财龙珠管理系统 》

前几天生财社群里龙珠交易频繁，这种虚拟奖品从设定到管理到流通，总是背负了很大的人力成本。一边要记录存量，一边要记录流水，如果有交易，这边减那边加，还挺容易弄错的。

类似这样的需求也还有很多，理发店办卡充次数，小商店想搞个会员系统试试效果，小学老师想弄个小红花清单，社群搞了个积分规则，其实都需要这样的功能。

没错，我做了个Excel模板，用了一些计算公式和宏代码，实现了一个本地化的小型的资产管理系统，完全Hold得住龙珠管理、交易的日常需求，将很多过程都变成自动或半自动处理，使用、教程等见附件和配图哦~

也可以关注下图公众号【一条枸杞】回复『表格』直接下载哈，有使用问题也可以公众号或私信我来解决。

今天主要是想分享下做这个小产品的过程中的一些思考和心路历程。

■ 1.思路

面对这样的需求猜测，一闪而过的念头是「自己做个管理系统吧」！产品思维和技术思维开始在脑子里躁动。

“这做个管理系统稳稳的！小半个月就能开发出雏形，套个免费的模板可以做出很好的视觉效果！”

内心开始有这种声音。但是从理性出发，如果做成这种云产品，程序安全、服务安全、部署方式、迁移、定制化……涉及到了很多很麻烦的、我又不想舍弃的地方。

突然我想起，前几天公布的生财有术龙珠清单，，都是用Excel做的啊。

那我也在Excel上搞一个！

■ 2.产品的程序思维

首先，要从产品的程序思维出发，即便是基于Excel，也要有一定的自动化在里面。

产品的程序思维是什么？就是在你的产品设计过程中，你要假设有一个无敌的程序员在你身边帮你实现一切需求，以最优用户体验为标杆。

“可是我身边没有程序员啊？”，没关系，这只是基础。在这样设定的产品基础下，你的产品有很多无敌好用的功能，这些功能即将受制于你的技术、时间、知识广度等很多因素的影响，将会使易用性为了实现出来而做多次退让。

比如，我一开始会希望这个管理系统，可以在记录交易的过程中，自动生成余额、排行榜、用户列表，可搜索可追溯。这是我能想到的比较好的体验，也是我的产品思维。

但回到技术实现上，如果没有用户表来规范，可能会在交易记录过程中出现误操作，同一个用户的行为记录成了两个用户的。于是这个时候就要用技术思维，把设计数据库那一套经验拿出来，规定好各种所需字段，然后在Excel公式的一些特性上做好字段顺序的排列……这里都是一些技术思路，不赘述。

生财里有技术能力的老哥真的不少，很多时候我们技术人想问题会很基于直观技术实现，而忽略了一些巧劲，就容易陷入懵逼。
产品就是产品，在想产品的时候要无视技术，在想技术的时候要让步产品，有时候就是那份在公司和PM互帮互助奋力想着产品的目的是什么，除了直接实现还有没有其他的方法可以达到同样的效果。

而产品的程序思维，就是让我们短期忘记技术，纯用产品思维把自己要做的小产品达到能想到的最便利体验，然后逐步退让给技术实现；退让到技术可实现之后，再从技术可实现版本从时间精力等方面退让到MVP版本。

当然，最终肯定要从MVP开始“做”，但做小产品的时候，也要让自己时刻拥有一种“未完成感”，这种“未完成感”一边是吊着自己的完成欲望，一边是防止自己一次性达到“完成感”而没有做充足的试错，而搞错了方向。

保证内核足矣，习惯小步快跑。

■ 3.无码思维

直接使用代码来实现各种东西固然很牛逼，体验也棒，但是也要考虑能不能不写代码就完成这件事，或者怎样写很少的代码完成很复杂的事。

让自己的内心逐渐“去代码化”。

很多人都不会写代码，但很多人心里都有自己的产品想做出来，想试试。

能用工具生成就用工具生成，能不写代码就不写代码。很多时候有很多东西都是不用写代码就能完成的，也正因如此我一直在保持【黑科技大全】星球的更新，强调“让有想法的人，不用PS一样做设计，美感不够工具来凑，不写代码照样发布网站，无须必要的时候不去系统学新技能。”

成为一个“无码工具人”，是我一直想做到的。很多事情凭借工具就可以搞定，这是这个时代带来的普惠。见过很多想法多脑洞大的人，总是吐槽自己就缺个技术了，但其实他们想做的很多东西，现有的工具都能解决。

我身边的同事都知道我怼伪需求的时候最喜欢用一句话：“如果这套商业化逻辑需要大量开发资源，就去先试一下跑通流程。我相信凭借业务人员的技能，仅使用微信 淘宝两个APP做技术承载，就能解决很多应用场景。”

所以，这样一个小型的管理系统，使用Excel也行啊，我何苦非要用我Web开发的那一堆技能。

无论是从我的开发速度，还是MVP产品的需求，都是Excel更快。

为了做这个工具，我确实很认真地学了一下VBA。学了多久？不到30分钟。

扫一遍大概的对象方法名称，了解能力之后，其他的问题针对性Google，绝大多数场景都有前人趟过水。

■ 4.不会技术没关系，用搜索解决技术问题

其实很多“不会写代码”的人会相对高估绝大部分技术人，有很多技术人遇到技术问题，也是靠百度谷歌解决的，所以用搜索就能解决很多技术方案的实现过程。

比如我在做这个Excel文件的时候，也遇到了很多障碍，比如我一开始想做到点击一个按钮就能给个表排序，但是我不会相关的代码，甚至当时也没看VBA，怎么办？这就涉及到一些具体的搜索技巧。

这段本来我打算写一下搜索词的演进策略，但是知识的诅咒有点强(讲述得不太好理解)、内容也比较深，可能会看懵，类似的搜索技能我后面会尝试先推出一些好上手的攻略的，比如之前的这一篇 &lt;e type=&#34;web&#34; href=&#34;https://t.zsxq.com/mI6y3nU&#34; title=&#34;https://t.zsxq.com/mI6y3nU&#34; /&gt; 。

非技术人员搜纯技术问题其实是需要积累的，一开始搜不到不要慌，慢慢的你会发现有那么几个网站总能解决你在某一领域的技术问题，后面你可以面向那个网站的站长、作者们的写作习惯，搜一些他们会提到的关键词，命中率更高。

搜索上亘古不变的真理：『你能搜到的，必然是别人公布的』。

■ 5.总结

最近越来越认同这样的道理：

并不是技术成本越高，产品价值就越高。
产品价值也未必完全需要成本来打造。
有些一念之间的可能性和小想法，本身就足够成为一个产品。
做一堆小产品，能让自己的生活更美好。
如果能顺便让别人的生活更好，就更棒了。
能顺便让很多人的生活更好，就是好的产品了。
过程中还能顺便赚钱，那就更快乐了。
—— 这样的事情每天都在我们身边上演。

感谢阅读到这里，关于面向小白的各种黑科技、无代码造物技能等相关的想法和经验，欢迎加我微信聊啊~记得备注【生财有术】~
</div>

## 配图
<div class="image" align="center">

<img src="https://images.zsxq.com/FjKjmbwXY7rEaUigVO2rQsbCcTyg?imageMogr2/auto-orient/thumbnail/800x/format/jpg/blur/1x0/quality/75&amp;e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:UcZgKL29ZiASINc8QW449Gkq2Rg=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

<img src="https://images.zsxq.com/FjIw9X7aHTW0MrZXuLwJUW4Fv1C_?imageMogr2/auto-orient/thumbnail/800x/format/jpg/blur/1x0/quality/75&amp;e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:Ru3ANCTuCoTSteYP9dO86m9kEFQ=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

<img src="https://images.zsxq.com/FpW5aZK-xL0tiVQB94Fwh4CCO2zO?imageMogr2/auto-orient/thumbnail/800x/format/jpg/blur/1x0/quality/75&amp;e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:jlK8U_bxQfgrYQFQOcqtpwqVsT0=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

<img src="https://images.zsxq.com/FhHpidLd7CGTskPbVTTNiOuC48BZ?e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:LzbhgKnHFYYxLCf4GJCk0cZZWWo=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

<img src="https://images.zsxq.com/FgRLhPR-LkMo6rSXRlcKemMWqoSF?e=1590940799&amp;token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:pkS22UnDzdKDEDihUbNe8M2nshw=" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>

</div>

## 评论

<div align="left">
<div>

<blockquote >
<span> <strong>于洋 : 万能的枸杞老师~~ </strong></span>
</blockquote>

<blockquote >
<span> <strong>北城之北 : 枸杞能造原子弹！ </strong></span>
</blockquote>

<blockquote >
<span> <strong>乔里奥 : 哇，厉害！枸杞老师深夜放大招！作为一个（前）程序员，必须为“无码思维”打call，这种思维方式值得每一个程序员用心体会，无码胜有码[强][强][强] </strong></span>
</blockquote>

<blockquote >
<span> <strong>Matrix : 打call </strong></span>
</blockquote>

<blockquote >
<span> <strong>荔枝小一萌 : 我明白你的意思了。这种管理系统的应用场景很宽泛，不仅可以用于奖励积分制、龙珠交易记录，甚至还可以衍生到记录用户在某件商品上的消费次数。类似的高频场景，还有很多，但是我没懂它是怎么实现效率提升的。 </strong></span>
</blockquote>

<blockquote >
<span> <strong>一条枸杞 : 解放了计算过程中的人力，人只需要录入，每月分别的结果，总榜，余额榜都可以一键查看。
很多时候我们记录社区用户积分，只记录结果 ，记录过程太费时间了。
但只有记录了过程，才能让全部操作可追溯。
我是用这个小工具完成了平时应该做但懒得做的那部分呀~ </strong></span>
</blockquote>

<blockquote >
<span> <strong>一条枸杞 : 冲冲冲！ </strong></span>
</blockquote>

<blockquote >
<span> <strong>一条枸杞 : 造这种大件还真不行，造个小型一次性空气冲击炮还是可以的[害羞](矿泉水瓶压缩大法) </strong></span>
</blockquote>

</div>
</div>