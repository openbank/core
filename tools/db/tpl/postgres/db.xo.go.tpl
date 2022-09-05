{{ define "db" -}}
type Repository struct {
	db db.DB
}

func New(db db.DB) db.Storage {
	return &Repository{
		db: db, 
	}
}

func (repo *Repository) WithTx(tx *sql.Tx) db.Storage {
	return &Repository{
		db: tx, 
	}
}

var _ db.Storage = new(Repository)
{{- end }}
