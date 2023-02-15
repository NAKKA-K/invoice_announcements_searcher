# インボイス公表情報なんとか検索したい

## 準備

- docker が必要
- `./data` に json を落としてくる
  - https://www.invoice-kohyo.nta.go.jp/download/zenken#

## 使い方

起動と json データの取り込み
```
$ make all
go mod download
docker run --cidfile meilisearch.cid \
		--name meilisearch \
		--rm -itd -p 7700:7700 -v \
		~/fluct_backbone/tools/invoice/meili_data:/meili_data getmeili/meilisearch:v0.30
go run ./...
2023/02/09 12:00:12 start to load json
2023/02/09 12:00:39 loaded: ./data/h_all_20230131_002.json, status: SUCCESS, duration 4.2642394s
2023/02/09 12:00:44 loaded: ./data/h_all_20230131_001.json, status: SUCCESS, duration 6.6648279s
2023/02/09 12:00:45 loaded: ./data/h_all_20230131_003.json, status: SUCCESS, duration 6.6648279s
2023/02/09 12:00:46 loaded: ./data/h_all_20230131_004.json, status: SUCCESS, duration 6.6648279s
2023/02/09 12:00:46 loaded: ./data/h_all_20230131_005.json, status: SUCCESS, duration 6.6648279s
2023/02/09 12:00:46 finish loading json
```

検索
```
$ make search Q=fluct
{
  "hits": [
    {
      "registratedNumber": "T8011001056816",
      "name": "株式会社ｆｌｕｃｔ",
      "normalizedName": "株式会社fluct"
    }
  ],
  "query": "\"fluct\"",
  "processingTimeMs": 7,
  "limit": 20,
  "offset": 0,
  "estimatedTotalHits": 1
}
```

## その他

全角半角とかは正規化している
