# pbmeta
通过读取FileDescriptorSet数据，并对核心protobuf结构进行结构封装
类似于C++中的此类封装
同时带有对SourceCodeInfo支持，可以proto中对结构字段的各种注释做meta扩展

# protoc-gen-meta
通过protoc的plugin机制扩展，将plugin.proto中定义的CodeGeneratorRequest输出成文件

本插件输出的FileDescriptorSet中包含有SourceCodeInfo，可以借由注释来拓展protobuf的meta描述能力


注意: protoc的-o功能也可以输出FileDescriptorSet,但出于性能和数据量考虑，并不包含SourceCodeInfo部分

# 原理参考
我的博客有一些原理描述
http://www.cppblog.com/sunicdavy/archive/2015/03/01/209894.html

# 官方descriptor.proto升级方法

go get https://github.com/google/protobuf
go get https://github.com/golang/protobuf

确保本项目与上面两个项目在同一个GOPATH下

修改compiler/plugin.pb.go的 import google_protobuf "google/protobuf"
为 import google_protobuf "github.com/davyxu/pbmeta/proto"

# Windows
运行proto/GenerateProto.bat
