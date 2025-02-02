package data

import (
	"errors"

	"github.com/gocql/gocql"
)

type ShortenRepo struct {
	db *gocql.Session
}

func NewShortenRepo(db *gocql.Session) ShortenRepo {
	return ShortenRepo{
		db: db,
	}
}

func (repo ShortenRepo) GetFullUrl(linkId string) (string, error) {
	iter := repo.db.Query(`
        SELECT fullUrl FROM verkurzen.urls WHERE linkId = ?
    `, linkId).Iter()

    if iter.NumRows() == 0 {
        return "", errors.New("No rows found")
    }

	var fullUrl string
	iter.Scan(&fullUrl)

    if err := iter.Close(); err != nil {
		return "", err
	}

	return fullUrl, nil
}

func (repo ShortenRepo) StoreLink(linkId string, fullUrl string) error {
	id := gocql.TimeUUID()

	err := repo.db.Query(`
        INSERT INTO verkurzen.urls (id, linkId, fullUrl) VALUES (?, ?, ?)
    `, id, linkId, fullUrl).Exec()

	return err
}

func (repo ShortenRepo) Migrate() error {
	err := repo.db.Query(`
        CREATE KEYSPACE IF NOT EXISTS verkurzen WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}
    `).Exec()
	if err != nil {
		return err
	}

	err = repo.db.Query(`
        CREATE TABLE IF NOT EXISTS verkurzen.urls (id UUID PRIMARY KEY, linkId TEXT, fullUrl TEXT)
    `).Exec()
	if err != nil {
		return err
	}

	err = repo.db.Query(`
        CREATE INDEX IF NOT EXISTS url_linkId ON verkurzen.urls (linkId)
    `).Exec()
	if err != nil {
		return err
	}

	return nil
}
