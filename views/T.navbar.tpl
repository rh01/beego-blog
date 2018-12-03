{{ define "navbar" }}
<div class="navbar navbar-default">
    <div class="container">
        <a href="/" class="navbar-brand">我的博客</a>
        <ul class="nav navbar-nav">
            <li {{if .IsHome}} class="active"{{ end }}><a href="/">首页</a></li>

            <li {{ if .IsCategory }} class="active" {{ end }}> <a href="/category">分类</a></li>
            <li  {{ if .IsTopic }} class="active"{{ end }}><a href="/topic">文章</a></li>
        </ul>
        <div class="pull-right">
            <ul class="nav navbar-nav">
                {{ if .IsLogin }}
                    <li><a href="/login/exit">退出</a></li>
                {{ else }}
                    <li><a href="/login">管理員登陸</a></li>
                {{ end }}
            </ul>
        </div>
    </div>
</div>
{{ end }}