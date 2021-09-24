package stmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// 选择分类
func select_category() int {
	category_id := user_input("分类编号")
	if sn, err := strconv.Atoi(category_id); err == nil {
		if sn >= 0 && sn < len(data) {
			return sn
		} else {
			e.Println("输入的编号错误！")
		}
	} else {
		e.Println("输入的编号错误！")
	}
	return -1
}

// 添加分类
func Add_category() {
	name := user_input("新分类名称")
	new_category := Category{name, nil}
	data = append(data, new_category)
	Save_to_file()
	s.Println("添加成功！")
}

// 删除分类
func Del_category() {
	sn := select_category()
	if sn != -1 {
		// data = append(data[:current_category], data[current_category+1:]...)
		data = append(data[:sn], data[sn+1:]...)
		Save_to_file()
		s.Println("删除成功！")
	} else {
		e.Println("删除失败")
	}
}

// 修改分类
func Edit_category() {
	sn := select_category()
	new_name := user_input("分类新名称")
	if sn != -1 {
		data[sn].Name = new_name
		Save_to_file()
		s.Println("修改成功！")
	} else {
		e.Println("修改失败")
	}
}

// 选择工具
func select_tools() int {
	tools_id := user_input("工具编号")
	if sn, err := strconv.Atoi(tools_id); err == nil {
		if sn >= 0 && sn < len(data[current_category].Tools) {
			return sn
		} else {
			e.Println("输入的编号错误！")
		}
	} else {
		e.Println("输入的编号错误！")
	}
	return -1
}

// 选择运行模式
func select_run_type() int {
	var run_type int

	print_sn("0")
	fmt.Println("后台")
	print_sn("1")
	fmt.Println("终端")
	run_type_id := user_input("运行模式")
	switch run_type_id {
	case "0":
		run_type = 0
	case "1":
		run_type = 1
	default:
		e.Println("输入的编号错误！")
		run_type = select_run_type()
	}
	return run_type
}

// 添加工具
func Add_tools() {
	name := user_input("工具名称")
	run_cmd := user_input("启动命令")
	run_type := select_run_type()
	new_tools := Tools{name, run_cmd, run_type}
	data[current_category].Tools = append(data[current_category].Tools, new_tools)
	Save_to_file()
	s.Println("添加成功！")
}

// 删除工具
func Del_tools() {
	sn := select_tools()
	if sn != -1 {
		tools_list := data[current_category].Tools
		tools_list = append(tools_list[:sn], tools_list[sn+1:]...)
		data[current_category].Tools = tools_list
		Save_to_file()
		s.Println("删除成功！")
	} else {
		e.Println("删除失败")
	}
}

// 修改工具
func Edit_tools() {
	sn := select_tools()
	new_name := user_input("工具新名称")
	new_run_cmd := user_input("工具新启动命令")
	new_run_type := select_run_type()
	if sn != -1 {
		data[current_category].Tools[sn].Name = new_name
		data[current_category].Tools[sn].Run = new_run_cmd
		data[current_category].Tools[sn].RunType = new_run_type
		Save_to_file()
		s.Println("修改成功！")
	} else {
		e.Println("修改失败")
	}
}

// 用户输入
func user_input(prompt string) string {
	var input_content string
	fmt.Printf("%s: ", prompt)
	Scanf(&input_content)
	if input_content == "" {
		input_content = user_input(prompt)
	}
	return input_content
}

// 保存到文件
func Save_to_file() bool {
	json_data, _ := json.Marshal(data)
	err := ioutil.WriteFile(config_path, json_data, 0644)
	return err == nil
}
