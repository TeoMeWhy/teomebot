package streamelements

type AddPointsRequest struct {
	UserName string `json:"user"`
	Amount   int64  `json:"amount"`
}
