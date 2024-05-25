package request

type BreakPoint struct {
	FileName string `json:"fileName" form:"fileName"`
	FileMd5  string `json:"fileMd5" form:"fileMd5"`
}
