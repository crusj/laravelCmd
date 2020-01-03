package writer

import (
	"fmt"
	"github.com/crusj/logger"
	"github.com/stoewer/go-strcase"
	"os"
	"os/exec"
	"strings"
)

type requestAttributeWriter struct {
	Dir,
	StartFlag,
	EndFlag string
}

func NewRequestAttributeWriter(dir, startFlag, EndFlag string) *requestAttributeWriter {
	return &requestAttributeWriter{
		Dir:       dir,
		StartFlag: startFlag,
		EndFlag:   EndFlag,
	}
}
func (requestWriter *requestAttributeWriter) FilesPath(routes []Route) []string {
	filesPath := make([]string, 0)

	hash := make(map[string]Empty)
	for _, route := range routes {
		switch strings.ToUpper(route.Method.String()) {
		case "POST":
			fileName := strcase.UpperCamelCase("store_"+route.ModuleName+"_"+route.ActionName) + ".php"
			if route.ActionName == "" {
				fileName = strcase.UpperCamelCase("store_"+route.ModuleName) + ".php"
			}
			ruleName := strings.Split(fileName, ".")[0]
			if _, exists := hash[fileName]; exists {
				continue
			} else {
				filesPath = append(filesPath, requestWriter.Dir+fileName)
				hash[fileName] = Empty{}
			}
			//文件不存在则创建
			if _, err := os.Stat(requestWriter.Dir + fileName); os.IsNotExist(err) {
				cmd := exec.Command("php", "artisan", "make:requestAttribute", ruleName)
				rsl, err := cmd.Output()
				if string(rsl) == "Request created successfully." {
					logger.Info("创建request,%s", ruleName)
				} else {
					logger.Warn("创建request失败,%s,%s", ruleName, err)
				}
			}
		case "PUT":
			fileName := strcase.UpperCamelCase("update_"+route.ModuleName+"_"+route.ActionName) + ".php"
			if route.ActionName == "" {
				fileName = strcase.UpperCamelCase("update_"+route.ModuleName) + ".php"
			}
			if _, exists := hash[fileName]; exists {
				continue
			} else {
				filesPath = append(filesPath, requestWriter.Dir+fileName)
				hash[fileName] = Empty{}
			}
		default:
			continue
		}
	}
	return filesPath
}
func (requestWriter *requestAttributeWriter) Between() *Between {
	return &Between{
		StartFlag: requestWriter.StartFlag,
		EndFlag:   requestWriter.EndFlag,
	}
}
func (requestWriter *requestAttributeWriter) Contents(routes []Route) [][]string {

	contents := make([][]string, 0)
	for _, route := range routes {
		//function name
		methodName := strings.ToUpper(route.Method.String())

		switch methodName {
		case "POST", "PUT":
			str := make([]string, 0)
			for _, rp := range route.RequestParam {
				r := &requestAttribute{
					Name:        rp.Name,
					Description: rp.Desc,
				}
				str = append(str, r.String())
			}
			contents = append(contents, str)
		default:
			continue
		}
	}
	return contents
}

type requestAttribute struct {
	Name        string
	Description string
}

func (request *requestAttribute) String() string {
	space12 := strings.Repeat(" ", 12)
	tpl := space12 + "'%s' => '%s',"
	descS := strings.Split(request.Description, "|")
	desc := ""
	if len(descS) > 1 {
		desc = strings.Join(descS[1:], ",")
	} else {
		desc = append(descS, request.Name)[0]
	}
	return fmt.Sprintf(tpl, request.Name, desc)

}
