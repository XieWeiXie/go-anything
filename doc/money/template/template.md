<h1 align="center">{{.Title}}</h1>

{{$Talk := "talk"}}

{{ if eq .Type  $Talk}}
<p align="center">
    <a>{{.Owner.Name}} || {{.CreateTime}}</a>
</p>

<div align="center">
    <img src="{{.Owner.AvatarURL}}" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>
{{else}}
<p align="center">
    <a>{{.Question.Owner.Name}} || {{.CreateTime}}</a>
</p>
<div align="center">
    <img src="{{.Question.Owner.AvatarURL}}" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
</div>
{{end}}

{{ if eq .Type  $Talk}}

## 正文

<div>
{{.Text}}
</div>

### 文章配图

<div class="image" align="center">
{{range .Images}}
<img src="{{.URL}}" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>
{{end}}
</div>

{{else}}

## Q&A

### Q
<div class="question">

<div align="center">
<p align="center">
    <a>{{.Question.Owner.Name}}</a>
</p>
<img src="{{.Question.Owner.AvatarURL}}" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
<br>
{{.Question.QuestionText}}
</div>

#### 提问配图

<div class="image" align="center">
{{range .Question.Images}}
<img src="{{.URL}}" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>
{{end}}
</div>
</div>

### A

<div class="answer">
<div align="center">
<p align="center">
    <a>{{.Answer.Owner.Name}}</a>
</p>
<img src="{{.Answer.Owner.AvatarURL}}" width="100" height="100" style="border:1px solid;border-radius:50%; color:#ffffff"/>
<br>
{{.Answer.AnswerText}}
</div>


#### 回答配图

<div class="image" align="center">
{{range .Answer.Images}}
<img src="{{.URL}}" width="33%" height="33%" style="border:1px solid;border-radius:50%; color:#3c3f41"/>
{{end}}
</div>
</div>

{{end}}
## 评论

<div align="left">
<div>
{{ range .Comments}}
<blockquote >
<span> <strong>{{.AuthorName}} : {{.Content}} </strong></span>
</blockquote>
{{end}}
</div>
</div>