package main

import "strings"

var tmplstr = strings.TrimSpace(`
<!DOCTYPE html>
<html>
<head>
<title>{{ .title }}</title>
<style>
#pics {
  line-height: 0;
  -webkit-column-count: 5;
  -webkit-column-gap: 0px;
  -moz-column-count: 5;
  -moz-column-gap: 0px;
  column-count: 5;
  column-gap: 0px;
}

#pics img {
  width: 100% !important;
  height: auto !important;
}

@media (max-width: 1200px) {
  #pics {
  -moz-column-count: 4;
  -webkit-column-count: 4;
  column-count: 4;
  }
}
@media (max-width: 1000px) {
  #pics {
  -moz-column-count: 3;
  -webkit-column-count: 3;
  column-count: 3;
  }
}
@media (max-width: 800px) {
  #pics {
  -moz-column-count: 2;
  -webkit-column-count: 2;
  column-count: 2;
  }
}
@media (max-width: 400px) {
  #pics {
  -moz-column-count: 1;
  -webkit-column-count: 1;
  column-count: 1;
  }
}
</style>
</head>
<body>

<section id="pics">
{{ range $pic := .pics -}}
<a href="{{.Source}}"><img src="{{.Source}}" height={{.Height}} width={{.Width}}/></a>
{{ end -}}
</section>

</body>
</html>
`)
