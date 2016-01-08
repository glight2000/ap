package action

import (
	"encoding/json"
	"os"
	"github.com/Unknwon/log"
	"bufio"
	"io"
)
func DecodeCfg() *Config{
	filePath := os.Getenv("GOBIN") + "ap.cfg"

	data := make([]byte,0)
	err := ReadBlock(filePath, 1024, func(pieces []byte) {
		data = append(data, pieces...)
	})

	if(err != nil){
		log.Fatal("Read config file failed:", filePath, "\n", err)
	}

	configPtr := &Config{}

	err = json.Unmarshal(data, configPtr)

	if(err!= nil){
		log.Fatal("Decode config file failed:", filePath, "\n", err)
		return nil
	}

	return configPtr
}

func ReadBlock(filePth string, bufSize int, handle func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, bufSize) //一次读取多少个字节
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		handle(buf[:n]) // n 是成功读取字节数
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

type Config struct {
	To []ToItem
	Customize []CustomizeItem
}

type ToItem struct {
	Alias []string
	Target string
}

type CustomizeItem struct {
	Shortcut string
	Targets []CustomizeTarget
}

type CustomizeTarget struct {
	App string
	Argument_filter string
	Ext_arguments []string
	IsArgumentsInherit bool
}