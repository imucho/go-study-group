package chapter1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/apbgo/go-study-group/chapter1/lib"
)

// Calc opには+,-,×,÷の4つが渡ってくることを想定してxとyについて計算して返却(正常時はerrorはnilでよい)
// 想定していないopが渡って来た時には0とerrorを返却
func Calc(op string, x, y int) (int, error) {

	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	// TODO Q1
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "×":
		return x * y, nil
	case "÷":
		return x / y, nil
	}

	return 0, fmt.Errorf("invalid op=%s", op)
}

// StringEncode 引数strの長さが5以下の時キャメルケースにして返却、それ以外であればスネークケースにして返却
func StringEncode(str string) string {
	// ヒント：長さ(バイト長)はlen(str)で取得できる
	// chapter1/libのToCamelとToSnakeを使うこと

	// TODO Q2
	if len(str) <= 5 {
		return lib.ToCamel(str)
	}

	return lib.ToSnake(str)
}

// Sqrt 数値xが与えられたときにz²が最もxに近い数値zを返却
func Sqrt(x float64) float64 {

	// TODO Q3
	z := 1.0
	prev := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 1e-5 {
			break
		}
		prev = z
	}
	return z
}

// Pyramid x段のピラミッドを文字列にして返却
// 期待する戻り値の例：x=5のとき "1\n12\n123\n1234\n12345"
// （x<=0の時は"error"を返却）
func Pyramid(x int) string {
	// ヒント：string <-> intにはstrconvを使う
	// int -> stringはstrconv.Ioa() https://golang.org/pkg/strconv/#Itoa

	// TODO Q4
	if x <= 0 {
		return "error"
	}
	ret := ""
	tmp := ""
	for i := 0; i < x; i++ {
		tmp += strconv.Itoa(i + 1)
		if i != 0 {
			ret += "\n"
		}
		ret += tmp
	}

	return ret
}

// StringSum x,yをintにキャストし合計値を返却 (正常終了時、errorはnilでよい)
// キャスト時にエラーがあれば0とエラーを返却
func StringSum(x, y string) (int, error) {

	// ヒント：string <-> intにはstrconvを使う
	// string -> intはstrconv.Atoi() https://golang.org/pkg/strconv/#Atoi

	// TODO Q5
	xnum, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}
	ynum, err := strconv.Atoi(y)
	if err != nil {
		return 0, err
	}

	return xnum + ynum, nil
}

// SumFromFileNumber ファイルを開いてそこに記載のある数字の和を返却
func SumFromFileNumber(filePath string) (int, error) {
	// ヒント：ファイルの扱い：os.Open()/os.Close()
	// bufio.Scannerなどで１行ずつ読み込むと良い

	// TODO Q6 オプション
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		ret += n
	}

	return ret, nil
}
