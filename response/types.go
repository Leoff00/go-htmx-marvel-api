package response

type Thumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type Result struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Thumbail    Thumbnail `json:"thumbnail"`
}

type DataType struct {
	Results []Result `json:"results"`
}
type ResponseObj struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   DataType `json:"data"`
}
