package createpost

import (
	"database/sql"

	"github.com/google/uuid"
)

func createPostUseCase(db *sql.DB, postData newPost) error {
	stmt, err := db.Prepare("insert into posts(id, title, content, author) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.New().String(), postData.Title, postData.Content, postData.Author)
	if err != nil {
		return err
	}

	return nil
}
