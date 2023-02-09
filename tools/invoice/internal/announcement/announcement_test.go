package announcement

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnouncement_UnmarshalJSON(t *testing.T) {
	input := `
[
  {
    "sequenceNumber" : "184764",
    "registratedNumber" : "T6011001033049",
    "process" : "01",
    "correct" : "0",
    "kind" : "2",
    "country" : "1",
    "latest" : "1",
    "registrationDate" : "2023-10-01",
    "updateDate" : "2022-05-30",
    "disposalDate" : "",
    "expireDate" : "",
    "address" : "東京都渋谷区道玄坂１丁目２１番１号渋谷ソラスタ１５階",
    "addressPrefectureCode" : "13",
    "addressCityCode" : "113",
    "addressRequest" : "",
    "addressRequestPrefectureCode" : "",
    "addressRequestCityCode" : "",
    "kana" : "",
    "name" : "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ, CARTA、ｶﾙﾀ 1！",
    "addressInside" : "",
    "addressInsidePrefectureCode" : "",
    "addressInsideCityCode" : "",
    "tradeName" : "",
    "popularName_previousName" : ""
  }
]
`
	var actual []Announcement
	if err := json.Unmarshal([]byte(input), &actual); err != nil {
		t.Error(err)
	}

	expected := []Announcement{
		{
			RegistratedNumber: "T6011001033049",
			OriginName:        "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ, CARTA、ｶﾙﾀ 1！",
			NormalizedName:    "株式会社CARTA HOLDINGS, CARTA、カルタ 1!",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestToMaps(t *testing.T) {
	input := []Announcement{
		{
			RegistratedNumber: "T6011001033049",
			OriginName:        "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
			NormalizedName:    "株式会社CARTA HOLDINGS",
		},
		{
			RegistratedNumber: "T6011001033049",
			OriginName:        "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
			NormalizedName:    "株式会社CARTA HOLDINGS",
		},
	}

	actual, err := ToMaps(input)
	if err != nil {
		t.Error(err)
	}

	expected := []map[string]interface{}{
		{
			"registratedNumber": "T6011001033049",
			"name":              "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
			"normalizedName":    "株式会社CARTA HOLDINGS",
		},
		{
			"registratedNumber": "T6011001033049",
			"name":              "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
			"normalizedName":    "株式会社CARTA HOLDINGS",
		},
	}
	assert.Equal(t, expected, actual)
}
