package config

type Config struct {
    NumClients uint `json:"num_clients"`
    Requests []Request `json:"requests"`
}

type Request struct {
    Type string `json:"type"`
    URI string `json:"uri"`
}
