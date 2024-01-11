package response

type Result struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataType struct {
	Results []Result `json:"results"`
}
type ResponseObj struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   DataType `json:"data"`
}
