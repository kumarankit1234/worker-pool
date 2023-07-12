package htmlparser

type HtmlParser interface {
	Parse(html string) []string
}

type htmlParserImpl struct {
	links []string
}

func New() HtmlParser {
	return &htmlParserImpl{
		links: []string{
			"abcd", "efgh", "ijkl", "mnol", "thul", "utlasd",
		},
	}
}

func (h *htmlParserImpl) Parse(html string) []string {
	//if rand.Intn(100) < 30 {
	//	return []string{}
	//}
	//r := rand.Intn(len(h.links) - 2)
	return []string{h.links[0], h.links[1]}
}
