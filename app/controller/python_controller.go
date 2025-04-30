package controller

import (
	"xeditor/app"
	"xeditor/app/interpreter"
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

func (p *PythonController) GetInterpreterList() map[string]interface{} {
	result := make(map[string]interface{})
	result["available"] = make([]string, 0)
	result["current"] = p.interpreter.PyExecutable

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
