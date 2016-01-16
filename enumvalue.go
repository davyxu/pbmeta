package pbmeta

import (
	//	"log"
	pbprotos "github.com/davyxu/pbmeta/proto"
)

type EnumValueDescriptor struct {
	Define *pbprotos.EnumValueDescriptorProto
	CommentMeta
}

func (self *EnumValueDescriptor) Name() string {
	return self.Define.GetName()
}

func (self *EnumValueDescriptor) Value() int32 {
	return self.Define.GetNumber()
}

func newEnumValueDescriptor(raw *pbprotos.EnumValueDescriptorProto, comment *pbprotos.SourceCodeInfo_Location) *EnumValueDescriptor {
	md := &EnumValueDescriptor{
		Define:      raw,
		CommentMeta: newCommentMeta(comment),
	}

	//log.Printf("enum value: %s %s", md.Define.GetName(), md.comment.GetTrailingComments())

	return md

}
