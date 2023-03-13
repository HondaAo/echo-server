package responses

type WordResponse struct {
	WordID  uint64 `json:"word_id"`
	Word    string `json:"word"`
	Mean    string `json:"mean"`
	Comment string `json:"comment"`
	Level   uint64 `json:"level"`
}
