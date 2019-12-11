package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

/* 若要开始检查，直接在终端内启动程序，将要检测的程序名作为第一个参数传入（不需要./）
   如果运行的是Python文件，请添加第二个参数（python2 or python3）*/

var (
	redBg   = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	greenBg = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	reset     = string([]byte{27, 91, 48, 109})
	dirFiles  []string
	pyVersion = 0
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, file := range files {
		dirFiles = append(dirFiles, file.Name())
	}
	osInfo := runtime.GOOS
	fmt.Printf("Environment: %s\n", osInfo)

	if len(os.Args) <= 1 {
		fmt.Println("please pass your file name as the first parameter")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		err = resultCheck("./" + os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if len(os.Args) == 3 {
		pyArg := os.Args[2]
		pyArg = strings.TrimSpace(pyArg)
		if pyArg == "python3" {
			pyVersion = 3
		} else if pyArg == "python2" {
			pyVersion = 2
		} else {
			fmt.Println("cannot identify your python version")
			os.Exit(1)
		}
		err = resultCheck(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	os.Exit(0)
}

func resultCheck(fileName string) error {
	correctResults, err := checkFiles()
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
		err = checkResult(fileName, testFile, correctResults[i])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	return nil

}

func checkIfFileExists(fileName string) bool {
	for _, eachName := range dirFiles {
		if eachName == fileName[2:] {
			return true
		}
	}
	return false
}

func checkFiles() ([]string, error) {
	/* 返回测试文件序列及结果文件内容 */
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
				err = nil
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

func checkResult(fileName, testFileName string, supposedSeq string) error {
	var cmd *exec.Cmd

	if pyVersion == 0 {
		cmd = exec.Command(fileName, testFileName)
	} else if pyVersion == 3 {
		cmd = exec.Command("python3", fileName, testFileName)
	} else {
		cmd = exec.Command("python2", fileName, testFileName)
	}

	resultBuf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	//resultBuf := string("@joseph-laptop:~/go/src/exercises/os_experiment_result_check$ go build joseph@joseph-laptop:~/go/src/exercises/os_experiment_result_check$ ./os_experiment_result_check os_experiment_1 Environment: linux./0.txt   passed")
	result := string(resultBuf)
	result = strings.Replace(result, "\n", "", -1)
	result = strings.Replace(result, "\r", "", -1)
	result = strings.TrimSpace(result)
	resultBufList := strings.Split(result, " ")
	supposedSeqList := strings.Split(supposedSeq, " ")

	// 长度大于期望串长度的输出会被忽略
	var incorrect = false
	var outputResult, outputSupposed []string
	for index, eachResult := range supposedSeqList {
		if index < len(resultBufList) && eachResult == resultBufList[index] {
			outputResult = append(outputResult, eachResult)
			outputSupposed = append(outputSupposed, eachResult)
		} else if index < len(resultBufList) && eachResult != resultBufList[index] {
			outputResult = append(outputResult, redBg+resultBufList[index]+reset+"   ")
			outputSupposed = append(outputSupposed, greenBg+eachResult+reset+"   ")
			incorrect = true
		} else {
			outputSupposed = append(outputSupposed, supposedSeqList[index:]...)
			incorrect = true
			break
		}
	}
	if outputResult == nil || outputSupposed == nil {
		return errors.New("output strings are none")
	}

	if incorrect {
		fmt.Printf("%s%s%s%s\n", redBg, testFileName, "   incorrect", reset)
	} else {
		fmt.Printf("%s%s%s%s\n", greenBg, testFileName, "   passed", reset)
	}
	fmt.Printf("%-5d", 0)
	for i := 1; i < len(supposedSeqList); i++ {
		fmt.Printf("%-4d", i)
	}
	fmt.Println()

	fmt.Printf("%-5s", outputResult[0])
	for i := 1; i < len(outputResult); i++ {
		fmt.Printf("%-4s", outputResult[i])
	}
	fmt.Println()

	fmt.Printf("%-5s", outputSupposed[0])
	for i := 1; i < len(outputSupposed); i++ {
		fmt.Printf("%-4s", outputSupposed[i])
	}
	fmt.Println()
	return nil
}
