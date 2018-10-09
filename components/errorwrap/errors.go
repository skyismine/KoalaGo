package errorwrap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

type Error interface {
	Caller() []CallerInfo
	Wrapped() []error
	Code() int
	error
	private()
}

type _Error struct {
	XCode int				`json:"Code"`
	XError error			`json:"Error"`
	XCaller []CallerInfo	`json:"Caller,omitempty"`
	XWrapped []error		`json:"Wraped,omitempty"`
}

type CallerInfo struct {
	FunName string
	FileName string
	FileLine int
}


//API 接口函数

func New(msg string) error {
	return &_Error{
		XError:errors.New(msg),
		XCaller:caller(2),
	}
}

func NewFrom(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(Error); ok {
		return e
	}

	return &_Error{
		XCaller:caller(2),
		XError:err,
	}
}

func Newf(format string, args ...interface{}) error {
	return &_Error{
		XError:fmt.Errorf(format, args...),
		XCaller:caller(2),
	}
}

func NewWithCode(code int, msg string) error {
	return &_Error{
		XCode:code,
		XError:errors.New(msg),
		XCaller:caller(2),
	}
}

func NewWithCodef(code int, format string, args ...interface{}) error {
	return &_Error{
		XCode:code,
		XError:fmt.Errorf(format, args...),
		XCaller:caller(2),
	}
}

func Wrap(err error, msg string) error {
	p := &_Error{
		XCaller:caller(2),
		XWrapped:[]error{err},
		XError:errors.New(fmt.Sprintf("%s -> {%v}", msg, err)),
	}
	if e, ok := err.(Error); ok {
		p.XWrapped = append(p.XWrapped, e.Wrapped()...)
	}
	return p
}

func Wrapf(err error, format string, args ...interface{}) error {
	p := &_Error{
		XCaller:caller(2),
		XWrapped:[]error{err},
		XError:errors.New(fmt.Sprintf("%s -> {%v}", fmt.Sprintf(format, args...), err)),
	}
	if e, ok := err.(Error); ok {
		p.XWrapped = append(p.XWrapped, e.Wrapped()...)
	}
	return p
}

func WrapWithCode(err error, code int, msg string) error {
	p := &_Error{
		XCaller:caller(2),
		XWrapped:[]error{err},
		XError:errors.New(fmt.Sprintf("%s -> {%v}", msg, err)),
		XCode:code,
	}
	if e, ok := err.(Error); ok {
		p.XWrapped = append(p.XWrapped, e.Wrapped()...)
	}
	return p
}

func WrapWithCodef(err error, code int, format string, args ...interface{}) error {
	p := &_Error{
		XCaller:caller(2),
		XWrapped:[]error{err},
		XError:errors.New(fmt.Sprintf("%s -> {%v}", fmt.Sprintf(format, args...), err)),
		XCode:code,
	}
	if e, ok := err.(Error); ok {
		p.XWrapped = append(p.XWrapped, e.Wrapped()...)
	}
	return p
}

func FromJson(json string) error {
	panic("not implement")
	return nil
}

func ToJson(err error) string {
	if p, ok := (err).(*_Error); ok {
		return p.String()
	}
	p := &_Error{XError:err}
	return p.String()
}


//接口实现函数

func (e *_Error) Error() (string) {

	return e.XError.Error()
}

func (e *_Error) Caller() []CallerInfo {
	return e.XCaller
}

func (e *_Error) Wrapped() []error {
	return e.XWrapped
}

func (e *_Error) Code() int {
	return e.XCode
}

func (e *_Error) String() string {
	return jsonEncodeString(e)
}

func (e *_Error) private() {

}


//工具函数
func caller(skip int) []CallerInfo {
	var infos []CallerInfo
	for ; ; skip++ {
		name, file, line, ok := callerInfo(skip+1)
		if !ok {
			return infos
		}
		if strings.HasPrefix(name, "runtime.") {
			return infos
		}
		infos = append(infos, CallerInfo{
			FunName:name,
			FileName:file,
			FileLine:line,
		})
	}
	//log unreached
}

var (
	reEmpty   = regexp.MustCompile(`^\s*$`)
	reInit    = regexp.MustCompile(`init·?\d+$`) // main.init·1
	reClosure = regexp.MustCompile(`func·?\d+$`) // main.func·001
)
func callerInfo(skip int) (name, file string, line int, ok bool) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		name = "???"
		file = "???"
		line = 1
		return
	}

	name = runtime.FuncForPC(pc).Name()
	if reInit.MatchString(name) {
		name = reInit.ReplaceAllString(name, "init")
	}
	if reClosure.MatchString(name) {
		name = reClosure.ReplaceAllString(name, "func")
	}

	//在最后一个路径分隔符处截断文件名
	if idx := strings.LastIndex(file, "/"); idx >= 0 {
		file = file[idx+1:]
	} else if idx := strings.LastIndex(file, "\\"); idx >= 0 {
		file = file[idx+1:]
	}

	return
}

func jsonEncodeIndent(m interface{}) []byte {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return nil
	}
	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1) // <
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1) // >
	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1) // &
	return data
}

func jsonEncodeString(m interface{}) string {
	return string(jsonEncodeIndent(m))
}