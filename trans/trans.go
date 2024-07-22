package trans

import (
	"encoding/json"
	"io"
	"net/http"
)

type Book struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json: price`
}

func AddBook(w http.ResponseWriter, r *http.Request) (Book, error) {
	reqB, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return Book{}, err
	}

	b := Book{}

	if err = json.Unmarshal(reqB, &b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	return b, nil

}
