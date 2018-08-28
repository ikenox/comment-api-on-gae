package comment

import "time"

type ReplyID int64

type Reply struct {
	replyID     ReplyID
	text        string
	commenterID CommenterID
	repliedAt   time.Time
	commentID   CommentID
}

func newReply(replyID ReplyID, commentID CommentID, commenterID CommenterID, text string, repliedAt time.Time) *Reply {
	return &Reply{
		replyID:     replyID,
		text:        text,
		commenterID: commenterID,
		repliedAt:   repliedAt,
		commentID:   commentID,
	}
}
