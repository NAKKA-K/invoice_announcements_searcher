package announcement

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/width"
)

type Announcement struct {
	RegistratedNumber string `json:"registratedNumber" mapstructure:"registratedNumber"`
	OriginName        string `json:"name" mapstructure:"name"`
	NormalizedName    string `json:"-" mapstructure:"normalizedName"`
}

type Document map[string]any

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

func ToJSONMaps(announcements []Announcement) ([]Document, error) {
	documents := make([]Document, 0, len(announcements))
	for _, a := range announcements {
		var doc Document
		err := mapstructure.Decode(a, &doc)
		if err != nil {
			return nil, fmt.Errorf("fail to struct -> map: %w", err)
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

// normalize は検索のために半角全角を統一する関数
// 漢字, ひらがな, カナ, ﾊﾝｶｸｶﾅ => full-width
// alphabet, number         => half-width
func normalize(s string) string {
	return width.Fold.String(s)
}
