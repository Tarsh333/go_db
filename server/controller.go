package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/tarsh333/go_db/constants"
	"github.com/tarsh333/go_db/model"
	"github.com/tarsh333/go_db/utils"
)

// this will contain crud fxns that will be called from http.handleFunc

func handleReq(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value("params").(model.RequestParams)
	path := filepath.Join("db", params.Cluster, params.Database, filepath.Join(params.CollectionName+".json"))
	mutex := utils.GetOrCreateMutex(path)
	mutex.Lock()
	defer mutex.Unlock()
	if !ok {
		http.Error(w, "Request params not found", http.StatusInternalServerError)
		return
	}
	fmt.Println("params ", params, params.Action)
	if params.Action == "get" {
		json.NewEncoder(w).Encode(GetCollectionData(params))
		return
	}
	if params.Action == "post" {
		json.NewEncoder(w).Encode(AddCollection(params))
		return
	}
	if params.Action == "edit" {
		json.NewEncoder(w).Encode(UpdateCollection(params))
		return
	}
}

func AddCollection(params model.RequestParams) model.Response {
	fmt.Println("add collection")
	if params.Cluster == "" || params.Database == "" || params.CollectionName == "" {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidParams}
	}
	if !utils.IsValidJSON([]byte(params.Data)) {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidDataFormat}
	}
	err := utils.AddFile(filepath.Join("db", params.Cluster, params.Database), filepath.Join(params.CollectionName+".json"), params.Data)
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	return model.Response{Data: "", Success: true, ErrorMessage: ""}
}

func UpdateCollection(params model.RequestParams) model.Response {
	fmt.Println("update collection")
	if params.Cluster == "" || params.Database == "" || params.CollectionName == "" {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidParams}
	}
	if !utils.IsValidJSON([]byte(params.Data)) {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidDataFormat}
	}
	existingData, err := utils.ReadFile(filepath.Join("db", params.Cluster, params.Database, params.CollectionName+".json"))
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	appendedData, err := utils.MergeJSONStrings(existingData, params.Data)
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	fmt.Println(existingData, appendedData)
	err = utils.AddFile(filepath.Join("db", params.Cluster, params.Database), filepath.Join(params.CollectionName+".json"), appendedData)
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	return model.Response{Data: "", Success: true, ErrorMessage: ""}
}

func GetCollectionData(params model.RequestParams) model.Response {
	fmt.Println("get collection")
	// time.Sleep(100 * time.Second)

	if params.Cluster == "" || params.Database == "" || params.CollectionName == "" {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidParams}
	}
	data, err := utils.ReadFile(filepath.Join("db", params.Cluster, params.Database, params.CollectionName+".json"))
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	return model.Response{Data: data, Success: true, ErrorMessage: ""}
}
