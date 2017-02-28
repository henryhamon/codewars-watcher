package main

type configFile struct {
	Datastore   string `json:"datastore"`
	Ticker      int    `json:"ticker"`
	DatabaseURL string `json:"databaseurl"`
	API         bool   `json:"api"`
}

func defaultConfigFile() configFile {
	return configFile{Datastore: "Mongo", Ticker: 12}
}
