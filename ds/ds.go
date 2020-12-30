package ds

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Postrequest struct {
	Name    string `json:"name"`
	Roll_no string `json:"roll_no"`
}
