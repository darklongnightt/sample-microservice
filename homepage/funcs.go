package homepage

import (
	"encoding/csv"
	"io"
	"mime/multipart"
)

func (h *Handlers) getProfile() *Profile {
	return &Profile{"Xavier", []string{"Calisthenics", "Coding"}}
}

func (h *Handlers) readFile(file multipart.File) error {
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			return err
		}
		if err == io.EOF {
			break
		}
		h.logger.Println(record)
	}
	return nil
}
