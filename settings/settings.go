package settings

var Port string
var DBname string

func Load() {
	Port = ":8080"
	DBname = "mediaLibraryDB2"
}
