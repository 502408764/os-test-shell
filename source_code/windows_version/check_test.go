package main

import (
	"fmt"
	"testing"
)

func TestCheckResult(t *testing.T) {
	err := checkResult("D:\\GoProjects\\src\\exercises\\os_experiment_result_check\\os_experiment_1.exe",
		"D:\\GoProjects\\src\\exercises\\os_experiment_result_check\\0.txt", "init x x x x p p q q r r x p q r x x x p x")
	if err != nil {
		fmt.Println(err.Error())
	}
}
