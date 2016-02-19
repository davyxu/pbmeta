package pbmeta

import (
	"strings"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type CommentMeta struct {
	comment *pbprotos.SourceCodeInfo_Location

	trailing string
	leading  string
}

func (self *CommentMeta) TrailingComment() string {
	return self.trailing

}

func (self *CommentMeta) LeadingComment() string {
	return self.leading

}

const oldStyleMarker = "@"

const newStyleMarker = "[tabtoy]"

func (self *CommentMeta) findOldStyleComment() string {
	if pos := strings.Index(self.trailing, oldStyleMarker); pos == 0 {
		return strings.TrimSpace(self.trailing[len(oldStyleMarker):])
	} else if strings.Index(self.leading, oldStyleMarker) == 0 {
		return strings.TrimSpace(self.leading[len(oldStyleMarker):])
	}

	return ""
}

// 解析带有tag的comment, 类似于go结构体中的命名tag
func (self *CommentMeta) ParseTaggedComment() []*TaggedComment {

	commentArray := make([]*TaggedComment, 0)

	// 解析新的样式
	commentArray = parseTaggedComment(self.leading, commentArray)
	commentArray = parseTaggedComment(self.trailing, commentArray)

	// 兼容老的样式
	if oldcomment := self.findOldStyleComment(); oldcomment != "" {

		commentArray = append(commentArray, &TaggedComment{
			Name: "@",
			Data: oldcomment,
		})

	}

	return commentArray
}

func (self *CommentMeta) parse() {
	self.trailing = strings.TrimSpace(self.comment.GetTrailingComments())
	self.leading = strings.TrimSpace(self.comment.GetLeadingComments())

	//log.Printf("<%s> <%s>", self.trailing, self.leading)
}

func newCommentMeta(comment *pbprotos.SourceCodeInfo_Location) CommentMeta {
	cm := CommentMeta{
		comment: comment,
	}

	cm.parse()

	return cm
}
