package responses

type Script struct {
	VideoID   string         `json:"video_id"`
	ScriptID  uint64         `json:"script_id"`
	Text      string         `json:"text"`
	Japanese  string         `json:"japanese"`
	TimeStamp float64        `json:"time_stamp"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Idioms    []ScriptIdioms `json:"idioms"`
}
