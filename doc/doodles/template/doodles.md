<head>
style{
    image {
    
    }
    img {
        width: 33%;
    }
    .info {
        font-siz: 14px;
        color: #333333;
        margin: 5px 5px 5px 5px;
    }
}
</head>
<h1 align="center"> 谷歌涂鸦 </h1>




## {{.Month}} 月

<div class="image">

{{range .Doodles}}
<img src="{{.AlternateUrl}}" alt="{{.ShareText}}"/>
<div class="info"><div class="title">{{.Title}}</div><div class="date">{{.Date | toDateFunc }}</div></div>
{{end}}
</div>








