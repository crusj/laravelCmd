package writer

import (
	"fmt"
	"github.com/stoewer/go-strcase"
	"strings"
)

type controllerWriter struct {
	Dir,
	StartFlag,
	EndFlag string
}

func NewControllerWriter(dir, startFlag, EndFlag string) *controllerWriter {
	return &controllerWriter{
		Dir:       dir,
		StartFlag: startFlag,
		EndFlag:   EndFlag,
	}
}
func (controllerWriter *controllerWriter) FilesPath(routes []Route) []string {
	filesPath := make([]string, 0)

	hash := make(map[string]Empty)
	for _, route := range routes {
		fileName := strcase.UpperCamelCase(route.ModuleName) + ".php"
		if _, exists := hash[fileName]; exists {
			continue
		} else {
			filesPath = append(filesPath, controllerWriter.Dir+fileName)
			hash[fileName] = Empty{}
		}
	}
	return filesPath
}
func (controllerWriter *controllerWriter) Between() *Between {
	return &Between{
		StartFlag: controllerWriter.StartFlag,
		EndFlag:   controllerWriter.EndFlag,
	}
}
func (controllerWriter *controllerWriter) Contents(routes []Route) [][]string {
	eightSpace := strings.Repeat(" ", 8)
	fourSpace := strings.Repeat(" ", 4)
	tweleveSpace := strings.Repeat(" ", 12)
	funcTpl := fourSpace + "//%s\n" +
		fourSpace + "public function %s(%s) {\n" +
		eightSpace + "%s" +
		eightSpace + "%s" +
		eightSpace + "%s" +
		fourSpace + "}"
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
		//serviceFactory
		serviceInstance := fmt.Sprintf("$service = ServiceFactory::%s();\n", strcase.LowerCamelCase(route.ModuleName))
		//resultName
		resultName := ""
		switch strings.ToUpper(route.Method.String()) {
		case "GET":
			resultName = "data"
		case "POST":
			resultName = "id"
		case "PUT", "DELETE":
			resultName = "ok"
		}
		//rsl
		actionRsl := fmt.Sprintf("$%s = $service->%s(%s);\n", resultName, actionName, idStr)
		//check and return
		ifStr := ""
		switch strings.ToUpper(route.Method.String()) {
		case "GET":
			ifStr = "if(empty($%s)) {\n" + tweleveSpace + "$this->%s('暂无数据');\n" +
				eightSpace + "}else {\n" +
				tweleveSpace + "$this->success($%s);\n" +
				eightSpace + "}\n"
			if route.SetId {
				ifStr = fmt.Sprintf(ifStr, resultName, "failObject", resultName)
			} else {
				ifStr = fmt.Sprintf(ifStr, resultName, "failArray", resultName)
			}
		case "POST":
			ifStr = "if($%s <= 0) {\n" + tweleveSpace + "$this->failObject($service->getErrorMsg());\n" +
				eightSpace + "}else {\n" +
				tweleveSpace + "$this->success(['id' => $%s]);\n" +
				eightSpace + "}\n"
			ifStr = fmt.Sprintf(ifStr, resultName, resultName)
		case "PUT":
			ifStr = "if($%s === false) {\n" + tweleveSpace + "$this->failObject($service->getErrorMsg());\n" +
				eightSpace + "}else {\n" +
				tweleveSpace + "$this->successNull();\n" +
				eightSpace + "}\n"
			ifStr = fmt.Sprintf(ifStr, resultName)
		case "DELETE":
			ifStr = "if($%s === false) {\n" + tweleveSpace + "$this->failObject($service->getErrorMsg());\n" +
				eightSpace + "}else {\n" +
				tweleveSpace + "$this->successNull();\n" +
				eightSpace + "}\n"
			ifStr = fmt.Sprintf(ifStr, resultName)
		}

		function := fmt.Sprintf(funcTpl,
			route.ActionTitle,
			actionName,
			idStr,
			serviceInstance,
			actionRsl,
			ifStr,
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
