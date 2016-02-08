package pbmeta

import (
	pbprotos "github.com/davyxu/pbmeta/proto"
)

type FieldDescriptor struct {
	Define *pbprotos.FieldDescriptorProto
	CommentMeta
	dp *DescriptorPool
}

func (self *FieldDescriptor) Name() string {
	return self.Define.GetName()
}

func (self *FieldDescriptor) TypeName() string {
	return self.Define.GetTypeName()
}

// pb定义中插件返回的格式时 .package.type.value的格式, 这里返回去掉头的.
func (self *FieldDescriptor) FullTypeName() string {
	return normalizeFullName(self.Define.GetTypeName())
}

func (self *FieldDescriptor) MessageDesc() *Descriptor {

	return self.dp.MessageByFullName(self.FullTypeName())
}

func (self *FieldDescriptor) EnumDesc() *EnumDescriptor {
	return self.dp.EnumByFullName(self.FullTypeName())
}

func (self *FieldDescriptor) IsRequired() bool {
	return self.Define.GetLabel() == pbprotos.FieldDescriptorProto_LABEL_REQUIRED
}

func (self *FieldDescriptor) IsOptional() bool {
	return self.Define.GetLabel() == pbprotos.FieldDescriptorProto_LABEL_OPTIONAL
}

func (self *FieldDescriptor) IsRepeated() bool {
	return self.Define.GetLabel() == pbprotos.FieldDescriptorProto_LABEL_REPEATED
}

func (self *FieldDescriptor) Label() pbprotos.FieldDescriptorProto_Label {
	return self.Define.GetLabel()
}
func (self *FieldDescriptor) Type() pbprotos.FieldDescriptorProto_Type {
	return self.Define.GetType()
}

func (self *FieldDescriptor) DefaultValue() string {
	return self.Define.GetDefaultValue()
}

func newFieldDescriptor(raw *pbprotos.FieldDescriptorProto,
	comment *pbprotos.SourceCodeInfo_Location,
	dp *DescriptorPool) *FieldDescriptor {

	md := &FieldDescriptor{
		Define:      raw,
		dp:          dp,
		CommentMeta: newCommentMeta(comment),
	}

	//log.Printf("field: %s %s", md.Define.GetName(), md.comment.GetTrailingComments())

	return md

}
