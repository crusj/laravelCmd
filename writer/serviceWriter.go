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
	funcTpl := "    //%s\n    public function %s(%s) :%s {\n%s\n        return %s;\n    }"
	hash := make(map[string][]string)
	hashMap := make([]string, 0)
	for _, route := range routes {
		//function name
		actionName := ""
		if route.ActionName == "" {
			switch strings.ToUpper(route.Method.String()) {
			case "GET":
				if route.SetId {
					actionName = "show"
				}else{
					actionName = "index"
				}
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
			queryParamToCode(route.QueryParams),
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
func queryParamToCode(queryParams []*QueryParam) string {

	var queryParamsStr []string
	eightSpace := strings.Repeat(" ", 8)
	tpl := eightSpace+"//%s\n" + eightSpace + "$%s = request()->get('%s', %s);"
	for _, queryParam := range queryParams {
		descEach := strings.Split(queryParam.Desc, "|")
		//参数类型
		paramType := ""
		//参数默认值
		paramDefaut := ""
		//参数描述
		paramDesc := ""
		switch len(descEach) {
		case 0:
		case 1:
			paramDesc = descEach[0]
		case 2:
			paramDesc = descEach[1]
			paramType = strings.ToLower(descEach[0])
		case 3:
			paramDesc = descEach[2]
			paramType = strings.ToLower(descEach[0])
			paramDefaut = descEach[1]
		default:
			paramType = strings.ToLower(descEach[0])
			paramDefaut = descEach[1]
			paramDesc = strings.Join(descEach[2:], "")
		}
		if paramDefaut == "" {
			if paramType == "" {
				paramDefaut = "''"
			} else {
				switch paramType {
				case "string":
					paramDefaut = "''"
				case "int":
					paramDefaut = "0"
				case "bool":
					paramDefaut = "false"
				default:
					paramDefaut = "''"
				}
			}
		}
		tplStr := fmt.Sprintf(tpl,
			paramDesc,
			strcase.LowerCamelCase(queryParam.Name),
			queryParam.Name,
			paramDefaut,
		)
		queryParamsStr = append(queryParamsStr, tplStr)
	}
	return strings.Join(queryParamsStr, "\n")
}
func NewServiceWriter(dir, startFlag, EndFlag string) *serviceWriter {
	return &serviceWriter{
		Dir:       dir,
		StartFlag: startFlag,
		EndFlag:   EndFlag,
	}
}
