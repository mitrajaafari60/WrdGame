//define dataModel For request/response Model
package DataModel

const (
	Grey   Color = "Grey"
	Green  Color = "Green"
	Yellow Color = "Yellow"
)

type Status string
type Color string

const (
	ERROR   Status = "error"
	SUCCESS Status = "success"
	FAILED  Status = "failed"
)

type GameInput struct {
	Word string `json:"word,omitempty"`
}

type GameResponse struct {
	Value  [WordsLen]Color `json:"resultColores,omitempty"`
	Status Status          `json:"status,omitempty"`
}
