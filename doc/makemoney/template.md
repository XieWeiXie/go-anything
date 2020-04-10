<h1 align="center">{{.Title}}</h1>
<p align="center">
    <a>{{.Owner.Name}} / {{.CreateTime}}</a>
</p>

<div align="center">
    <img src="{{.Owner.AvatarURL}}" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>

## 正文

<div>
{{.Text}}
</div>

## 配图
<div class="image" align="center">
{{range .Images}}
<img src="{{.URL}}" width="200" height="200" style="border:1px solid;border-radius:50%; color:#ffffff"/>
{{end}}
</div>

## 评论

<div align="left">
<div>
{{ range .Comments}}
<blockquote >
<span> <strong>{{.AuthorName}} // {{.Content}} </strong></span>
</blockquote>
{{end}}
</div>
</div>