package storage

type Route struct {
	Id         int
	Url        string
	Method     string
	Headers    map[string]string
	StatusCode int `json:"status_code"`
	Body       string
}

type Setting struct {
	Id     int
	Option string
	Value  string
}
