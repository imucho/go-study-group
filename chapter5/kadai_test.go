package chapter5

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		assert.NoError(t, Validate(1, 1))
	})
	t.Run("異常系　ファイル未指定", func(t *testing.T) {
		t.Parallel()
		assert.EqualError(t, Validate(0, 1), "ファイルパスを指定してください。")
	})
	t.Run("異常系　-f 1未満", func(t *testing.T) {
		t.Parallel()
		assert.EqualError(t, Validate(1, 0), "-f は1以上である必要があります")
	})
	t.Run("異常系　-f 1未満、ファイル未指定", func(t *testing.T) {
		t.Parallel()
		assert.EqualError(t, Validate(0, 0), "ファイルパスを指定してください。")
	})
}

func TestCut(t *testing.T) {
	csvString := "1,GoodAfternoon ,Illustrator,FALSE\n2,Hi,Gopher,TRUE\n3,GoodMorning,Doctor,TRUE\n4,Hello,Gopher,FALSE\n5,GoodEvening , Singer,TRUE"
	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader(csvString)
		writer := &bytes.Buffer{}
		err := Cut(reader, writer, ",", 2)
		assert.NoError(t, err)
		expected := "GoodAfternoon \nHi\nGoodMorning\nHello\nGoodEvening \n"
		assert.Equal(t, expected, writer.String())
	})
	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader(csvString)
		writer := &bytes.Buffer{}
		err := Cut(reader, writer, "\t", 1)
		assert.NoError(t, err)
		expected := "1,GoodAfternoon ,Illustrator,FALSE\n2,Hi,Gopher,TRUE\n3,GoodMorning,Doctor,TRUE\n4,Hello,Gopher,FALSE\n5,GoodEvening , Singer,TRUE\n"
		assert.Equal(t, expected, writer.String())
	})
	t.Run("異常系　範囲外", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader(csvString)
		writer := &bytes.Buffer{}
		assert.EqualError(t, Cut(reader, writer, ",", 5), "-fの値に該当するデータがありません")
	})
}

func BenchmarkCut(b *testing.B) {
	b.ResetTimer()
	reader := strings.NewReader("1,GoodAfternoon ,Illustrator,FALSE\n2,Hi,Gopher,TRUE\n3,GoodMorning,Doctor,TRUE\n4,Hello,Gopher,FALSE\n5,GoodEvening , Singer,TRUE")
	writer := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		Cut(reader, writer, ",", 2)
	}
}
