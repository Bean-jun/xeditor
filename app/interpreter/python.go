package interpreter

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"strings"
	"xeditor/app/utils"
)

var (
	//go:embed ..\..\bin\uv.exe
	uv                       []byte
	defaultPythonInterpreter PythonInterpreter
)

func init() {
	tmpExe := filepath.Join(os.TempDir(), "uv.exe")
	log.Println(tmpExe)
	err := os.WriteFile(tmpExe, uv, 0755)
	if err != nil {
		log.Println("unzip python interpreter error: ", err.Error())
		return
	}
	defaultPythonInterpreter.UVExecutable = tmpExe
}

type PythonInterpreter struct {
	UVExecutable string
	PyExecutable string
}

func NewPythonInterpreter() *PythonInterpreter {
	return &defaultPythonInterpreter
}

func (p *PythonInterpreter) ExecuteCmd(args ...string) (string, string, error) {
	return utils.ExecuteCmdNoWindow(p.UVExecutable, args...)
}

func (p *PythonInterpreter) GetAllInterpreterList() ([]string, error) {
	o1, _, err := p.ExecuteCmd("python", "list")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(o1, "\n")
	availableList := make([]string, 0)
	for _, line := range lines {
		if line != "" {
			e := strings.Split(line, " ")
			if len(e) <= 2 {
				continue
			}
			version := utils.RefindNumber(e[0])
			if version != "" {
				availableList = append(availableList, version)
			}
		}
	}
	return availableList, nil
}

func (p *PythonInterpreter) GetAvailableInterpreterList() ([]string, error) {
	o1, _, err := p.ExecuteCmd("python", "list")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(o1, "\n")
	availableList := make([]string, 0)
	for _, line := range lines {
		if line != "" && !strings.Contains(line, "download available") {
			e := strings.Split(line, " ")
			if len(e) <= 2 {
				continue
			}
			availableList = append(availableList, utils.RefindNumber(e[0]))
		}
	}
	return availableList, nil
}

func (p *PythonInterpreter) InstallInterpreter(interpreter string) error {
	err := os.Setenv("UV_PYTHON_INSTALL_MIRROR", "https://ghproxy.cn/https://github.com/indygreg/python-build-standalone/releases/download")
	if err != nil {
		return err
	}
	_, _, err = p.ExecuteCmd("python", "install", interpreter)
	if err != nil {
		return err
	}
	return nil
}

func (p *PythonInterpreter) UnInstallInterpreter(interpreter string) error {
	_, _, err := p.ExecuteCmd("python", "uninstall", interpreter)
	if err != nil {
		return err
	}
	return nil
}

func (p *PythonInterpreter) SetInterpreter(interpreter string) error {
	o1, _, err := p.ExecuteCmd("python", "find", interpreter)
	if err != nil {
		return err
	}
	p.PyExecutable = strings.TrimRight(o1, "\n")
	return nil
}
