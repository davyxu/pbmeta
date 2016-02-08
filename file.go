package pbmeta

import (
	"strconv"
	"strings"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type FileDescriptor struct {
	Define *pbprotos.FileDescriptorProto

	EnumSet
	MessageSet

	commentMap map[string]*pbprotos.SourceCodeInfo_Location

	dp *DescriptorPool
}

func (self *FileDescriptor) buildComment() {

	if self.Define.SourceCodeInfo == nil {
		return
	}

	for _, loc := range self.Define.SourceCodeInfo.Location {

		if len(loc.GetPath()) == 0 {
			continue
		}

		pathStrList := make([]string, len(loc.GetPath()))

		for i, v := range loc.GetPath() {
			pathStrList[i] = strconv.Itoa(int(v))
		}

		commentPath := strings.Join(pathStrList, ".")

		self.commentMap[commentPath] = loc
	}
}

func (self *FileDescriptor) parse() {

	//log.Printf("%s", self.Define.GetName())

	self.buildComment()

	self.MessageSet.parse(self.Define, self.Define.MessageType, "MessageType", self)

	self.EnumSet.parse(self.Define, self.Define.EnumType, self)
}

func (self *FileDescriptor) FileName() string {
	return self.Define.GetName()
}

func (self *FileDescriptor) PackageName() string {
	return self.Define.GetPackage()
}

func (self *FileDescriptor) Comment(path string) *pbprotos.SourceCodeInfo_Location {
	if v, ok := self.commentMap[path]; ok {
		return v
	}

	return nil
}

func newFileDescriptor(raw *pbprotos.FileDescriptorProto, dp *DescriptorPool) *FileDescriptor {

	fd := &FileDescriptor{
		Define:     raw,
		dp:         dp,
		commentMap: make(map[string]*pbprotos.SourceCodeInfo_Location),
		EnumSet:    newEnumSet(dp),
		MessageSet: newMessageSet(dp),
	}

	fd.parse()

	return fd
}
