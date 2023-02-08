package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/text/width"
)

type Announcement struct {
	RegistratedNumber string `json:"registratedNumber"`
	OriginName        string `json:"name"`
	NormalizedName    string `json:"normalizedName"`
}

func LoadFromJson(name string) ([]Announcement, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("fail to reading file: %w", err)
	}

	var data []Announcement
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("fail to parsing json: %w", err)
	}

	return data, nil
}

func (d *Announcement) UnmarshalJSON(b []byte) error {
	type Alias Announcement

	var aux Alias
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	d.RegistratedNumber = aux.RegistratedNumber
	d.OriginName = aux.OriginName
	d.NormalizedName = normalize(aux.OriginName)

	return nil
}

func ToSliceMap(announcements []Announcement) []map[string]interface{} {
	var documents []map[string]interface{}
	b, _ := json.Marshal(announcements)
	json.Unmarshal(b, &documents)
	return documents
}

// 漢字, ひらがな, カナ, ﾊﾝｶｸｶﾅ => full-width
// alphabet, number         => half-width
func normalize(s string) string {
	return width.Fold.String(s)
}
