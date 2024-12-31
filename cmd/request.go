package main

type Request struct {
	Type     string   `json:"type"`
	Message  string   `json:"message"`
	Token    string   `json:"token,omitempty"`
	Position Position `json:"position,omitempty"`
	Init     Init     `json:"init,omitempty"`
	Lang     string   `json:"lang,omitempty"`
	Limit    int      `json:"limit,omitempty"`
	Length   int      `json:"length,omitempty"`
}

type Init struct {
	Lang    string `json:"lang,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Length  int    `json:"length,omitempty"`
	Trial   int    `json:"trial,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
}
