package domain

type PageId int64
type Page struct {
	pageId  PageId
	pageUrl *PageUrl
}

func NewPage(pageId PageId, pageUrl *PageUrl) *Page {
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

func (p *PageUrl) Url() string {
	return p.url
}

func NewPageUrl(url string) *PageUrl {
	return &PageUrl{
		url: url,
	}
}
