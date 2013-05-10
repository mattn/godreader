{{define "entry"}}
{{$root := .Root}}{{$value := .Value}}
<html>
<head>
<meta charset="UTF-8">
<title>GodReader</title>
</head>
<body>
<h4>{{$value.Title | html}}</h4>
{{html $value.Content}}
<hr />
<a accesskey="0" href=".">refresh</a>
<a accesskey="8" href="{{$root}}">back</a>
</body>
</html>
{{end}}
