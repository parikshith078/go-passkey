package util

type DB struct {
	Data []DBItem `json:"data"`
}

type DBItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

