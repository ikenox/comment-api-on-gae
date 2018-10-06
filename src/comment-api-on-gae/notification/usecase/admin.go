package usecase

import (
	"comment-api-on-gae/env"
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

var systemAdmin = fmt.Sprintf("ikenox+%s@gmail.com", env.ProjectID)

// TODO CleanArchitectureで書き直し
// 今後ユーザーへの通知とかユーザーごとの通知設定とかを取り扱うのであればやった方良さそう
func NotifyCommentPosted(ctx context.Context, commentId int64, name, text string) {
	msg := &mail.Message{
		Sender:   systemAdmin,
		To:       []string{systemAdmin},
		Subject:  "コメントが投稿されました",
		Body:     fmt.Sprintf("commentId:%d\nname:%s\ntext:%s\n", commentId, name, text),
		HTMLBody: fmt.Sprintf("commentId:%d\nname:%s\ntext:%s\n", commentId, name, text),
	}
	err := mail.Send(ctx, msg)
	if err != nil {
		log.Errorf(ctx, "mail sending error: %v", err.Error())
	}
}

func NotifyCommentDeleted(ctx context.Context, commentId int64, name, text string){
	msg := &mail.Message{
		Sender:   systemAdmin,
		To:       []string{systemAdmin},
		Subject:  "コメントが削除されました",
		Body:     fmt.Sprintf("commentId:%d\nname:%s\ntext:%s\n", commentId, name, text),
		HTMLBody: fmt.Sprintf("commentId:%d\nname:%s\ntext:%s\n", commentId, name, text),
	}
	err := mail.Send(ctx, msg)
	if err != nil {
		log.Errorf(ctx, "mail sending error: %v", err.Error())
	}
}
