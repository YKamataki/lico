package conf

import (
	"fmt"
	"regexp"
	"strings"

	p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
)



func ReadCreatedList() (*List, error) {
	path := vars.GetCreated()

	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}

	var list List

	for lineNo, line := range lines {
		list = append(list, NewEntryWithIndex("", p.Path(strings.TrimSpace(line)), lineNo+1))
	}

	return &list, nil
}

// 設定ファイルを読み込みます
func ReadConf() (*List, error) {
	path := vars.GetList()

	// parse config
	lines, err := FormatTemplate(path)
	if err != nil {
		return nil, err
	}

	list := List{}

	commentReg, _ := regexp.Compile("^ *#")
	emptyReg, _ := regexp.Compile("^ *$")

	for lineNo, line := range lines {
		// コメントと空行を除外
		if commentReg.MatchString(line) || emptyReg.MatchString(line) {
			continue
		}

		// :で分割
		splited := strings.Split(line, ":")
		if len(splited) >= 4 || len(splited) <= 1{   // if 1<= x <= 4; then
			return nil, fmt.Errorf("wrong syntax in line: %v", lineNo+1)
		}

		// 代入ReadConf
		repoPath := p.Path(strings.TrimSpace(splited[0]))
		homePath := p.Path(strings.TrimSpace(splited[1]))

		// 作成
		item := NewEntryWithIndex(repoPath, homePath, lineNo+1)

		// オプション付き
		if len(splited) == 3{
			option, err := ParseEntryOption(strings.TrimSpace(splited[2]))
			if err !=nil{
				return nil, err
			}
			item.Option = option
		}

		// 一覧追加
		list = append(list, item)
	}
	return &list, nil
}




