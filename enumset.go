package pbmeta

import (
	"fmt"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type EnumSet struct {
	enumMap   map[string]*EnumDescriptor
	enumArray []*EnumDescriptor
	dp        *DescriptorPool
}

func (self *EnumSet) parse(parentDef interface{}, enumArray []*pbprotos.EnumDescriptorProto, fd *FileDescriptor) {

	fieldNumber, err := getFieldNumber(parentDef, "EnumType")
	if err != nil {
		log.Errorln(err)
		return
	}

	self.enumArray = make([]*EnumDescriptor, len(enumArray))

	index := 0
	for _, v := range enumArray {

		path := fmt.Sprintf("%d.%d", fieldNumber, index)

		loc := fd.Comment(path)

		newEnum := newEnumDescriptor(fd, v, loc, path)

		// 注册到全局
		self.dp.registerEnum(fd, newEnum)

		self.enumMap[v.GetName()] = newEnum
		self.enumArray[index] = newEnum
		index++
	}
}

func (self *EnumSet) Enum(index int) *EnumDescriptor {

	return self.enumArray[index]
}

func (self *EnumSet) EnumCount() int {
	return len(self.enumArray)
}

func (self *EnumSet) EnumByName(name string) *EnumDescriptor {

	if v, ok := self.enumMap[name]; ok {
		return v
	}

	return nil
}

func (self *EnumSet) EnumValueByName(name string) *EnumValueDescriptor {
	for _, v := range self.enumMap {

		if vv := v.ValueByName(name); vv != nil {
			return vv
		}

	}

	return nil

}

func newEnumSet(dp *DescriptorPool) EnumSet {
	return EnumSet{
		dp:      dp,
		enumMap: make(map[string]*EnumDescriptor),
	}
}
