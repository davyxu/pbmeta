package pbmeta

import (
	"fmt"
	"strings"

	pbprotos "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type FieldDescriptor struct {
	Define *pbprotos.FieldDescriptorProto
	CommentMeta
	dp *DescriptorPool

	msg *Descriptor
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

	// TODO 内嵌结构估计得使用和枚举一样的方法进行
	if d := self.msg.NestedMsg.MessageByName(self.FullTypeName()); d != nil {
		return d
	}

	return self.dp.MessageByFullName(self.FullTypeName())
}

func (self *FieldDescriptor) EnumDesc() *EnumDescriptor {

	// 这里很奇怪, 按descriptor.proto里FieldDescriptorProto的type_name注释说:
	// "."开头的名字是全局空间, 然而内嵌的类型并不是这样导出的
	// 而要找到结构体的完全路径是非常困难的, 因此这里用一个取巧的方法:
	// 使用字段所在消息的名字+字段所在消息内嵌枚举的名字不断进行尝试, 直到
	// 符合自己枚举的全名尾缀
	for i := 0; i < self.msg.EnumSet.EnumCount(); i++ {

		en := self.msg.EnumSet.Enum(i)

		tailName := fmt.Sprintf("%s.%s", self.msg.Define.GetName(), en.Name())

		if strings.HasSuffix(self.FullTypeName(), tailName) {
			return en
		}

	}

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

func (self *FieldDescriptor) IsMessageType() bool {
	return self.Define.GetType() == pbprotos.FieldDescriptorProto_TYPE_MESSAGE
}

func (self *FieldDescriptor) IsEnumType() bool {
	return self.Define.GetType() == pbprotos.FieldDescriptorProto_TYPE_ENUM
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

func newFieldDescriptor(parent *Descriptor, raw *pbprotos.FieldDescriptorProto,
	comment *pbprotos.SourceCodeInfo_Location,
	dp *DescriptorPool) *FieldDescriptor {

	md := &FieldDescriptor{
		Define:      raw,
		dp:          dp,
		msg:         parent,
		CommentMeta: newCommentMeta(comment),
	}

	//log.Printf("field: %s %s", md.Define.GetName(), md.comment.GetTrailingComments())

	return md

}
