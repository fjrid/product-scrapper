package repository

import (
	"os"
)

func (r *RepoImpl) CreateTable() (err error) {
	qBytes, err := os.ReadFile("./migration/up.sql")
	if err != nil {
		return
	}

	_, err = r.db.Exec(string(qBytes))
	if err != nil {
		return
	}

	return
}

func (r *RepoImpl) Droptable() (err error) {
	qBytes, err := os.ReadFile("./migration/down.sql")
	if err != nil {
		return
	}

	_, err = r.db.Exec(string(qBytes))
	if err != nil {
		return
	}

	return
}
