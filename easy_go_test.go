package easygo

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewEasyGo(t *testing.T) {
	easyGo2 := NewEasyGo()
	easyGo2.maxConn = 2
	easyGo2.timeOut = 1
	easyGo2.name = "custom_name"
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *EasyGo
	}{
		{
			name: "null option",
			args: args{options: []Option{}},
			want: NewEasyGo(),
		},
		{
			name: "input option",
			args: args{
				options: []Option{
					WithMaxConn(2), WithTimeOut(1), WithName("custom_name")},
			},
			want: easyGo2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEasyGo(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEasyGo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEasyGo_Run(t *testing.T) {
	type fields struct {
		baseServer            *http.Server
		serveHandler          http.Handler
		name                  string
		port                  string
		ctx                   easyGoCtx
		timeOut               time.Duration
		maxConn               int64
		appConfigYamlFilePath string
		DEBUG                 bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO.md: Add test cases.
		{
			name: "easy-go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			easyGo := &EasyGo{
				baseServer:            tt.fields.baseServer,
				serveHandler:          tt.fields.serveHandler,
				name:                  tt.fields.name,
				port:                  tt.fields.port,
				ctx:                   tt.fields.ctx,
				timeOut:               tt.fields.timeOut,
				maxConn:               tt.fields.maxConn,
				appConfigYamlFilePath: tt.fields.appConfigYamlFilePath,
			}
			easyGo.Run()
		})
	}
}
