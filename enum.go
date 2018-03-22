package pbmeta

import (
	"bytes"
	"fmt"

	pbprotos "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type EnumDescriptor struct {
	Define *pbprotos.EnumDescriptorProto
	CommentMeta

	valueMap    map[string]*EnumValueDescriptor
	valueArray  []*EnumValueDescriptor
	valueNumMap map[int32]*EnumValueDescriptor

	parentPath string
}

func (self *EnumDescriptor) Name() string {
	return self.Define.GetName()
}

func (self *EnumDescriptor) parse(fd *FileDescriptor) {
	// 取描述定义中的序号
	fieldNumber, err := getFieldNumber(self.Define, "Value")
	if err != nil {
		log.Errorln(err)
		return
	}

	self.valueArray = make([]*EnumValueDescriptor, len(self.Define.Value))

	index := 0

	for _, def := range self.Define.Value {

		path := fmt.Sprintf("%s.%d.%d", self.parentPath, fieldNumber, index)

		newValue := newEnumValueDescriptor(def, fd.Comment(path))

		// 添加新的字段描述符
		self.valueMap[def.GetName()] = newValue
		self.valueArray[index] = newValue
		self.valueNumMap[def.GetNumber()] = newValue
		index++
	}
}

func (self *EnumDescriptor) ValueByName(name string) *EnumValueDescriptor {
	if v, ok := self.valueMap[name]; ok {
		return v
	}

	return nil
}

func (self *EnumDescriptor) ValueByNumber(number int32) *EnumValueDescriptor {
	if v, ok := self.valueNumMap[number]; ok {
		return v
	}

	return nil
}

func (self *EnumDescriptor) Value(index int) *EnumValueDescriptor {
	return self.valueArray[index]
}

func (self *EnumDescriptor) ValueCount() int {
	return len(self.valueArray)
}

func (self *EnumDescriptor) String() string {

	var buffer bytes.Buffer

	for _, def := range self.Define.Value {
		buffer.WriteString(fmt.Sprintf("%s=%d ", def.GetName(), def.GetNumber()))
	}

	return buffer.String()
}

func newEnumDescriptor(fd *FileDescriptor,
	raw *pbprotos.EnumDescriptorProto,
	comment *pbprotos.SourceCodeInfo_Location,
	parentPath string) *EnumDescriptor {

	md := &EnumDescriptor{
		Define:      raw,
		valueMap:    make(map[string]*EnumValueDescriptor),
		valueNumMap: make(map[int32]*EnumValueDescriptor),
		CommentMeta: newCommentMeta(comment),
		parentPath:  parentPath,
	}

	//log.Printf("enum: %s %s", md.Define.GetName(), md.comment.GetLeadingComments())

	md.parse(fd)

	return md

}
