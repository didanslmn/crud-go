package tamplate

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(
		filepath.Join("views", "layout.html"),
		filepath.Join("views", name),
	))

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
