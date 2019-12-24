package linux_version

import (
	"exercises/os_experiment_result_check"
	"fmt"
	"testing"
)

func TestCheckResult(t *testing.T) {
	err := main.checkResult("/home/joseph/go/src/exercises/os_experiment_result_check/os_experiment_1",
		"/home/joseph/go/src/exercises/os_experiment_result_check/0.txt", "init x x x x p p q q r r x p q r x x x p x")
	if err != nil {
		fmt.Println(err.Error())
	}
}
