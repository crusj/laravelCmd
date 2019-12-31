package writer

import (
	"fmt"
	myFile "github.com/crusj/file"
	"github.com/crusj/logger"
	"math/rand"
	"time"
)

type Empty struct{}

type Between struct {
	StartFlag,
	EndFlag,
	startTag,
	endTag string
}
type Writer interface {
	FilePath() string
	Between() *Between
	Content([]Route) []string
}

func Write(parser Parser, writer Writer) error {
	file, err := myFile.NewFile(writer.FilePath())
	if err != nil {
		return err
	}
	between := writer.Between()
	err = file.Scan([]myFile.AddTag{between.addStartTag, between.addEndTag}...)
	if err != nil {
		return err
	}
	err = file.InsertBetween(between.startTag, between.endTag, writer.Content(parser.Parse()))
	if err != nil {
		return err
	}
	logger.Info("ok 写入成功")
	return nil
}
func (between *Between) addStartTag(line int, content string) string {
	if content == between.StartFlag {
		rand.Seed(time.Now().Unix())
		between.startTag = fmt.Sprintf("%d_%s", rand.Intn(1000000), between.StartFlag)
		return between.startTag
	}
	return ""
}
func (between *Between) addEndTag(line int, content string) string {
	if content == between.EndFlag {
		rand.Seed(time.Now().Unix())
		between.endTag = fmt.Sprintf("%d_%s", rand.Intn(1000000), between.EndFlag)
		return between.endTag
	}
	return ""
}

//向多个文件写入数据
type Writers interface {
	FilesPath([]Route) []string
	Between() *Between
	Contents([]Route) [][]string
}

func Writes(parser Parser, writers Writers) error {

	routes := parser.Parse()
	contents := writers.Contents(routes)
	for i, filePath := range writers.FilesPath(routes) {
		file, err := myFile.NewFile(filePath)
		if err != nil {
			logger.Error("文件路径错误%s,将被忽略,%s", filePath, err)
			continue
		}
		if contents[i] == nil {
			logger.Warn("文件路径%s,内容为空将被忽略", filePath)
			continue
		}
		between := writers.Between()
		err = file.Scan([]myFile.AddTag{between.addStartTag, between.addEndTag}...)
		if err != nil {
			logger.Error("文件%s扫描错误,将被忽略,%s", filePath, err)
			continue
		}
		err = file.InsertBetween(between.startTag, between.endTag, contents[i])
		if err != nil {
			logger.Error("文件%插入错误,将被忽略,%s", filePath, err)
			continue
		}
		logger.Info("文件%s写入成功", filePath)
	}
	return nil
}
