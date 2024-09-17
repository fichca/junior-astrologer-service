package model

type APODResponse struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Date        string `json:"date"`
	Url         string `json:"url"`
}
