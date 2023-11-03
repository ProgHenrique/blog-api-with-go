package deletepost

import (
	"database/sql"
)

func deletePostUseCase(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
