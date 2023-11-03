package updatepost

import (
	"database/sql"
)

func updatePostUseCase(db *sql.DB, postData updatePost) error {
	stmt, err := db.Prepare("update posts set title = coalesce(nullif(?, ''), title), content = coalesce(nullif(?, ''), content), author = coalesce(nullif(?, ''), author) where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postData.Title, postData.Content, postData.Author, postData.ID)
	if err != nil {
		return err
	}

	return nil
}
