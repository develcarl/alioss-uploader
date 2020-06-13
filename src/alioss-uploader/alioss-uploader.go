package main

import (
	"fmt"
	"os"

	"github.com/develcarl/alioss-uploader/src/operate"
)

func main() {
	var args []string = os.Args
	if len(args) == 1 {
		panic("参数不可为空")
	}
	for i := 1; i < len(args); i++ {
		remotePath, err := operate.PutObject(args[i])
		if err != nil {
			operate.HandleError(err)
		}
		fmt.Println("http://typora-pic-plat.oss-cn-beijing.aliyuncs.com/" + remotePath)
	}

}
