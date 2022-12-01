package dto

// baseDataTemplateColumnDTO represents a column of a DT
type baseDataTemplateColumnDTO struct {
	Name string `json:"name"`
}

// NewDataTemplateColumnDTO for creating/adding new columns
type NewDataTemplateColumnDTO struct {
	baseDataTemplateColumnDTO
}
