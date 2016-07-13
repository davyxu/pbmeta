package pbmeta

import (
	"fmt"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type MessageSet struct {
	msgMap   map[string]*Descriptor
	msgArray []*Descriptor
	dp       *DescriptorPool
}

func (self *MessageSet) parse(parentDef interface{}, msgArray []*pbprotos.DescriptorProto, fieldName string, fd *FileDescriptor) {

	fieldNumber, err := getFieldNumber(parentDef, fieldName)
	if err != nil {
		log.Errorln(err)
		return
	}

	self.msgArray = make([]*Descriptor, len(msgArray))

	index := 0
	for _, def := range msgArray {

		path := fmt.Sprintf("%d.%d", fieldNumber, index)

		newMsg := newMessageDescriptor(fd, def, fd.Comment(path), path, self.dp)

		

		self.msgMap[def.GetName()] = newMsg
		self.msgArray[index] = newMsg
		index++
	}
}

func (self *MessageSet) Message(index int) *Descriptor {

	return self.msgArray[index]
}

func (self *MessageSet) MessageCount() int {
	return len(self.msgArray)
}

func (self *MessageSet) MessageByName(name string) *Descriptor {

	if v, ok := self.msgMap[name]; ok {
		return v
	}

	return nil
}

func (self *MessageSet) DebugPrint() {

	for _, v := range self.msgMap {
		log.Debugf("%s", v.Name())
	}
}

func newMessageSet(dp *DescriptorPool) MessageSet {
	return MessageSet{
		dp:     dp,
		msgMap: make(map[string]*Descriptor),
	}
}
