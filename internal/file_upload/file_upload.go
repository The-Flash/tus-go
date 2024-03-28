package fileupload

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type FileUpload struct{}

func (f *FileUpload) Create(fid string, size int64, sizeIsDeffered bool, offset int64, metadata map[string]string) error {
	fileInfo := FileInfo{
		Fid:            fid,
		Size:           size,
		SizeIsDeffered: sizeIsDeffered,
		Offset:         offset,
		Metadata:       metadata,
	}
	err := create(fileInfo)
	if err != nil {
		return err
	}
	return nil
}

func create(f FileInfo) error {
	file, err := os.Create(path.Join("..", "files", fmt.Sprintf("%s.json", f.Fid)))
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(f)
	if err != nil {
		return err
	}
	return nil
}
