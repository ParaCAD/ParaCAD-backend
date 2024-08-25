package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type GetTemplateResponse struct {
	TemplateUUID        uuid.UUID                      `json:"template_uuid"`
	TemplateName        string                         `json:"template_name"`
	TemplateDescription string                         `json:"template_description"`
	TemplatePreview     []byte                         `json:"template_preview"`
	Parameters          []GetTemplateResponseParameter `json:"template_parameters"`

	OwnerUUID uuid.UUID `json:"owner_uuid"`
	OwnerName string    `json:"owner_name"`
}

type GetTemplateResponseParameter struct {
	ParameterDisplayName  string                                  `json:"parameter_display_name"`
	ParameterName         string                                  `json:"parameter_name"`
	ParameterDefaultValue any                                     `json:"parameter_default_value"`
	ParameterConstraints  []GetTemplateResponseParameterConstrain `json:"parameter_constraints"`
}

type GetTemplateResponseParameterConstrain struct {
	Type  constrainType `json:"type"`
	Value any           `json:"value"`
}

type constrainType string

const (
	MinValue  constrainType = "min_value"
	MaxValue  constrainType = "max_value"
	Step      constrainType = "step"
	MinLength constrainType = "min_length"
	MaxLength constrainType = "max_length"
)

func (c *Controller) HandleGetTemplate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateUUID, err := uuid.Parse(p.ByName("uuid"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	template, err := c.db.GetTemplateByUUID(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusNotFound, err)
		return
	}

	owner, err := c.db.GetUserByUUID(template.OwnerUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusConflict, err)
		return
	}

	response := GetTemplateResponse{
		TemplateUUID:        template.UUID,
		TemplateName:        template.Name,
		TemplateDescription: template.Description,
		TemplatePreview:     template.Preview,
		Parameters:          []GetTemplateResponseParameter{},

		OwnerUUID: owner.UUID,
		OwnerName: owner.Username,
	}

	for _, parameter := range template.Parameters {
		response.Parameters = append(response.Parameters, parameterToResponseParameter(parameter))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
}

func parameterToResponseParameter(parameter database.Parameter) GetTemplateResponseParameter {
	responseParameter := GetTemplateResponseParameter{
		ParameterDisplayName: parameter.GetDisplayName(),
		ParameterName:        parameter.GetName(),
	}
	switch parameter.GetType() {
	case database.ParameterTypeString:
		p := parameter.(database.StringParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstrain{
			{
				Type:  MinLength,
				Value: p.MinLength,
			},
			{
				Type:  MaxLength,
				Value: p.MaxLength,
			},
		}
	case database.ParameterTypeInt:
		p := parameter.(database.IntParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstrain{
			{
				Type:  MinValue,
				Value: p.MinValue,
			},
			{
				Type:  MaxValue,
				Value: p.MaxValue,
			},
		}

	case database.ParameterTypeFloat:
		p := parameter.(database.FloatParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
		responseParameter.ParameterConstraints = []GetTemplateResponseParameterConstrain{
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
	case database.ParameterTypeBool:
		p := parameter.(database.BoolParameter)
		responseParameter.ParameterDefaultValue = p.DefaultValue
	}
	return responseParameter
}
