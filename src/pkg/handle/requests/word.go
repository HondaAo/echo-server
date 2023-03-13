package requests

type WordRequest struct {
	Levels []uint64 `json:"levels"`
	Amount uint64   `json:"amount"`
}
