package renderer

import (
	"html/template"
	"net/http"

	decoder "../decoder"
)

var HtmlTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>CYOA</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Lines}}
          <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>
`

func RenderArc(w http.ResponseWriter, arc decoder.Arc) (err error) {
	tpl := template.Must(template.New("").Parse(HtmlTemplate))
	return tpl.Execute(w, arc)
}
