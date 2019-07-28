package keystorage

type Backend interface {
	Read(filename string) (map[string]string, error)
	Write(filename string) error
	Add(name, url string) error
}
