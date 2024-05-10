package urlshortner

type URLData struct {
	ShortURL       string
	OriginalURL    string
	CreationTime   int64
	ExpirationTime int64
	ClickCount     int64
}

type StoreProvider interface {
	SaveURL(urlData URLData)
	UpdateURL(urlData URLData)
	FetchURL(shortURL string)
	IncrementClickCount(shortURL string, addClickCount int64)
}
