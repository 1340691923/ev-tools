package main

import (
	"flag"
	"fmt"
	"github.com/1340691923/eve-plugin-sdk-go/build"
	"os"
	"path/filepath"
)

func main() {
	// 解析命令行参数
	sourceDir := flag.String("source", "", "要压缩的文件夹路径")
	destZip := flag.String("output", "dist/source_code.zip", "输出的zip文件路径")
	excludeDir := flag.String("exclude", "node_modules,dist,.idea,.vscode", "要排除的文件夹路径")
	flag.Parse()

	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("获取执行路径时出错:", err)
		return
	}

	// 获取执行文件所在的目录
	execDir := filepath.Dir(execPath)
	execFileName := filepath.Base(execPath)
	// 检查输入是否有效
	if *sourceDir == "" {
		sourceDir = &execDir
	}

	os.MkdirAll("dist", 777)

	err = build.CompressPathToZip(*sourceDir, *excludeDir, execFileName, *destZip)

	if err != nil {
		fmt.Println("压缩过程中出错:", err)
		os.Exit(1)
	}

	fmt.Println("压缩完成，输出文件:", *destZip)
}
