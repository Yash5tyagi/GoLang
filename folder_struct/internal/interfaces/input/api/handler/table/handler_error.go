package table

import (
	"krriya/pkg/response"
	"net/http"
)

func (user *Handler) handleError(w http.ResponseWriter, r *http.Request, err error) {

	var apiResponse *response.APIResponse

	switch {

	default:
		// Handle other errors
		apiResponse = response.NewAPIResponse("failed", http.StatusInternalServerError)
		apiResponse.Error = "Something went wrong"
	}

	response.ResponseJSON(w, r, apiResponse)
}
