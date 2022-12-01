// Package routerbuilder responsible for building the routes
package routerbuilder

import (
	"excalibur/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// RouterBuilder object responsible for building routes
type RouterBuilder interface {
	Init() *mux.Router
}

type routerBuilder struct {
	r  *mux.Router
	hc handler.HandlerCollection
}

// New creates a new RouterBuilder instance
func New(hc handler.HandlerCollection) RouterBuilder {
	router := mux.NewRouter()

	return &routerBuilder{
		r:  router,
		hc: hc,
	}
}

func (b *routerBuilder) Init() *mux.Router {
	b.buildDataTemplateRouter()

	return b.r
}

func (b *routerBuilder) buildDataTemplateRouter() {
	h := b.hc.DataTemplateHandler

	b.r.HandleFunc("/datatemplate", h.GetAllDataTemplates).Methods(http.MethodGet)
	b.r.HandleFunc("/datatemplate/{id:[0-9a-f]{24}}", h.GetDataTemplateByID).Methods(http.MethodGet)
	b.r.HandleFunc("/datatemplate", h.CreateDataTemplate).Methods(http.MethodPost)
	b.r.HandleFunc("/datatemplate/{id:[0-9a-f]{24}}", h.UpdateDataTemplateByID).Methods(http.MethodPut)
	b.r.HandleFunc("/datatemplate/{id:[0-9a-f]{24}}", h.DeleteDataTemplateByID).Methods(http.MethodDelete)

	// Column configuration and updates

	b.r.HandleFunc("/datatemplate/{id:[0-9a-f]{24}}/column", h.AddNewColumn).Methods(http.MethodPost)
}
