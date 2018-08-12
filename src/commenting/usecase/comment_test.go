package usecase_test

import (
	"commenting/interface/repository"
	"commenting/usecase"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"strings"
	"testing"
)

func TestPostComment(t *testing.T) {
	instance, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	defer instance.Close()
	req, err := instance.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := appengine.NewContext(req)

	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)

	type param struct {
		test   string
		pageId string
		name   string
		text   string
	}

	validCases := []*param{
		{
			test:   "normal",
			pageId: "pageId123",
			name:   "hoge",
			text:   "This is comment.",
		},
		{
			test:   "anonymous post",
			pageId: "pageId123",
			name:   "",
			text:   "This is comment.",
		},
		{
			test:   "japanese",
			pageId: "pageId123",
			name:   "おなまえ",
			text:   "これはコメントです",
		},
		{
			test:   "emoji",
			pageId: "pageId123",
			name:   "🤤",
			text:   "😨",
		},
		{
			test:   "long characters",
			pageId: strings.Repeat("a", 64),
			name:   strings.Repeat("a", 20),
			text:   strings.Repeat("a", 1000),
		},
		{
			test:   "long japanese characters",
			pageId: "pageId123",
			name:   strings.Repeat("あ", 20),
			text:   strings.Repeat("あ", 1000),
		},
		{
			test:   "long emoji characters",
			pageId: "pageId123",
			name:   strings.Repeat("🤤", 20),
			text:   strings.Repeat("😨", 1000),
		},
	}

	invalidCases := []*param{
		{
			test:   "empty PageId",
			pageId: "",
			name:   "hoge",
			text:   "This is comment.",
		},
		{
			test:   "empty text",
			pageId: "pageId123",
			name:   "hoge",
			text:   "",
		},
		{
			test:   "long PageId",
			pageId: strings.Repeat("a", 65),
			name:   "name",
			text:   "text",
		},
		{
			test:   "long name",
			pageId: "pageId123",
			name:   strings.Repeat("a", 21),
			text:   "text",
		},
		{
			test:   "long text",
			pageId: "pageId123",
			name:   "name",
			text:   strings.Repeat("a", 1001),
		},
	}

	for _, c := range validCases {
		_, res := u.PostComment(c.pageId, c.name, c.text)
		if res.Code() != usecase.OK {
			t.Errorf("%s: got '%v: %v' want '%v'", c.test, res.Code(), res.Message(), usecase.OK)
		}
	}

	for _, c := range invalidCases {
		_, res := u.PostComment(c.pageId, c.name, c.text)
		if res.Code() != usecase.ErrInvalid {
			t.Errorf("%s: got '%v' want '%v'", c.test, res.Code(), usecase.ErrInvalid)
		}
	}
}
