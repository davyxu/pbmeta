# pbmeta
Google Protobuf的Golang的Descriptor系列的封装扩展
带有对SourceCodeInfo支持，可以proto中对结构字段的各种注释做meta扩展

# protoc-gen-meta
通过protoc的plugin机制扩展，将plugin.proto中定义的CodeGeneratorRequest输出成文件

本插件输出的FileDescriptorSet中包含有SourceCodeInfo，可以借由注释来拓展protobuf的meta描述能力


注意: protoc的-o功能也可以输出FileDescriptorSet,但出于性能和数据量考虑，并不包含SourceCodeInfo部分

# 使用方法

* 编译出的protoc-gen-meta

* 编写批处理/shell调用protoc调用protoc-gen-meta导出你的proto文件的二进制描述

* 通过pbmeta库读取

# 功能实现及限制

* 嵌套结构不能被很好的支持, 建议使用全局结构


# 原理参考
我的博客有核心原理描述
http://www.cppblog.com/sunicdavy/archive/2015/03/01/209894.html

# 官方descriptor.proto升级方法
如果官方有descriptor的更新,可以根据下面方法进行更新,重新生成golang代码

go get https://github.com/google/protobuf

go get https://github.com/golang/protobuf

go install https://github.com/golang/protobuf/protoc-gen-go

确保本项目与上面两个项目在同一个GOPATH下

修改compiler/plugin.pb.go的 import google_protobuf "google/protobuf"
为 import google_protobuf "github.com/davyxu/pbmeta/proto"

* Windows 运行方法 proto/GenerateProto.bat
