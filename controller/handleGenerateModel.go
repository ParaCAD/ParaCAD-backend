package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ParaCAD/ParaCAD-backend/generator"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/ParaCAD/ParaCAD-backend/utils/logging"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type GenerateModelRequest struct {
	Parameters []GenerateModelParameter `json:"parameters"`
}

type GenerateModelParameter struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type GenerateModelResponse []byte

func (c *Controller) HandleGenerateModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templateUUID, err := uuid.Parse(p.ByName("UUID"))
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	request := GenerateModelRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	template, err := c.db.GetTemplateContentByUUID(templateUUID)
	if err != nil {
		utils.HandleErr(r, w, http.StatusNotFound, err)
		return
	}

	generatorTemplate := generator.FilledTemplate{
		UUID:     template.UUID,
		Template: []byte(template.Template),
		Params:   []generator.Parameter{},
	}

	for _, templateParameter := range template.Parameters {
		requestParameterIdx := slices.IndexFunc(request.Parameters, func(i GenerateModelParameter) bool {
			return i.ParameterName == templateParameter.GetName()
		})
		if requestParameterIdx == -1 {
			utils.HandleErr(r, w, http.StatusBadRequest, fmt.Errorf("missing parameter %s", templateParameter.GetName()))
			return
		}
		err = templateParameter.VerifyValue(request.Parameters[requestParameterIdx].ParameterValue)
		if err != nil {
			utils.HandleErr(r, w, http.StatusBadRequest, err)
			return
		}
		generatorTemplate.Params = append(generatorTemplate.Params, generator.Parameter{
			Key:   templateParameter.GetName(),
			Value: request.Parameters[requestParameterIdx].ParameterValue,
		})
	}

	generatedModel, err := c.generator.GenerateModel(generatorTemplate)
	if err != nil {
		logging.AddErrorToRequestCtx(r, err, http.StatusInternalServerError)
		utils.HandleErr(r, w, http.StatusInternalServerError, fmt.Errorf("failed to generate preview from template: %w", errors.New("generation failed, please check your template code and parameters")))
		return
	}

	utils.SendFile(w, generatedModel, "model.stl")
}
