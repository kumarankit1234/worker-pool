package downloader

type Downloader interface {
	Download(url string) string
}

type downloaderImpl struct {
}

func New() Downloader {
	return &downloaderImpl{}
}

func (d *downloaderImpl) Download(url string) string {
	return ""
}
