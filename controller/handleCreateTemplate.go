package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/auth"
	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/database/dbparameter"
	"github.com/ParaCAD/ParaCAD-backend/generator"
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

	// Validate request (field lengths, parameter types and constraints)
	err = validateCreateTemplateRequest(request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	userUUID, _, err := auth.GetUserIDAndRoleFromRequest(r)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, fmt.Errorf("failed to get user ID from request: %w", err))
		return
	}

	// Generate model from template to ensure it is valid
	defaultFilledTemplate := templateRequestToDefaultFilledTemplate(request)
	_, err = c.generator.GenerateModel(defaultFilledTemplate)
	if err != nil {
		utils.HandleErr(r, w, http.StatusConflict, fmt.Errorf("failed to generate model from template: %w", err))
		return
	}

	// Generate preview
	preview, err := c.generator.GeneratePreview(defaultFilledTemplate)
	if err != nil {
		utils.HandleErr(r, w, http.StatusConflict, fmt.Errorf("failed to generate preview from template: %w", err))
		return
	}

	templateUUID := uuid.New()
	template, err := templateRequestToTemplate(templateUUID, userUUID, request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	// Save template to database and image store
	err = c.imageStore.SaveFile(*template.Preview, preview)
	if err != nil {
		utils.HandleErr(r, w, http.StatusFailedDependency, fmt.Errorf("failed to save template preview: %w", err))
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

func templateRequestToTemplate(templateUUID uuid.UUID, userUUID uuid.UUID, request CreateTemplateRequest) (database.Template, error) {
	template := database.Template{
		UUID:        templateUUID,
		OwnerUUID:   userUUID,
		Name:        request.TemplateName,
		Description: request.TemplateDescription,
		Preview:     utils.GetPtr(templateUUID.String() + ".png"),
		Template:    request.TemplateContent,
		Parameters:  []dbparameter.Parameter{},
	}

	return template, nil
}

func templateRequestToDefaultFilledTemplate(request CreateTemplateRequest) generator.FilledTemplate {
	filledTemplate := generator.FilledTemplate{
		UUID:     uuid.Nil,
		Template: []byte(request.TemplateContent),
	}

	for _, parameter := range request.Parameters {
		filledTemplate.Params = append(filledTemplate.Params, generator.Parameter{
			Key:   parameter.ParameterName,
			Value: parameter.ParameterDefaultValue,
		})
	}

	return filledTemplate
}
