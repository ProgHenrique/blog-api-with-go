package deletepost

import (
	"database/sql"
	"net/http"

	"github.com/ProgHenrique/api-blog/app/models"
	"github.com/ProgHenrique/api-blog/app/modules"
	"github.com/gorilla/mux"
)

func DeletePostController(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["_id"]
	if !ok {
		errorMessage := models.Response{
			Message: "Post id is missing!",
		}
		modules.ResponseJSON(res, http.StatusBadRequest, errorMessage)
		return
	}

	err := deletePostUseCase(db, id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	deleteMessage := models.Response{
		Message: "Post deleted sucess!",
	}

	modules.ResponseJSON(res, http.StatusOK, deleteMessage)
}
