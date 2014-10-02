package main

import "os"
import "text/template"

const t = 
`{{range $i := .nephews}}Hello {{$i}}
{{end}}`


func main() {
	nephews := []string{"Huey", "Dewey", "Louie"}
	data := map[string]interface{}{
		"nephews": nephews,
	}
	template.Must(template.New("").Parse(t)).Execute(os.Stdout, data)
}
