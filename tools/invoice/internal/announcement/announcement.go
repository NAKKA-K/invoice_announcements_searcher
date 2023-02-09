package announcement

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

func ToMaps(announcements []Announcement) ([]map[string]interface{}, error) {
	b, err := json.Marshal(announcements)
	if err != nil {
		return nil, fmt.Errorf("fail to marshal [announcements to bytes]: %w", err)
	}

	var documents []map[string]interface{}
	err = json.Unmarshal(b, &documents)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal [announcements bytes to slice map]: %w", err)
	}

	return documents, nil
}

// normalize は検索のために半角全角を統一する関数
// 漢字, ひらがな, カナ, ﾊﾝｶｸｶﾅ => full-width
// alphabet, number         => half-width
func normalize(s string) string {
	return width.Fold.String(s)
}
