package file

import (
	"encoding/json"
	"fmt"
	"io"
	"kssyncservice_go/config"
	"kssyncservice_go/pkg/res"
	"net/http"
	"time"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type FileHandlerDeps struct {
	Config *config.Config
}

type FileHandler struct {
	Config *config.Config
}

func NewFileHandler(router *http.ServeMux, deps FileHandlerDeps) {
	handler := &FileHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/file", handler.getFile())
}

func (handler *FileHandler) getFile() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		conf := handler.Config.ExternalFileConfig
		response, err := http.Get(fmt.Sprintf("http://%s:%s%s", conf.FILE_SERVICE_HOST, conf.FILE_SERVICE_PORT, conf.FILE_SERVICE_METHOD))
		if err != nil {
			fmt.Println(time.Now(), " - error fetch data!")
		}
		defer response.Body.Close()

		reader := transform.NewReader(response.Body, charmap.Windows1251.NewDecoder())

		body, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println(time.Now(), " - error read data!")
		}

		var data FileData
		json.Unmarshal(body, &data)

		var sliceOfMaps []map[string]interface{}

		for _, row := range data.Table_Data {
			result := make(map[string]interface{})
			for idx, column := range data.Columns {
				result[column] = row[idx]
			}
			sliceOfMaps = append(sliceOfMaps, result)
		}

		res.Json(w, sliceOfMaps, 200)
	}
}
