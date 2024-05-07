package utils

import (
	"encoding/json"
	"net/http"

	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
)

func DecodeItemFromRequest(w http.ResponseWriter, r *http.Request) (*itemDomain.Item, error) {
	var item itemDomain.Item
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &item, nil
}
