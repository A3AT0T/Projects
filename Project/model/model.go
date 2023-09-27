package model

type MyUser struct {
	Name   string   `json:"name"`
	Age    uint64   `json:"age"`
	Active bool     `json:"active"`
	Mass   float64  `json:"mass"`
	Books  []string `json:"books"`
}

type User struct {
	Name   string
	Age    int64
	Active bool
	Mass   float64
	Books  []string
}
