package domain

type PageId int64
type Page struct {
	Entity
	pageId  PageId
	pageUrl *PageUrl
}

func NewPage(pageId PageId, url string) *Page {
	pageUrl := newPageUrl(url)
	return &Page{
		pageId:  pageId,
		pageUrl: pageUrl,
	}
}

func (p *Page) PageId() PageId {
	return p.pageId
}

func (p *Page) PageUrl() *PageUrl {
	return p.pageUrl
}

type PageUrl struct {
	ValueObject
	url string
}

func newPageUrl(url string) *PageUrl {
	return &PageUrl{
		url: url,
	}
}
