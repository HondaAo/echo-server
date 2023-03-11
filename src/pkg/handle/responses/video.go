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
		Level:     GetVideoLevel(uint64(video.Level)),
		Display:   video.Display,
		CreatedAt: video.CreatedAt.Format("2006-01-02"),
		UpdatedAt: video.UpdatedAt.Format("2006-01-02"),
	}
}

func getCategory(categoryID entity.CategoryID) string {
	switch categoryID {
	case entity.Drama:
		return "ドラマ"
	case entity.Anime:
		return "アニメ"
	case entity.Economy:
		return "経済学"
	case entity.Geograpy:
		return "地理学"
	case entity.Politics:
		return "政治"
	case entity.Movie:
		return "映画"
	default:
		return "ドラマ"
	}
}
