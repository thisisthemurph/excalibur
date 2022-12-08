package dto

// DataTemplateColumnDTO represents a column of a DT
// swagger:model
type DataTemplateColumnDTO struct {
	// the name of the column
	//
	// required: true
	Name string `json:"name"`
}
