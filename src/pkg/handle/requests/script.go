package requests

type ScriptRequest struct {
	ScriptID     uint64   `json:"script_id"`
	Text         string   `json:"text"`
	Japanese     string   `json:"japanese"`
	TimeStamp    float64  `json:"timestamp"`
	ScriptIdioms []string `json:"idioms"`
}
