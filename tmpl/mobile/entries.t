{{define "entries"}}
{{$root := .Root}}{{$value := .Value}}
<html>
<head>
<meta charset="UTF-8">
<title>GodReader</title>
</head>
<body>
<h4>GodReader</h4>
{{range $entry := $value}}<a href="{{urlquery $entry.Link}}">{{$entry.Site | html}} {{$entry.Title | html}}</a><br />
{{end}}
<hr />
<a accesskey="0" href=".">refresh</a>
</body>
</html>
{{end}}
