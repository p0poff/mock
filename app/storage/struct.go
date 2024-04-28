package storage

type Route struct {
	Url     string
	Method  string
	Headers []map[string]string
	Body    string
}

type Setting struct {
	Id     int
	Option string
	Value  string
}

type Routers struct {
	Routes []Route
}

type Settings struct {
	Settings []Setting
}
