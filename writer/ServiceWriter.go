package writer

import (
	"fmt"
	"github.com/stoewer/go-strcase"
	"strings"
)

//服务文件方法写入
type serviceWriter struct {
	Dir,
	StartFlag,
	EndFlag string
}

func (serviceWriter serviceWriter) FilesPath(routes []Route) []string {
	filesPath := make([]string, 0)

	hash := make(map[string]Empty)
	for _, route := range routes {
		fileName := strcase.UpperCamelCase(route.ModuleName) + ".php"
		if _, exists := hash[fileName]; exists {
			continue
		} else {
			filesPath = append(filesPath, serviceWriter.Dir+fileName)
			hash[fileName] = Empty{}
		}
	}
	return filesPath
}
func (serviceWriter serviceWriter) Between() *Between {
	return &Between{
		StartFlag: serviceWriter.StartFlag,
		EndFlag:   serviceWriter.EndFlag,
	}
}
func (serviceWriter serviceWriter) Contents(routes []Route) [][]string {
	funcTpl := "    //%s\n    public function %s(%s) :%s {\n        return %s;\n    }"
	hash := make(map[string][]string)
	hashMap := make([]string, 0)
	for _, route := range routes {
		//function name
		actionName := ""
		if route.ActionName == "" {
			switch strings.ToUpper(route.Method.String()) {
			case "GET":
				actionName = "index"
			case "POST":
				actionName = "store"
			case "PUT":
				actionName = "update"
			case "DELETE":
				actionName = "destroy"
			}
		} else {
			actionName = route.ActionName
			switch strings.ToUpper(route.Method.String()) {
			case "POST":
				actionName = "store" + strcase.UpperCamelCase(actionName)
			case "PUT":
				actionName = "update" + strcase.UpperCamelCase(actionName)
			case "DELETE":
				actionName = "destroy" + strcase.UpperCamelCase(actionName)
			}
		}
		//id
		idStr := ""
		if route.SetId == true {
			idStr = "int $id"
		}
		//return type
		returnType := ""
		defaultReturnValue := ""
		switch strings.ToUpper(route.Method.String()) {
		case "GET":
			returnType = "array"
			defaultReturnValue = "[]"
		case "POST":
			returnType = "int"
			defaultReturnValue = "0"
		case "PUT", "DELETE":
			returnType = "bool"
			defaultReturnValue = "false"
		}
		function := fmt.Sprintf(funcTpl,
			route.ActionTitle,
			actionName,
			idStr,
			returnType,
			defaultReturnValue,
		)
		if _, exists := hash[route.ModuleName]; exists {
			hash[route.ModuleName] = append(hash[route.ModuleName], function)
		} else {
			hash[route.ModuleName] = []string{function}
			hashMap = append(hashMap, route.ModuleName)
		}
	}
	actions := make([][]string, 0)
	for _, moduleName := range hashMap {
		actions = append(actions, hash[moduleName])
	}
	return actions
}

func NewServiceWriter(dir, startFlag, EndFlag string) *serviceWriter {
	return &serviceWriter{
		Dir:       dir,
		StartFlag: startFlag,
		EndFlag:   EndFlag,
	}
}
