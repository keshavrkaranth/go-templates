package handlers

import (
	"net/http"
	"templates-practice/pkg/config"
	"templates-practice/pkg/models"
	"templates-practice/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Homepage(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	ipAddr := r.RemoteAddr
	m.App.Session.Put(r.Context(),"IP",ipAddr)

	stringMap["name"] = "Keshav R Karanth "
	stringMap["ip"] = ipAddr
	render.RenderTemplate(w, "homepage.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
