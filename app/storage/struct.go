package storage

type Route struct {
	Id      int
	Url     string
	Method  string
	Headers map[string]string
	Body    string
}

type Setting struct {
	Id     int
	Option string
	Value  string
}
