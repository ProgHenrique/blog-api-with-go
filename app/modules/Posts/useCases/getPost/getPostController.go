package getpost

import (
	"database/sql"
	"net/http"

	"github.com/ProgHenrique/api-blog/app/models"
	"github.com/ProgHenrique/api-blog/app/modules"
	"github.com/gorilla/mux"
)

func GetPostController(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["_id"]
	if !ok {
		errorMessage := models.Response{
			Message: "Post id is missing!",
		}
		modules.ResponseJSON(res, http.StatusBadRequest, errorMessage)
		return
	}
	post, err := getPostUseCase(db, id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	modules.ResponseJSON(res, http.StatusOK, post)
}
