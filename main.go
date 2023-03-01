package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"os/exec"
	"sync"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func Command(cmd string) bool {
	c := exec.Command("cmd", "/C", cmd) // windows
	//c := exec.Command(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return false
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return false
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	wg.Add(2)
	go read(&wg, stderr)
	go read(&wg, stdout)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	// 等待任务结束
	wg.Wait()
}

func read(wg *sync.WaitGroup, std io.ReadCloser) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			return
		}
		fmt.Print(ConvertByte2String(line, UTF8))

	}
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("*********************************************** power by AIIS ********************************************************")
	fmt.Println("*********************************************** start git batch opt... ***********************************************")
	fmt.Println("input y to continue:")
	for input.Scan() {
		line := input.Text()
		if line == "q" {
			break
		}
		if Command(line) {
			continue
		}
		fmt.Println("execute error, please check pre execute list")
	}
}
