package models

type ItemResponse struct {
	Item Item `json:"item"`
}

type ItemsResponse struct {
	Categories []string `json:"categories"`
	Items      []Item   `json:"items"`
}
