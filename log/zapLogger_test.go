package log

import (
	"testing"
)

//func TestZapLogger_Debug(t *testing.T) {
//	NewZapperLogger().Error("AAAAAAA")
//}

func TestStr(t *testing.T)  {

	for i := 0; i < 100000; i++ {
		//fmt.Sprintf("%s%s", a, b)
		NewZapperLogger().Error("AAAAAAA")
	}
}