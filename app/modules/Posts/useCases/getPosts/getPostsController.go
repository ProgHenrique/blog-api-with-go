package getposts

import (
	"database/sql"
	"net/http"

	"github.com/ProgHenrique/api-blog/app/modules"
)

func GetPostsController(db *sql.DB, res http.ResponseWriter, req *http.Request) {
	posts, err := getPostsUseCase(db)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	modules.ResponseJSON(res, http.StatusOK, posts)
}
