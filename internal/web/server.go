package server

import (
	"html/template"
	"net/http"

	repository "github.com/agschrei/cyclonedx-viewer/internal/repository"
	models "github.com/agschrei/cyclonedx-viewer/pkg/models"
	log "github.com/sirupsen/logrus"
)

var bomContainer models.CycloneDxBomContainer

func Serve(path string) {
	bomContainer = models.ParseCycloneDxBom(path)
	static := repository.FileRepository.Static
	http.Handle("/static/", http.FileServer(http.FS(static)))

	http.HandleFunc("/", RenderMainPage)
	http.HandleFunc("/raw/", RenderRawPage)
	log.Info("Serving page at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RenderRawPage(w http.ResponseWriter, r *http.Request) {
	templates := repository.FileRepository.Templates

	t, _ := template.ParseFS(templates, "raw.page.go.tpl")
	t.ParseFS(templates, "layouts/*", "partials/*")

	err := t.Execute(w, bomContainer.BomData)

	if err != nil {
		log.Info(err)
	}

}

func RenderMainPage(w http.ResponseWriter, r *http.Request) {
	templates := repository.FileRepository.Templates

	t, _ := template.ParseFS(templates, "home.page.go.tpl")
	t.ParseFS(templates, "layouts/*", "partials/*")
	t.Execute(w, bomContainer.Bom)

}
