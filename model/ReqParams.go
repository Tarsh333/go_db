package model

type RequestParams struct {
	Cluster        string `json:"cluster"`
	Database       string `json:"db"`
	CollectionName string `json:"collection"`
	Action         string `json:"a"`
	Data           string `json:"d"`
}
type Response struct {
	Success      bool   `json:"success"`
	Data         string `json:"data"`
	ErrorMessage string `json:"error"`
}
