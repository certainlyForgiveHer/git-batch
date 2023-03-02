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
	CICC    = "CICC"
	CICCWM  = "CICCWM"
)

func Command(cmd string) bool {
	c := exec.Command("cmd", "/C", cmd) // windows
	//c := exec.Command("bash", "-c", cmd) // mac linux
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
	c.Wait()
	return c.ProcessState.Success()
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
	v := make(map[string]string)
	v["cmd"] = "pu"
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("*********************************************** power by AIIS ********************************************************")
	fmt.Println("       ┌─┐       ┌─┐ + +")
	fmt.Println("    ┌──┘ ┴───────┘ ┴──┐++")
	fmt.Println("    │                 │")
	fmt.Println("    │       ───       │++ + + +")
	fmt.Println("    ███████───███████ │+")
	fmt.Println("    │                 │+")
	fmt.Println("    │       ─┴─       │")
	fmt.Println("    │                 │")
	fmt.Println("    └───┐         ┌───┘")
	fmt.Println("        │         │")
	fmt.Println("        │         │   + +")
	fmt.Println("        │         │")
	fmt.Println("        │         └──────────────┐")
	fmt.Println("        │                        │")
	fmt.Println("        │                        ├─┐")
	fmt.Println("        │                        ┌─┘")
	fmt.Println("        │                        │")
	fmt.Println("        └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +")
	fmt.Println("          │ ─┤ ─┤       │ ─┤ ─┤")
	fmt.Println("          └──┴──┘       └──┴──┘  + + + +")
	fmt.Println("                 神兽保佑")
	fmt.Println("                代码无BUG!")
	fmt.Println("*********************************************************************************************************************")
	printHelp(v)
	for input.Scan() {
		line := input.Text()
		if v["cmd"] == "" {
			v["cmd"] = line
		}
		switch v["cmd"] {
		case "sr":
			{
				showRepository(v)
				continue
			}
		case "scr":
			{
				setRepository(CICC, line, v)
				continue
			}
		case "swr":
			{
				setRepository(CICCWM, line, v)
				continue
			}
		case "pu":
			{
				if push(v) {
					continue
				}
			}
		case "sa":
			{
			}
		case "sb":
			{
				if showBranch(v) {
					continue
				}
			}
		case "pp":
			{
				printPreExecuteList(v)
				continue
			}
		case "h":
			{
				printHelp(v)
				continue
			}
		case "q":
			{
				println("bye")
				println("for the lich king!!!!!")
				return
			}
		default:
			break
		}
		fmt.Println("execute error, please check stdout and pre execute list")
		printPreExecuteList(v)
	}
}

func printPreExecuteList(v map[string]string) {
	fmt.Println("0. 确认创建gitlab api token")
	fmt.Println("1. 确认设置cicc/ciccwm仓库地址")
	fmt.Println("1. 确认当前分支可以push且远程仓库有push权限")
	fmt.Println("2. 确认设置仓库认证（cicc支持ssh/apitoken方式，ciccwm支持apitoken方式）")
	fmt.Println("3. 推送代码")
	v["cmd"] = ""
}

func printHelp(v map[string]string) {
	fmt.Println("h  --> help")
	fmt.Println("sa --> set ciccwm api token")
	fmt.Println("sb --> show current branch")
	fmt.Println("scr--> set cicc repository")
	fmt.Println("swr--> set ciccwm repository")
	fmt.Println("sr --> show current repository")
	fmt.Println("pp --> show pre execute check list")
	fmt.Println("pu --> push cicc/ciccwm, execute git fetch --all first")
	fmt.Println("q  --> quit")
	v["cmd"] = ""
}

func showRepository(v map[string]string) bool {
	res := Command("git remote -v")
	v["cmd"] = ""
	return res
}

func setRepository(s string, url string, v map[string]string) bool {
	if url == "scr" || url == "swr" || url == "" {
		fmt.Println("please input git remote url, q to back:")
		return true
	}
	if url == "q" {
		return true
	}
	Command("git remote remove " + s)
	res := Command("git remote add " + s + " " + url)
	v["cmd"] = ""
	return res
}

func showBranch(v map[string]string) bool {
	res := Command("git branch -vv")
	v["cmd"] = ""
	return res
}

func setApiToken(token string, v map[string]string) bool {
	return false
}

func push(v map[string]string) bool {
	Command("git fetch --all")
	res := Command("git push cicc") && Command("git push ciccwm")
	v["cmd"] = ""
	return res
}
