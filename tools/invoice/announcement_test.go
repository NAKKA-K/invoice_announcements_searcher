package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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
    "name" : "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
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
			OriginName:        "株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ",
			NormalizedName:    "株式会社CARTA HOLDINGS",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestNormalize(t *testing.T) {
	actual := normalize("株式会社ＣＡＲＴＡ　ＨＯＬＤＩＮＧＳ, CARTA、ｶﾙﾀ 1！")
	expected := "株式会社CARTA HOLDINGS, CARTA、カルタ 1!"

	assert.Equal(t, expected, actual, "日本語は全角、英数字記号は半角になる")
}
