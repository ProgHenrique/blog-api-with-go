package updatepost

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ProgHenrique/api-blog/app/models"
	"github.com/ProgHenrique/api-blog/app/modules"
	"github.com/gorilla/mux"
)

type updatePost struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func UpdatePostController(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["_id"]
	if !ok {
		errorMessage := models.Response{
			Message: "Post id is missing!",
		}
		modules.ResponseJSON(res, http.StatusBadRequest, errorMessage)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var postData updatePost
	postData.ID = id
	err = json.Unmarshal(body, &postData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = updatePostUseCase(db, postData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedMessage := models.Response{
		Message: "Post updated sucess!",
	}

	modules.ResponseJSON(res, http.StatusOK, updatedMessage)
}
