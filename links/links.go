package links

type Links interface {
	IsVisited(link string) bool
	AreSameDomain(firstLink string, secondLink string) bool
}

type linksImpl struct {
}

func New() Links {
	return &linksImpl{}
}

func (l *linksImpl) IsVisited(link string) bool {
	return false
}

func (l *linksImpl) AreSameDomain(firstLink string, secondLink string) bool {
	return true
}
