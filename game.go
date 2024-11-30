package main

// "exec"
import (
	"os"
	// "reflect"
	// "strings"
)

// "os/exec"

func startGame(cmd string) error {
	// fmt.Println(a)
	// params, ok := a.([]string)
	// if !ok {
	// 	return fmt.Errorf("INVALID_START_PARAM_TYPE:TYPE=%s:VALUE=%s", reflect.TypeOf(a), params)
	// }
	// var gtaSaExe string = fmt.Sprintf("%s\\gta_sa.exe", cmd)
	// _, err := os.Stat(gtaSaExe)
	// if err != nil {
	// 	return err
	// }
	_, err := os.StartProcess(cmd, []string{}, &os.ProcAttr{})
	return err
}
