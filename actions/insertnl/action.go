package insertnl

import (
	"bufio"
	"fmt"
	"insertnewline/operators/outfile"
	"os"
	"strings"
)

const newline = "\r\n"

func Action(fpath string, n int, d bool) error {
	file, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer file.Close()
	in := bufio.NewScanner(file)
	if in.Err() != nil {
		return in.Err()
	}
	c := insertNewLine(in, n, d)
	err = outfile.File(c)
	if err != nil {
		return err
	}
	fmt.Println("finished")
	return nil
}

func insertNewLine(sc *bufio.Scanner, num int, down bool) []byte {
	var str string
	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			fmt.Println(err)
			// 最後の行なら処理終了
			break
		}
		// rsが存在すればloop発生
		// 最初のcheck対象はrsそのまま
		ck := []rune(sc.Text())
		for {
			if len(ck) == 0 {
				// 入ってるということは改行がいる
				// str += newline
				break
			}
			if len(ck) <= num {
				// 禁則チェックなし。もしくは、check対象が num 以下であればそのまま追加してbreak
				str += string(ck) + newline
				break
			}
			if !check(ck[num:num+1]) || !down {
				// 禁則チェックなし、またはcheck対象のnum+1文字目が禁則でなかったらそのまま追加、check対象を更新
				str += string(ck[:num]) + newline
				ck = ck[num:]
				continue
			}
			// check対象のnum+1文字目が禁則だったらnum+2文字目を確認
			if check(ck[num+1 : num+2]) {
				// num+2文字目が禁則だったら[:num+2]まで追加。check対象を[num+2:]に更新
				str += string(ck[:num+2]) + newline
				ck = ck[num+2:]
				continue
			}
			// num+2文字目が禁則でなかったら[:num+1]まで追加。check対象を[num+1:]に更新
			str += string(ck[:num+1]) + newline
			ck = ck[num+1:]
		}
	}
	return []byte(convert(str, newline))
}

func check(eol []rune) bool {
	list := map[string]bool{
		",": true,
		".": true,
		"、": true,
		"。": true,
		"」": true,
		"』": true,
	}
	if _, ok := list[string(eol)]; ok {
		return ok
	}
	return false
}

func convert(str, code string) string {
	return strings.NewReplacer(
		"\r\n", code,
		"\r", code,
		"\n", code,
	).Replace(str)
}
