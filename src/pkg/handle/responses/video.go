package responses

type Video struct {
	VideoID   string  `json:"video_id"`
	URL       string  `json:"url"`
	Start     float64 `json:"start"`
	End       float64 `json:"end"`
	Level     string  `json:"level"`
	Display   bool    `json:"display"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
