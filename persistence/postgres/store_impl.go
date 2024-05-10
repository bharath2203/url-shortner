package postgres

import (
	"database/sql"

	"github.com/bharath2203/urlshortner"
)

type postgresURLStore struct {
	dbp *dbProvider
}

func NewPostgresURLStore() (*postgresURLStore, error) {
	ddp, err := NewDBProvider()
	if err != nil {
		return nil, err
	}
	return &postgresURLStore{
		dbp: ddp,
	}, nil
}

func (p *postgresURLStore) SaveURL(urlData urlshortner.URLData) {

}

func (p *postgresURLStore) UpdateURL(urlData urlshortner.URLData) {

}

func (p *postgresURLStore) FetchURL(shortURL string) {

}

func (p *postgresURLStore) incrementClickCount(shortURL string, addClickCount int64) {

}

type dbProvider struct {
	endpoint string
	db       *sql.DB
}

func NewDBProvider() (*dbProvider, error) {
	d := &dbProvider{
		endpoint: "connection_string_here",
	}
	err := d.init()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (p *dbProvider) init() error {
	db, err := sql.Open("postgres", p.endpoint)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *dbProvider) insertURLData(urlData urlshortner.URLData) error {
	query := "INSERT INTO url_data (short_url, original_url, creation_time, expiration_time, clicks) values ($1, $2, $3, $4, $5);"
	_, err := p.db.Exec(query, urlData.ShortURL, urlData.OriginalURL, urlData.CreationTime, urlData.ExpirationTime, urlData.ClickCount)
	if err != nil {
		return err
	}
	return nil
}
