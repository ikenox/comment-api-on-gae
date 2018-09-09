package domain

type PageID string

func NewPageID(pageID string) PageID {
	return PageID(pageID)
}

type Page struct {
	pageID PageID
}

func NewPage(pageId PageID) *Page {
	return &Page{
		pageID: pageId,
	}
}

func (p *Page) PageId() PageID {
	return p.pageID
}
