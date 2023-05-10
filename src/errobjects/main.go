package errobjects

type ErrorMessage struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Code        int64  `json:"code"`
	CodeStr     string `json:"codeStr"`
}
