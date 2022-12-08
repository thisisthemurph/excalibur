package handler

import (
	"errors"
	"excalibur/internal/handler/dto"
	"excalibur/internal/handler/errorhandler"
	"excalibur/internal/handler/hateoas"
	"excalibur/internal/handler/request"
	"excalibur/internal/handler/response"
	"excalibur/internal/models"
	"excalibur/internal/services"
	"fmt"
	"log"
	"net/http"
)

// DataTemplate handler
type DataTemplate interface {
	GetAllDataTemplates(w http.ResponseWriter, r *http.Request)
	GetDataTemplateByID(w http.ResponseWriter, r *http.Request)
	CreateDataTemplate(w http.ResponseWriter, r *http.Request)
	UpdateDataTemplateByID(w http.ResponseWriter, r *http.Request)
	AddNewColumn(w http.ResponseWriter, r *http.Request)
	DeleteDataTemplateByID(w http.ResponseWriter, r *http.Request)
}

type dataTemplate struct {
	eh      errorhandler.ErrorHandler
	log     log.Logger
	service services.DataTemplateService
}

// NewDataTemplateHandler instantiates handler
func NewDataTemplateHandler(s services.DataTemplateService, logger log.Logger) DataTemplate {
	eh := errorhandler.New()
	return &dataTemplate{
		eh:      eh,
		log:     logger,
		service: s,
	}
}

// swagger:route GET /datatemplate datatemplate listDataTemplates
// Return a list of all datatemplates from the database
//
//	Responses:
//		200: dataTemplateListResponse
func (d *dataTemplate) GetAllDataTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := d.service.GetAllDataTemplates()
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	response.Respond(w, templates, http.StatusOK)
}

// swagger:route GET /datatemplate/{id} datatemplate getDataTemplateByID
//
// Gets a DataTemplate by the specified ID
//
//	Responses:
//		200: dataTemplateResponse
func (d *dataTemplate) GetDataTemplateByID(w http.ResponseWriter, r *http.Request) {
	id, ok := request.GetParamFomRequest(r, "id")

	if !ok {
		response.ReturnError(w, errors.New("could not determine url parameter for DataTemplate ID"), http.StatusBadRequest)
		return
	}

	template, err := d.service.GetDataTemplateByID(id)
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	response.Respond(w, template, http.StatusOK)
}

// swagger:route POST /datatemplate datatemplate createDataTemplate
// Create a new datatemplate
//
// responses:
//	200: dataTemplateResponse

func (d *dataTemplate) CreateDataTemplate(w http.ResponseWriter, r *http.Request) {
	template, err := request.GetDtoFromJSONBody[dto.DataTemplateDTO](w, r)
	if err != nil {
		return
	}

	var columns []models.DataTemplateColumn
	newTemplate := models.DataTemplate{
		Name:    template.Name,
		Columns: columns,
	}

	oid, err := d.service.CreateDataTemplate(newTemplate)
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	result := dto.DataTemplateWithHateoasDTO{
		DataTemplateID: oid,
		H: hateoas.H{}.
			WithLink(hateoas.L{
				Href: "datatemplate/" + oid,
				Rel:  "datatemplate",
				Type: "GET",
			}),
	}

	response.Respond(w, result, http.StatusCreated)
}

// swagger:route PUT /datatemplate/{id} datatemplate updateDataTemplate
// Update a products details
//
// responses:
//	200: hateoasWithIdResponse

func (d *dataTemplate) UpdateDataTemplateByID(w http.ResponseWriter, r *http.Request) {
	id, ok := request.GetParamFomRequest(r, "id")
	if !ok {
		response.ReturnError(w, errors.New("could not determine url parameter for DataTemplate ID"), http.StatusBadRequest)
		return
	}

	update, err := request.GetDtoFromJSONBody[dto.DataTemplateDTO](w, r)
	if err != nil {
		return
	}

	err = d.service.UpdateDataTemplateByID(id, update.Name)
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	result := dto.DataTemplateWithHateoasDTO{
		DataTemplateID: id,
		H: hateoas.H{}.
			WithLink(hateoas.L{
				Href: fmt.Sprintf("datatemplate/%v", id),
				Rel:  "datatemplate",
				Type: "GET",
			}),
	}

	response.Respond(w, result, http.StatusOK)
}

// swagger:route POST /datatemplate/{id} datatemplate addNewColumnToDataTemplate
// Update a products details
//
// responses:
//	200: hateoasWithIdResponse

func (d *dataTemplate) AddNewColumn(w http.ResponseWriter, r *http.Request) {
	id, ok := request.GetParamFomRequest(r, "id")
	if !ok {
		response.ReturnError(w, errors.New("could not determine url parameter for DataTemplate ID"), http.StatusBadRequest)
		return
	}

	column, err := request.GetDtoFromJSONBody[dto.DataTemplateColumnDTO](w, r)
	if err != nil {
		return
	}

	dtc := models.DataTemplateColumn{
		Name: column.Name,
	}

	_, err = d.service.AddNewColumn(id, dtc)
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	result := dto.DataTemplateWithHateoasDTO{
		DataTemplateID: id,
		H: hateoas.H{}.
			WithGetLink(hateoas.L{
				Href: fmt.Sprintf("datatemplate/%v", id),
				Rel:  "datatemplate",
			}),
	}

	response.Respond(w, result, http.StatusOK)
}

// swagger:route POST /datatemplate/{id} datatemplate deleteDataTemplate
// Delete the datatemplate associated with the given ID
//
// responses:
//	200: dataTemplateResponse

func (d *dataTemplate) DeleteDataTemplateByID(w http.ResponseWriter, r *http.Request) {
	id, ok := request.GetParamFomRequest(r, "id")

	if !ok {
		response.ReturnError(w, errors.New("could not determine url parameter for DataTemplate ID"), http.StatusBadRequest)
		return
	}

	template, err := d.service.DeleteDataTemplateByID(id)
	if status := d.eh.HandleAPIError(w, err); status != http.StatusOK {
		return
	}

	response.Respond(w, template, http.StatusOK)
}
