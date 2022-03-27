package server

import (
	"html/template"
	"io"

	"github.com/lbrictson/auditmon/web"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func MustNewRenderer() *Renderer {
	r := Renderer{}
	r.templates = template.Must(template.ParseFS(web.Assets, "templates/*.tmpl"))
	return &r
}

type Renderer struct {
	templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	feData := make(map[string]any)
	if data != nil {
		d, ok := data.(map[string]any)
		if ok {
			feData = d
		}
	}
	feData["infoMessage"] = c.Get("infoMessage")
	feData["errorMessage"] = c.Get("errorMessage")
	feData["sessionInfo"] = c.Get("userData")
	err := t.templates.ExecuteTemplate(w, name, feData)
	if err != nil {
		log.Warnf("something went wrong rendering template %v with err %v", name, err.Error())
		return err
	}
	return nil
}
