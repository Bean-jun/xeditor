package controller

import (
	"xeditor/app"
	"xeditor/app/interpreter"
	"xeditor/app/utils"
)

type PythonController struct {
	app         *app.App
	interpreter *interpreter.PythonInterpreter
}

func NewPythonController(app *app.App) *PythonController {
	return &PythonController{
		app:         app,
		interpreter: interpreter.NewPythonInterpreter(),
	}
}

// GetInterpreterList 获取解释器状态信息
func (p *PythonController) GetInterpreterList() map[string]interface{} {
	result := make(map[string]interface{})
	result["available"] = make([]string, 0)
	result["current"] = utils.RefindNumber(p.interpreter.PyExecutable)

	list, err := p.interpreter.GetAllInterpreterList()
	if err != nil {
		return nil
	}

	available, err := p.interpreter.GetAvailableInterpreterList()
	if err != nil {
		return nil
	}

	result["list"] = list
	result["available"] = available
	return result
}

// InstallInterpreter 安装指定版本解释器
func (p *PythonController) InstallInterpreter(interpreter string) bool {
	err := p.interpreter.InstallInterpreter(interpreter)
	return err == nil
}

// 设置指定版本解释器为默认解释器
func (p *PythonController) SetInterpreter(interpreter string) bool {
	err := p.interpreter.SetInterpreter(interpreter)
	return err == nil
}

// 卸载指定版本解释器
func (p *PythonController) UnInstallInterpreter(interpreter string) bool {
	err := p.interpreter.UnInstallInterpreter(interpreter)
	return err == nil
}
