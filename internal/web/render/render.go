package render

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	log.Info("Rendering template %s", tmpl)
}
