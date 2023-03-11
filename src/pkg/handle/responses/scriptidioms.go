package responses

type ScriptIdioms struct {
	Idiom    Idioms `json:"idiom"`
	VideoID  string `json:"video_id"`
	ScriptID uint64 `json:"script_id"`
}

type Idioms struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Level   string `json:"level"`
}
