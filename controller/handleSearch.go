package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ParaCAD/ParaCAD-backend/database"
	"github.com/ParaCAD/ParaCAD-backend/utils"
	"github.com/julienschmidt/httprouter"
)

type SearchRequest struct {
	Query              string `json:"query"`
	SearchDescriptions bool   `json:"search_descriptions"`
	Sorting            string `json:"sorting"`
	PageNumber         int    `json:"page_number"`
	PageSize           int    `json:"page_size"`
}

type SearchResponse struct {
	Results []TemplatePreview `json:"results"`
}

type TemplatePreview struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Preview   string  `json:"preview"`
	Created   string  `json:"created"`
	OwnerUUID *string `json:"owner_uuid"`
	OwnerName *string `json:"owner_name"`
}

const maxPageSize = 25

func (c *Controller) HandleSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	request := SearchRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	err = request.Validate()
	if err != nil {
		utils.HandleErr(r, w, http.StatusBadRequest, err)
		return
	}

	searchParameters := database.SearchParameters{
		Query:              request.Query,
		SearchDescriptions: request.SearchDescriptions,
		Sorting:            database.ToSorting(request.Sorting),
		PageNumber:         request.PageNumber,
		PageSize:           request.PageSize,
	}

	results, err := c.db.SearchTemplates(searchParameters)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}

	response := SearchResponse{
		Results: []TemplatePreview{},
	}

	for _, result := range results {
		response.Results = append(response.Results, searchResponseToTemplatePreview(result))
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.HandleErr(r, w, http.StatusInternalServerError, err)
		return
	}
}

func (r *SearchRequest) Validate() error {
	if r.PageNumber <= 0 {
		return errors.New("page number must be greater than 0")
	}
	if r.PageSize <= 0 {
		return errors.New("page size must be greater than 0")
	}
	if r.PageSize > maxPageSize {
		return fmt.Errorf("page size must be less than or equal to %d", maxPageSize)
	}
	return nil
}

func searchResponseToTemplatePreview(result database.SearchResult) TemplatePreview {
	template := TemplatePreview{
		UUID:    result.UUID,
		Name:    result.Name,
		Preview: utils.ValueOrDefault(result.Preview, "not-found.png"),
		Created: result.Created.Format("2006-01-02 15:04"),
	}
	if result.OwnerDeleted == nil {
		template.OwnerUUID = result.OwnerUUID
		template.OwnerName = result.OwnerName
	}
	return template
}
