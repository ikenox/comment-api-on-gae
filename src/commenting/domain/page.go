package domain

type PageID string

func NewPageID(pageID string) PageID {
	return PageID(pageID)
}

type Page struct {
	pageId PageID
}

func NewPage(pageId PageID) *Page {
	return &Page{
		pageId: pageId,
	}
}

func (p *Page) PageId() PageID {
	return p.pageId
}
