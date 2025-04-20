package Controller

import (
	"encoding/json"
	"net/http"

	"github.com/FACELESS-GOD/URLShortnerService.git/Helper/StructStore"
	Model "github.com/FACELESS-GOD/URLShortnerService.git/Package/Models"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Utility"
)

func AddURL(writer http.ResponseWriter, Req *http.Request) {
	response := StructStore.URLResponse{}

	url_req := StructStore.URLRequest{}

	Utility.ParseBody(Req, url_req)

	new_url, err := Model.AddURLTODb(url_req)
	if err != nil {
		response.IsAnyError = true
		response.ErrorMessages = append(response.ErrorMessages, err.Error())

		res, _ := json.Marshal(response)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(res)
		writer.Header().Set("Content-Type", "application/json")
		return
	}

	response.IsAnyError = true
	response.ErrorMessages = nil
	response.OriginalURL = url_req.URL
	response.NewURL = new_url
	res, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Header().Set("Content-Type", "application/json")
	return

}
