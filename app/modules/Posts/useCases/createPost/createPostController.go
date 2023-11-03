package createpost

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ProgHenrique/api-blog/app/models"
	"github.com/ProgHenrique/api-blog/app/modules"
)

type newPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func CreatePostController(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var postData newPost
	err = json.Unmarshal(body, &postData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = createPostUseCase(db, postData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	createdMessage := models.Response{
		Message: "Post created success!",
	}

	modules.ResponseJSON(res, http.StatusCreated, createdMessage)
}
