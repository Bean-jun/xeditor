package interpreter

import (
	_ "embed"
	"testing"
)

func TestPythonInterpreter_GetAvailableInterpreterList(t *testing.T) {
	tests := []struct {
		name string
		p    *PythonInterpreter
	}{
		// TODO: Add test cases.
		{name: "测试uv获取可用python解释器", p: &defaultPythonInterpreter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.GetAvailableInterpreterList()
		})
	}
}

func TestPythonInterpreter_InstallInterpreter(t *testing.T) {
	type args struct {
		interpreter string
	}
	tests := []struct {
		name string
		p    *PythonInterpreter
		args args
	}{
		// TODO: Add test cases.
		{name: "测试uv获取可用python解释器", p: &defaultPythonInterpreter, args: args{interpreter: "3.7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.InstallInterpreter(tt.args.interpreter)
		})
	}
}

func TestPythonInterpreter_SetInterpreter(t *testing.T) {
	type args struct {
		interpreter string
	}
	tests := []struct {
		name    string
		p       *PythonInterpreter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", p: &defaultPythonInterpreter, args: args{interpreter: "3.10"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.SetInterpreter(tt.args.interpreter); (err != nil) != tt.wantErr {
				t.Errorf("PythonInterpreter.SetInterpreter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPythonInterpreter_GetAllInterpreterList(t *testing.T) {
	tests := []struct {
		name string
		p    *PythonInterpreter
	}{
		// TODO: Add test cases.
		{name: "测试uv获取可用python解释器", p: &defaultPythonInterpreter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.GetAllInterpreterList()
		})
	}
}
