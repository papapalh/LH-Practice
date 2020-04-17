package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("管道问题还需要之后在了解")
	fmt.Println("Command 执行程序")
	aa := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`) // 使用给出的参数执行name指定的程序
	fmt.Println("    .Path   执行命令路径")
	fmt.Println("    .Args   执行参数")
	fmt.Println("    .Dir    执行目录")

	fmt.Println("StdoutPipe 返回一个在命令Start后与命令标准输出关联的管道(只是建立管道),Wait方法获知命令结束后会关闭这个管道")
	stdout,_ := aa.StdoutPipe()

	fmt.Println("Start   开始执行命令，但并不会等待该命令完成即返回")
	fmt.Println("Run     开始执行命令，等待该命令完成即返回")
	if err := aa.Start(); err != nil {
		fmt.Println(err)
		return
	}

	//runtime.BlockProfile()

	var person struct {
		Name string
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Wait    会阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的")
	if err := aa.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person)

	fmt.Print("Output  执行命令并返回标准输出的切片,命令完成后才返回 ")
	zz,_ := exec.Command("echo", "-n", "Output").Output()
	fmt.Println(string(zz))
}
