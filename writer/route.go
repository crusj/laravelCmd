package writer

import "strings"

type Method string

var ValidMethods = map[string]Empty{
	"GET":    {},
	"POST":   {},
	"PUT":    {},
	"DELETE": {},
}

func (method Method) IsValid() bool {
	if _, exists := ValidMethods[strings.ToUpper(string(method))]; exists {
		return true
	} else {
		return false
	}
}
func (method Method) String() string {
	return string(method)
}

type RouteRequest struct {
	//名称
	Name,
	//请求路径
	Path string
	//请求方法
	Method Method
	//请求参数
	Params map[string]interface{}

	//模块名
	ModuleName,
	//模块标题
	ModuleTitle,
	//动作名
	ActionName,
	//动作标题
	ActionTitle string
	SetId bool
}
type RouteResponse struct {
	//响应类型
	Type string
	//响应数据
	Body interface{}
}

type Route struct {
	RouteRequest
	RouteResponse
}
