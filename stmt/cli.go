package stmt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

// 工具
type Tools struct {
	Name    string `json:"name"`
	Run     string `json:"run"`
	RunType int    `json:"run_type"`
}

// 分类
type Category struct {
	Name  string  `json:"name"`
	Tools []Tools `json:"tools"`
}

var config_path string
var data []Category
var d = color.New(color.FgCyan, color.Bold)  // 输出重要信息
var s = color.New(color.FgGreen, color.Bold) // 输出成功信息
var e = color.New(color.FgRed)               // 输出错误信息
var current_category int                     // 当前分类id
var show_tools_status bool

func Cli(file_path string, logo_body []byte) {
	show_logo(logo_body)
	config_path = file_path
	read_all(file_path)

	for {
		show_category()
	}
}

// 显示logo
func show_logo(logo_body []byte) {
	fmt.Print(string(logo_body))
}

// 显示头部信息
func show_header(title string) {
	print("\n==========")
	d.Printf(" ⚡ %s ⚡ ", title)
	println("==========")
}

// 读取数据文件
func read_all(config_path string) {
	body, err := ioutil.ReadFile(config_path)
	if err != nil {
		data = []Category{}
	}
	json.Unmarshal(body, &data)
}

// 打印序号
func print_sn(sn string) {
	print("[")
	d.Printf("%s", sn)
	print("] ")
}

// 显示分类列表
func show_category() {
	show_header("分类列表")
	for k, v := range data {
		print_sn(strconv.Itoa(k))
		fmt.Println(v.Name)
	}
	print_sn("a/d/e/q")
	fmt.Println("添加/删除/修改/退出")

	loop := true
	for loop {
		user_input := user_select()

		switch user_input {
		case "a":
			Add_category()
			loop = false
		case "d":
			Del_category()
			loop = false
		case "e":
			Edit_category()
			loop = false
		case "q":
			quit()
			loop = false
		default:
			if sn, err := strconv.Atoi(user_input); err == nil {
				if sn >= 0 && sn < len(data) {
					current_category = sn
					show_tools_status = true
					for show_tools_status {
						show_tools()
					}
					loop = false
				} else {
					e.Println("输入的编号错误！")
				}
			} else {
				e.Println("输入的编号错误！")
			}
		}
	}
}

// 显示工具列表
func show_tools() {
	show_header("工具列表")
	tools_list := data[current_category].Tools
	for k, v := range tools_list {
		print_sn(strconv.Itoa(k))
		fmt.Println(v.Name)
	}
	print_sn("a/b/d/e/q")
	fmt.Println("添加/返回/删除/修改/退出")

	loop := true
	for loop {
		user_input := user_select()
		switch user_input {
		case "a":
			Add_tools()
			loop = false
		case "b":
			show_tools_status = false
			loop = false
		case "d":
			Del_tools()
			loop = false
		case "e":
			Edit_tools()
			loop = false
		case "q":
			quit()
			loop = false
		default:
			if sn, err := strconv.Atoi(user_input); err == nil {
				if sn >= 0 && sn < len(data[current_category].Tools) {
					cmd := tools_list[sn].Run
					run_type := tools_list[sn].RunType
					run_tools(cmd, run_type)
					loop = false
				} else {
					e.Println("输入的编号错误！")
				}
			} else {
				e.Println("输入的编号错误！")
			}
		}
	}
}

// 运行程序
func run_tools(cmd string, run_type int) {
	if run_type == 0 {
		if runtime.GOOS == "windows" {
			Exec_win_daemon(cmd)
		} else {
			Exec_linux_daemon(cmd)
		}
	} else {
		if runtime.GOOS == "windows" {
			Exec_win(cmd)
		} else {
			Exec_linux(cmd)
		}
	}

}

// 读取输入的内容
func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}

// 用户输入
func user_select() string {
	var user_input string
	d.Print("Select﹥")
	Scanf(&user_input)
	if user_input == "" {
		user_input = user_select()
	}
	return user_input
}

// 退出程序
func quit() {
	fmt.Println("👋 Goodbye!")
	os.Exit(0)
}
