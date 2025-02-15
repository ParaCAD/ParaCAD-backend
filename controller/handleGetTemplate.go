package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type GetTemplateResponse struct {
	TemplateUUID        uuid.UUID                      `json:"template_uuid"`
	TemplateName        string                         `json:"template_name"`
	TemplateDescription string                         `json:"template_description"`
	TemplatePreview     string                         `json:"template_preview"`
	Parameters          []GetTemplateResponseParameter `json:"template_parameters"`

	OwnerUUID uuid.UUID `json:"owner_uuid"`
	OwnerName string    `json:"owner_name"`
}

type GetTemplateResponseParameter struct {
	ParameterDisplayName  string                                   `json:"parameter_display_name"`
	ParameterName         string                                   `json:"parameter_name"`
	ParameterType         string                                   `json:"parameter_type"`
	ParameterDefaultValue any                                      `json:"parameter_default_value"`
	ParameterConstraints  []GetTemplateResponseParameterConstraint `json:"parameter_constraints"`
}

type GetTemplateResponseParameterConstraint struct {
	Type  constraintType `json:"type"`
	Value any            `json:"value"`
}

type constraintType string

const (
	MinValue  constraintType = "min_value"
	MaxValue  constraintType = "max_value"
	Step      constraintType = "step"
	MinLength constraintType = "min_length"
	MaxLength constraintType = "max_length"
)

func (c *Controller) HandleGetTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	template, err := c.db.GetTemplateWithOwnerByUUID(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusNotFound, err)
		return
	}

	response := GetTemplateResponse{
		TemplateUUID:        template.UUID,
		TemplateName:        template.Name,
		TemplateDescription: template.Description,
		TemplatePreview:     utils.ValueOrDefault(template.PreviewURL),
		Parameters:          []GetTemplateResponseParameter{},

		OwnerUUID: template.OwnerUUID,
		OwnerName: template.OwnerName,
	}

	for _, parameter := range template.Parameters {
		response.Parameters = append(response.Parameters, parameterToResponseParameter(parameter))
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
}

func parameterToResponseParameter(parameter dbparameter.Parameter) GetTemplateResponseParameter {
	responseParameter := GetTemplateResponseParameter{
		ParameterDisplayName: parameter.GetDisplayName(),
		ParameterName:        parameter.GetName(),
		ParameterType:        parameter.GetType().ParameterType().String(),
	}
	switch parameter.GetType() {
	case dbparameter.ParameterTypeString:
		p := parameter.(dbparameter.StringParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstraint{
			{
				Type:  MinLength,
				Value: p.MinLength,
			},
			{
				Type:  MaxLength,
				Value: p.MaxLength,
			},
		}
	case dbparameter.ParameterTypeInt:
		p := parameter.(dbparameter.IntParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstraint{
			{
				Type:  MinValue,
				Value: p.MinValue,
			},
			{
				Type:  MaxValue,
				Value: p.MaxValue,
			},
		}

	case dbparameter.ParameterTypeFloat:
		p := parameter.(dbparameter.FloatParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstraint{
			{
				Type:  MinValue,
				Value: p.MinValue,
			},
			{
				Type:  MaxValue,
				Value: p.MaxValue,
			},
			{
				Type:  Step,
				Value: p.Step,
			},
		}
	case dbparameter.ParameterTypeBool:
		p := parameter.(dbparameter.BoolParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
	}
	return responseParameter
}
