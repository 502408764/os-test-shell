package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/issue9/term/colors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	//greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	//redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	//reset     = string([]byte{27, 91, 48, 109})
	dirFiles []string
	thirdArg = false
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
	fmt.Printf("Environment: %s\n", strings.ToTitle(osInfo))

	if len(os.Args) <= 1 {
		fmt.Println("please pass your file name as the first parameter")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		err = resultCheck("./" + os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if len(os.Args) == 3 {
		//execArg := os.Args[2]
		//execArg = strings.TrimSpace(execArg)
		//if execArg == "python" {
		//
		//}
		//pyArg := os.Args[2]
		//pyArg = strings.TrimSpace(pyArg)
		//if pyArg == "python" {
		//	thirdArg = 3
		//} else {
		//	fmt.Println("cannot identify your python version")
		//	os.Exit(1)
		//}
		thirdArg = true
		if os.Args[2] != "java" && os.Args[2] != "python" {
			fmt.Println("cannot identify your args.")
		}
		err = resultCheck(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("too many or too few arguments.")
		os.Exit(1)
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

	if thirdArg == false {
		cmd = exec.Command(fileName, testFileName)
	} else {
		cmd = exec.Command(os.Args[2], fileName, testFileName)
	}

	resultBuf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	result := string(resultBuf)
	result = strings.Replace(result, "\n", "", -1)
	result = strings.Replace(result, "\r", "", -1)
	result = strings.TrimSpace(result)
	resultBufList := strings.Split(result, " ")
	supposedSeqList := strings.Split(supposedSeq, " ")

	// 长度大于期望串长度的输出会被忽略
	var incorrect = false
	var outputResult, outputSupposed []string
	var eachCorrect []bool
	for index, eachResult := range supposedSeqList {
		if index < len(resultBufList) && eachResult == resultBufList[index] {
			outputResult = append(outputResult, eachResult)
			outputSupposed = append(outputSupposed, eachResult)
			eachCorrect = append(eachCorrect, true)
		} else if index < len(resultBufList) && eachResult != resultBufList[index] {
			//outputResult = append(outputResult, redBg+resultBufList[index]+reset+"   ")
			//outputSupposed = append(outputSupposed, greenBg+eachResult+reset+"   ")
			outputResult = append(outputResult, resultBufList[index])
			outputSupposed = append(outputSupposed, eachResult)
			eachCorrect = append(eachCorrect, false)
			incorrect = true
		} else {
			outputSupposed = append(outputSupposed, supposedSeqList[index])
			eachCorrect = append(eachCorrect, false)
			incorrect = true
		}
	}
	if outputResult == nil || outputSupposed == nil {
		return errors.New("output strings are none")
	}

	if incorrect {
		_, _ = colors.Printf(colors.White, colors.Red, "%s: incorrect\n", testFileName)
		//fmt.Printf("%s%s%s%s\n", redBg, testFileName, "   incorrect", reset)
	} else {
		_, _ = colors.Printf(colors.White, colors.Green, "%s passed\n", testFileName)
		//fmt.Printf("%s%s%s%s\n", greenBg, testFileName, "   passed", reset)
	}
	//fmt.Printf("%-5d", 0)
	for i := 0; i < len(supposedSeqList); i++ {
		fmt.Printf("%-4d", i)
	}
	fmt.Println()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	//fmt.Printf("%-5s", outputResult[0])
	for i := 0; i < len(outputResult); i++ {
		if i < len(eachCorrect) && eachCorrect[i] == true {
			fmt.Printf("%-4s", outputResult[i])
		} else {
			_, _ = colors.Printf(colors.Default, colors.Red, "%-4s", outputResult[i])
		}
	}
	fmt.Println()

	//fmt.Printf("%-5s", outputSupposed[0])
	for i := 0; i < len(outputSupposed); i++ {
		if i < len(eachCorrect) && eachCorrect[i] == true {
			fmt.Printf("%-4s", outputSupposed[i])
		} else {
			_, _ = colors.Printf(colors.Default, colors.Green, "%-4s", outputSupposed[i])
		}
	}
	fmt.Println()
	return nil
}
