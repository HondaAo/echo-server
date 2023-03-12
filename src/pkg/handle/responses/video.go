package responses

import "github.com/HondaAo/snippet/src/pkg/services/videos/entity"

type Video struct {
	VideoID   string  `json:"video_id"`
	URL       string  `json:"url"`
	Title     string  `json:"title"`
	Start     float64 `json:"start"`
	End       float64 `json:"end"`
	Level     string  `json:"level"`
	Category  string  `json:"category"`
	Display   bool    `json:"display"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewVideoResponse(video *entity.Video) *Video {
	return &Video{
		VideoID:   video.VideoID,
		Title:     video.Title,
		URL:       video.URL,
		Start:     video.Start,
		Category:  getCategory(video.CategoryID),
		End:       video.End,
		Level:     getVideoLevel(uint64(video.Level)),
		Display:   video.Display,
		CreatedAt: video.CreatedAt.Format("2006-01-02"),
		UpdatedAt: video.UpdatedAt.Format("2006-01-02"),
	}
}

func getCategory(categoryID entity.CategoryID) string {
	switch categoryID {
	case entity.Drama:
		return "海外ドラマ"
	case entity.Anime:
		return "アニメ"
	case entity.Economy:
		return "経済学"
	case entity.Geograpy:
		return "地理学"
	case entity.News:
		return "ニュース"
	case entity.Movie:
		return "映画"
	default:
		return "海外ドラマ"
	}
}

func getType(enType entity.EnglishType) string {
	switch enType {
	case entity.America:
		return "アメリカ英語"
	case entity.England:
		return "イギリス英語"
	case entity.Other:
		return "その他"
	default:
		return "アメリカ英語"
	}
}

func getVideoLevel(intLevel uint64) string {
	switch intLevel {
	case 1:
		return "A1"
	case 2:
		return "A2"
	case 3:
		return "B1"
	case 4:
		return "B2"
	case 5:
		return "C1"
	default:
		return "F"
	}
}
