package pbmeta

import (
	"strings"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type TaggedComment struct {
	Name string
	Data string
}
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

const simpleStyleMarker = "@"

func (self *CommentMeta) findSimpleStyleComment() string {
	if pos := strings.Index(self.trailing, simpleStyleMarker); pos == 0 {
		return strings.TrimSpace(self.trailing[len(simpleStyleMarker):])
	} else if strings.Index(self.leading, simpleStyleMarker) == 0 {
		return strings.TrimSpace(self.leading[len(simpleStyleMarker):])
	}

	return ""
}

// 解析带有tag的comment, 类似于go结构体中的命名tag
func (self *CommentMeta) ParseTaggedComment() []*TaggedComment {

	var commentArray []*TaggedComment

	// 解析新的样式
	commentArray = parseTaggedComment(self.leading, commentArray)
	commentArray = parseTaggedComment(self.trailing, commentArray)

	// 兼容简单的样式
	if simplecomment := self.findSimpleStyleComment(); simplecomment != "" {

		commentArray = append(commentArray, &TaggedComment{
			Name: "@",
			Data: simplecomment,
		})

	}

	return commentArray
}

func (self *CommentMeta) FindTaggedComment(name string) (string, bool) {

	taggedComment := self.ParseTaggedComment()
	if len(taggedComment) == 0 {
		return "", false
	}

	for _, c := range taggedComment {
		if c.Name == name {
			return c.Data, true
		}
	}

	return "", false

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
