package pbmeta

import (
	"fmt"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type Descriptor struct {
	Define *pbprotos.DescriptorProto
	EnumSet
	CommentMeta

	fieldMap     map[string]*FieldDescriptor
	fieldArray   []*FieldDescriptor
	fieldNumMap  map[int32]*FieldDescriptor
	fieldDescMap map[*FieldDescriptor]*FieldDescriptor

	parentPath string

	dp *DescriptorPool
}

func (self *Descriptor) parse(fd *FileDescriptor) {

	// 取描述定义中的序号
	fieldNumber, err := getFieldNumber(self.Define, "Field")
	if err != nil {
		log.Errorln(err)
		return
	}

	self.fieldArray = make([]*FieldDescriptor, len(self.Define.Field))

	index := 0

	for _, v := range self.Define.Field {

		path := fmt.Sprintf("%s.%d.%d", self.parentPath, fieldNumber, index)

		newField := newFieldDescriptor(v, fd.Comment(path), self.dp)

		// 添加新的字段描述符
		self.fieldMap[v.GetName()] = newField
		self.fieldArray[index] = newField
		self.fieldNumMap[v.GetNumber()] = newField
		self.fieldDescMap[newField] = newField
		index++
	}

	self.EnumSet.parse(self.Define, self.Define.EnumType, fd)

}

func (self *Descriptor) Name() string {
	return self.Define.GetName()
}

func (self *Descriptor) FieldByName(name string) *FieldDescriptor {
	if v, ok := self.fieldMap[name]; ok {
		return v
	}

	return nil
}

func (self *Descriptor) FieldByNumber(number int32) *FieldDescriptor {
	if v, ok := self.fieldNumMap[number]; ok {
		return v
	}

	return nil
}

func (self *Descriptor) Contains(fd *FieldDescriptor) bool {
	if v, ok := self.fieldDescMap[fd]; ok && v == fd {
		return true
	}

	return false
}

func (self *Descriptor) Field(index int) *FieldDescriptor {
	return self.fieldArray[index]
}

func (self *Descriptor) FieldCount() int {
	return len(self.fieldArray)
}

func newMessageDescriptor(fd *FileDescriptor,
	raw *pbprotos.DescriptorProto,
	comment *pbprotos.SourceCodeInfo_Location,
	parentPath string,
	dp *DescriptorPool) *Descriptor {

	md := &Descriptor{
		Define:       raw,
		dp:           dp,
		fieldMap:     make(map[string]*FieldDescriptor),
		fieldNumMap:  make(map[int32]*FieldDescriptor),
		fieldDescMap: make(map[*FieldDescriptor]*FieldDescriptor),
		EnumSet:      newEnumSet(dp),
		CommentMeta:  newCommentMeta(comment),
		parentPath:   parentPath,
	}

	//log.Printf("msg: %s %s", md.Define.GetName(), md.comment.GetLeadingComments())

	md.parse(fd)

	return md

}
