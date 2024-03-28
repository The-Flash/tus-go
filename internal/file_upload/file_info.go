package fileupload

type Metadata map[string]string

type FileInfo struct {
	Fid            string   `json:"fid"`
	Size           int64    `json:"size"`
	SizeIsDeffered bool     `json:"sizeIsDeffered"`
	Offset         int64    `json:"offset"`
	Metadata       Metadata `json:"metadata"`
}
