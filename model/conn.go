package model

type ConnectionInfo struct {
	Method string `json:"Method"`
	Header string `json:"Header"`
	Path   string `json:"Path"`
	Params string `json:"Params"`
}
