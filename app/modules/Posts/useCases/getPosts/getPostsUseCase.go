package getposts

import (
	"database/sql"

	"github.com/ProgHenrique/api-blog/app/models"
)

func getPostsUseCase(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query("select id, title, author, content, created_at, updated_at from posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.Title, &p.Author, &p.Content, &p.Created_at, &p.Updated_at)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}
