package domain

type PageId int
type Page struct {
	Entity
	pageId  *PageId
	pageUrl *pageUrl
}

func NewPage(pageId *PageId, url string) *Page {
	pageUrl := newPageUrl(url)
	return &Page{
		pageId:  pageId,
		pageUrl: pageUrl,
	}
}

type pageUrl struct {
	ValueObject
	url string
}

func newPageUrl(url string) *pageUrl {
	return &pageUrl{
		url: url,
	}
}
