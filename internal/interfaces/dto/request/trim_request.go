package request

type TrimVideoRequest struct {
	TrimStart string `json:"trimStart" default:"00:00:01"`
	TrimEnd   string `json:"trimEnd" default:"00:00:03"`
}
