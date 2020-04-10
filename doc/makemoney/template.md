<h1 align="center">{{.Title}}</h1>
<p align="center">
    <a>{{.Owner.Name}} / {{.CreateTime}}</a>
</p>

<div align="center">
    <img src="{{.Owner.AvatarURL}}" width="33%" height="33%"/>
</div>

<div>
{{.Text}}
</div>

<div class="image" align="center">
{{range .Images}}
<img src="{{.URL}}" width="33%" height="33%"/>
{{end}}
</div>

<div align="left">
<div>
{{ range .Comments}}
<blockquote >
<span> <strong>{{.AuthorName}} // {{.Content}} </strong></span>
</blockquote>
{{end}}
</div>
</div>