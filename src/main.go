package src

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

// arg  执行的bat文件完整路径
// 返回错误信息及消息提示
func ExecCommand(arg string) (error, string) {
	c := exec.Command("cmd", "/C", arg)
	w := bytes.NewBuffer(nil)
	c.Stderr = w
	message := "执行" + arg + "文件抽取数据成功"
	_, err1 := os.Stat(arg)
	var err error
	//判断文件是否存在
	if err1 != nil {
		err := c.Run()
		if err != nil {
			fmt.Printf("Run returns: %s\n", err)
		}
		//处理中文乱码
		garbledStr := ConvertByte2String(w.Bytes(), GB18030)
		message = err1.Error() + garbledStr
		//文件不存在并且执行报错
		return err, message
	} else {
		err = c.Run()
		if err != nil {
			//处理中文乱码
			garbledStr := ConvertByte2String(w.Bytes(), GB18030)
			//文件存在 但执行bat文件报错
			return err, garbledStr
		}
	}
	//文件存在并且执行bat文件成功
	return err, message
}
