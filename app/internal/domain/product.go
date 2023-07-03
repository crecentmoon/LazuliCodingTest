package domain

import "encoding/json"

type Product struct {
	JAN                 string          `json:"jan"`
	ProductName         string          `json:"product_name"`
	Attributes          json.RawMessage `json:"attributes"`
	Maker               string          `json:"maker"`
	Brand               string          `json:"brand"`
	TagsFromDescription []string        `json:"tags_from_description"`
	TagsFromReview      []string        `json:"tags_from_review"`
}
