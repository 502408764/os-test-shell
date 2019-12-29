package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/issue9/term/colors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type checkBox struct {
	outputStr string
	supposedStr string
	correct bool
}

var (
	dirFileNames []string
	thirdArg     = false
)

func main() {
	// 获取当前目录下的文件，并将文件名都存入dirFileNames
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln("cannot read files from current directory")
	}
	for _, eachFcb := range files {
		dirFileNames = append(dirFileNames, eachFcb.Name())
	}

	osInfo := runtime.GOOS
	fmt.Printf("Environment: %s\n\n", osInfo)

	if len(os.Args) <= 1 {
		fmt.Println("please pass your file name as the first parameter")
		os.Exit(1)
	} else if len(os.Args) == 3 {
		thirdArg = true
		err = repeatCheck()
	} else {
		err = repeatCheck()
	}

	if err != nil {
		log.Fatalln(err.Error())
	}
	os.Exit(0)
}

// checkIfFileExists 检查txt文件是否在当前工作目录
func checkIfFileExists(fileName string) bool {
	for _, eachName := range dirFileNames {
		if eachName == fileName[2:] {
			return true
		}
	}
	return false
}

// readCorrectResults 函数返回测试序列结果文件内容
func readCorrectResults() ([]string, error) {
	resultFile, err := os.Open("./result.txt")
	if err != nil {
		return nil, err
	}
	defer resultFile.Close()

	bufReader := bufio.NewReader(resultFile)
	var lines []string
	for {
		line, _, err := bufReader.ReadLine() // 按行读
		if err != nil {
			if err == io.EOF {
				// 到文件尾就跳出
				break
			}
		} else {
			l := string(line)
			lines = append(lines, l)
		}
	}
	if len(lines) == 0 {
		return nil, errors.New("empty result file")
	}
	return lines, nil
}

// checkResult 执行一次学生的实验程序，将输出结果和
// 从参数接收到的字符串序列做比较，高亮显示不同的部分
func checkResult(testFileName, supposedSeq string) error {
	var cmd *exec.Cmd

	// 如果有第三个参数，就将第三个参数作为命令
	// 常见的参数有java、python2、python3
	if !thirdArg {
		cmd = exec.Command(os.Args[1], testFileName)
	} else {
		cmd = exec.Command(os.Args[2], os.Args[1], testFileName)
	}

	resultBuf, err := cmd.Output()
	if err != nil {
		return err
	}

	// 去掉换行符和不必要的前后空格
	result := string(resultBuf)
	result = strings.Replace(result, "\r", "", -1)
	result = strings.Replace(result, "\n", "", -1)
	result = strings.TrimSpace(result)

	// 把输出字符串和结果字符串劈成List逐项比较
	resultList := strings.Split(result, " ")
	supposedList := strings.Split(supposedSeq, " ")

	var boxes []checkBox
	var incorrect = false
	for index, supposed := range supposedList {
		// 长度大于期望串长度的输出会被忽略
		if index < len(resultList) && supposed == resultList[index] {
			boxes = append(boxes, checkBox{
				outputStr:   supposed,
				supposedStr: supposed,
				correct:     true,
			})
		} else if index < len(resultList) && supposed != resultList[index] {
			boxes = append(boxes, checkBox{
				outputStr:   resultList[index],
				supposedStr: supposed,
				correct:     false,
			})
			incorrect = true
		} else {
			boxes = append(boxes, checkBox{
				outputStr:   "",
				supposedStr: supposed,
				correct:     false,
			})
			incorrect = true
		}
	}

	// 打印本次测试的结果
	if incorrect {
		_, _ = colors.Printf(colors.White, colors.Red, "%s: incorrect", testFileName)
	} else {
		_, _ = colors.Printf(colors.White, colors.Green, "%s passed", testFileName)
	}
	fmt.Println()

	fmt.Printf("%-8s: ", "Index")
	for i := 0; i < len(supposedList); i++ {
		fmt.Printf("%-4d", i)
	}
	fmt.Println()

	fmt.Printf("%-8s: ", "Output")
	for _, eachBox := range boxes {
		if eachBox.correct {
			fmt.Printf("%-4s", eachBox.outputStr)
		} else {
			_, _ = colors.Printf(colors.Default, colors.Red, "%-4s", eachBox.outputStr)
		}
	}
	fmt.Println()

	fmt.Printf("%-8s: ", "Supposed")
	for _, eachBox := range boxes {
		if eachBox.correct {
			fmt.Printf("%-4s", eachBox.supposedStr)
		} else {
			_, _ = colors.Printf(colors.Default, colors.Green, "%-4s", eachBox.supposedStr)
		}
	}
	fmt.Println()
	return nil
}

// repeatCheck 根据结果文件的行数运行结果检查
func repeatCheck() error {
	correctResults, err := readCorrectResults()
	if err != nil {
		return err
	}
	fileNum := len(correctResults)

	for i := 0; i < fileNum; i++ {
		testFile := "./" + string(48+i) + ".txt"
		ifExist := checkIfFileExists(testFile)
		if !ifExist {
			fmt.Printf("cannot find file %s\n", testFile)
			continue
		}
		err = checkResult(testFile, correctResults[i])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	return nil
}
