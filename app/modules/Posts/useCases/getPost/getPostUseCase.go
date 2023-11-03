package getpost

import (
	"database/sql"

	"github.com/ProgHenrique/api-blog/app/models"
)

func getPostUseCase(db *sql.DB, id string) (*models.Post, error) {
	stmt, err := db.Prepare("select id, title, author, content, created_at, updated_at from posts where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var post models.Post
	err = stmt.QueryRow(id).Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.Created_at, &post.Updated_at)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
