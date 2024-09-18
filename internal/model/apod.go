package model

import (
	"encoding/json"
	"time"
)

type APODClientResponse struct {
	Title       string   `json:"title"`
	Explanation string   `json:"explanation"`
	Copyright   string   `json:"copyright"`
	Date        APODDate `json:"date"`
	Url         string   `json:"url"`
}

type APODDate time.Time

const DateFormat = "2006-01-02"

func (ad *APODDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	t, err := time.Parse(DateFormat, str)
	if err != nil {
		return err
	}

	*ad = APODDate(t)
	return nil
}

func (ad APODDate) MarshalJSON() ([]byte, error) {
	t := time.Time(ad)
	return json.Marshal(t.Format(DateFormat))
}

func NewAPOD(id, title, explanation, copyright string, date *time.Time) *APOD {
	return &APOD{
		Id:          id,
		Title:       title,
		Explanation: explanation,
		Date:        date,
		Copyright:   copyright,
	}
}

type APODResponse struct {
	Title       string     `json:"title"`
	Explanation string     `json:"explanation"`
	Copyright   string     `json:"copyright"`
	Date        *time.Time `json:"date"`
	Url         string     `json:"url"`
}

type APOD struct {
	Id          string     `db:"id"`
	Title       string     `db:"title"`
	Explanation string     `db:"explanation"`
	Copyright   string     `db:"copyright"`
	Date        *time.Time `db:"date"`
}
