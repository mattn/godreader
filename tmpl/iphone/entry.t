{{define "entry"}}
{{$root := .Root}}{{$value := .Value}}
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black">
<title>GodReader</title>
<link rel="stylesheet" href="{{$root}}assets/css/jquery.mobile-1.3.1.min.css" />
<script src="{{$root}}assets/javascript/jquery-1.9.1.min.js"></script>
<script src="{{$root}}assets/javascript/jquery.mobile-1.3.1.min.js"></script>
</head>
<body>
<div data-role="page" data-theme="b" data-fullscreen="true" data-url="{{$root}}{{$value.Id}}">
<div data-role="header" data-theme="b">
<h2>{{$value.Title | html}}</h2>
</div>
<div data-role="content">
{{html $value.Content}}
</div>
<div data-role="footer" data-theme="b" data-position="fixed">
<a accesskey="0" href="{{$root}}{{$value.Id}}">refresh</a>
<a accesskey="8" href="{{$root}}">back</a>
<a accesskey="5" href="{{$value.Link}}">go the URL</a>
</div>
</body>
</html>
{{end}}
