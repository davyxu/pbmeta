package main

import (
	"fmt"
	"io/ioutil"
	"os"

	pbcompiler "github.com/davyxu/pbmeta/proto/compiler"
	"github.com/golang/protobuf/proto"
)

func main() {

	pluginInput, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// 转换输入
	var req pbcompiler.CodeGeneratorRequest

	if err := proto.Unmarshal(pluginInput, &req); err != nil {
		fmt.Printf(err.Error())
		return
	}

	// 注意: 由于官方的插件设计输出需要通过CodeGeneratorResponse以文本形式（源码生成）输出
	// 但我们此处需要输出二进制数据，因此使用单独写文件方式输出二进制文件
	// 由于我们无法获得输入文件名，因此通过插件参数来指定
	// 官方的插件参数格式是--插件名_out=参数内容:输出文件夹
	// 这里我们将忽略输出文件夹部分，你可以直接在参数内容中指定输出路径及文件名

	// 创建输出文件，参数来源格式: --meta_out=输出文件名:.   例如--meta_out=test.pb:.
	file, err := os.Create(req.GetParameter())

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	file.Write(pluginInput)

	file.Close()

	// 写输出
	var res pbcompiler.CodeGeneratorResponse

	data, err := proto.Marshal(&res)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	os.Stdout.Write(data)

}
