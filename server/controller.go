package server

import (
	"fmt"
	"path/filepath"

	"github.com/tarsh333/go_db/constants"
	"github.com/tarsh333/go_db/model"
	"github.com/tarsh333/go_db/utils"
)

// this will contain crud fxns that will be called from http.handleFunc

func AddCollection(params model.RequestParams) model.Response {
	if params.Cluster == "" || params.Database == "" || params.CollectionName == "" {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidParams}
	}
	if !utils.IsValidJSON([]byte(params.Data)) {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidDataFormat}
	}
	err := utils.AddFile(filepath.Join("db", params.Cluster, params.Database), filepath.Join(params.CollectionName), params.Data)
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	return model.Response{Data: "", Success: true, ErrorMessage: ""}
}

func UpdateCollection(params model.RequestParams) model.Response {
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
	if params.Cluster == "" || params.Database == "" || params.CollectionName == "" {
		return model.Response{Data: "", Success: false, ErrorMessage: constants.Constants.InvalidParams}
	}
	data, err := utils.ReadFile(filepath.Join("db", params.Cluster, params.Database, params.CollectionName))
	if err != nil {
		return model.Response{Data: "", Success: false, ErrorMessage: err.Error()}
	}
	return model.Response{Data: data, Success: true, ErrorMessage: ""}
}
