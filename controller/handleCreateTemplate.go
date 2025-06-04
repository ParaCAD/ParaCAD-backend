package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

const (
	MinTemplateNameLength        = 3
	MaxTemplateNameLength        = 100
	MinTemplateDescriptionLength = 10
	MaxTemplateDescriptionLength = 1000
)

type CreateTemplateRequest struct {
	TemplateName        string `json:"template_name"`
	TemplateDescription string `json:"template_description"`
	TemplateContent     string `json:"template_content"`
	// TemplatePreview     string                           `json:"template_preview"` TODO
	Parameters []CreateTemplateRequestParameter `json:"template_parameters"`
}

type CreateTemplateRequestParameter struct {
	ParameterDisplayName  string                                     `json:"parameter_display_name"`
	ParameterName         string                                     `json:"parameter_name"`
	ParameterType         string                                     `json:"parameter_type"`
	ParameterDefaultValue string                                     `json:"parameter_default_value"`
	ParameterConstraints  []CreateTemplateRequestParameterConstraint `json:"parameter_constraints"`
}

type CreateTemplateRequestParameterConstraint struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type CreateTemplateResponse struct {
	TemplateUUID uuid.UUID `json:"template_uuid"`
}

func (c *Controller) HandleCreateTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	request := CreateTemplateRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = validateCreateTemplateRequest(request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("Creating template:\n[name]:\t%s\n[desc]:\t%s\n[cont]:\t%s\n", request.TemplateName, request.TemplateDescription, request.TemplateContent)

	templateUUID := uuid.New()

	template, err := templateRequestToTemplate(templateUUID, request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = c.db.CreateTemplate(template)
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, err)
		return
	}

	response := CreateTemplateResponse{
		TemplateUUID: template.UUID,
	}

	json.NewEncoder(w).Encode(response)
}

func validateCreateTemplateRequest(request CreateTemplateRequest) error {
	if len(request.TemplateName) < MinTemplateNameLength {
		return fmt.Errorf("template name must be at least %d characters long", MinTemplateNameLength)
	}
	if len(request.TemplateName) > MaxTemplateNameLength {
		return fmt.Errorf("template name must not exceed %d characters", MaxTemplateNameLength)
	}

	if len(request.TemplateDescription) < MinTemplateDescriptionLength {
		return fmt.Errorf("template description must be at least %d characters long", MinTemplateDescriptionLength)
	}
	if len(request.TemplateDescription) > MaxTemplateDescriptionLength {
		return fmt.Errorf("template description must not exceed %d characters", MaxTemplateDescriptionLength)
	}

	if len(request.TemplateContent) == 0 {
		return fmt.Errorf("template content must not be empty")
	}

	// TODO: pre-validate template parameters

	return nil
}

func templateRequestToTemplate(templateUUID uuid.UUID, request CreateTemplateRequest) (database.Template, error) {
	template := database.Template{
		UUID:        templateUUID,
		Name:        request.TemplateName,
		Description: request.TemplateDescription,
		Template:    request.TemplateContent,
		Parameters:  []dbparameter.Parameter{},
	}

	return template, nil
}
