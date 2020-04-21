package toast

import (
	"database/sql"
)

// Toast represents a Toast model in the database
type Toast struct {
	Toast string
}

// GetRandomToast returns a random toast
func (t *Toast) GetRandomToast(db *sql.DB) error {
	query, err := db.Prepare("SELECT toast FROM chest_toast WHERE explicit=false ORDER BY random()")
	if err != nil {
		return err
	}

	return query.QueryRow().Scan(&t.Toast)
}
