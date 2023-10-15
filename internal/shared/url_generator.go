package shared

import "strconv"

type UrlGenerator struct {
	host              string
	port              uint
	storageFolderName string
	useHttps          bool
}

func NewUrlGenerator(
	host string,
	port uint,
	storageFolderName string,
	useHttps bool,
) *UrlGenerator {
	return &UrlGenerator{
		host,
		port,
		storageFolderName,
		useHttps,
	}
}

func (that *UrlGenerator) GenerateUrlFromObjectPath(objectPath string) string {
	httpPath := "http://"
	if that.useHttps {
		httpPath = "https://"
	}

	return httpPath + that.host + ":" + strconv.Itoa(int(that.port)) + "/" + that.storageFolderName + objectPath
}
