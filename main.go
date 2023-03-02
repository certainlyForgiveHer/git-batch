package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	CICC    = "CICC"
	CICCWM  = "CICCWM"
)

// Command 执行命令, 输出到控制台
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

// CommandWithResult 执行命令, 获取输出结果
func CommandWithResult(cmd string) string {
	c := exec.Command("cmd", "/C", cmd) // windows
	//c := exec.Command("bash", "-c", cmd) // mac linux
	stdout, err := c.CombinedOutput()
	if err != nil {
		return ""
	}
	c.Run()
	return string(stdout)
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
		case "sca":
			{
				fmt.Println("暂不支持设置api token至CICC仓库, 请设置为ssh url")
				//setApiToken(CICC, line, v)
				continue
			}
		case "swa":
			{
				setApiToken(CICCWM, line, v)
				continue
			}
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
	fmt.Println("0. 确认仓库账户创建凭证, ssh -> public key, http/https -> api token")
	fmt.Println("1. 确认设置仓库地址")
	fmt.Println("1. 确认当前分支可以push且远程仓库有push权限")
	fmt.Println("2. 确认设置仓库认证(ssh/token方式)")
	fmt.Println("3. 推送代码")
	v["cmd"] = ""
}

func printHelp(v map[string]string) {
	fmt.Println("h  --> help")
	fmt.Println("swa--> set ciccwm api token")
	fmt.Println("sca--> set cicc api token")
	fmt.Println("sb --> show current branch")
	fmt.Println("scr--> set cicc repository")
	fmt.Println("swr--> set ciccwm repository")
	fmt.Println("sr --> show current repository")
	fmt.Println("pp --> show pre execute check list")
	fmt.Println("pu --> push cicc and ciccwm repo, execute git fetch first")
	fmt.Println("q  --> quit")
	v["cmd"] = ""
}

func showRepository(v map[string]string) bool {
	res := Command("git remote -v")
	fmt.Println(v["cmd"] + " complete!")
	v["cmd"] = ""
	return res
}

func setRepository(s string, url string, v map[string]string) bool {
	if url == "scr" || url == "swr" || url == "" {
		fmt.Println("please input git remote url, q to back:")
		return true
	}
	if url == "q" {
		fmt.Println(v["cmd"] + " complete!")
		v["cmd"] = ""
		return true
	}
	Command("git remote remove " + s)
	res := Command("git remote add " + s + " " + url)
	v["cmd"] = ""
	return res
}

func showBranch(v map[string]string) bool {
	res := Command("git branch -vv")
	fmt.Println(v["cmd"] + " complete!")
	v["cmd"] = ""
	return res
}

func setApiToken(s string, token string, v map[string]string) bool {
	if token == "swa" || token == "sca" || token == "" {
		fmt.Println("请确保git remote url已设置为http/https类型")
		fmt.Println("please input api token, q to back:")
		return true
	}
	if token == "q" {
		fmt.Println(v["cmd"] + " complete!")
		v["cmd"] = ""
		return true
	}
	url := CommandWithResult("git remote get-url " + s)
	_, tokenUrl, found := strings.Cut(url, "http://")
	fmt.Println("get tokenUrl:" + tokenUrl)
	if !found {
		_, tokenUrl, found := strings.Cut(url, "https://")
		if !found {
			fmt.Println("set api token failed")
			return false
		}
		fmt.Println(v["cmd"] + " complete!")
		v["cmd"] = ""
		return Command("git remote set-url " + s + " https://oauth2:" + token + "@" + tokenUrl)
	}
	fmt.Println(v["cmd"] + " complete!")
	v["cmd"] = ""
	return Command("git remote set-url " + s + " http://oauth2:" + token + "@" + tokenUrl)
}

func push(v map[string]string) bool {
	Command("git fetch " + CICC)
	Command("git fetch " + CICCWM)
	res := Command("git push "+CICC) && Command("git push "+CICCWM)
	fmt.Println(v["cmd"] + " complete!")
	v["cmd"] = ""
	return res
}
