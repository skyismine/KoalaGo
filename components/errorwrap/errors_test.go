package errorwrap

import (
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	err := New("test1")
	if err.Error() != "test1" {
		t.Fatal("test1 fail")
	}
}

func BenchmarkErrors(b *testing.B) {
	b.StopTimer()
	//此处添加一些不计入耗时的操作
	//...
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := New(fmt.Sprintf("benchmarkTest-----%d", i+1))
		err = Wrapf(err, "wrapf-----%d", i+1)
		fmt.Println(ToJson(err))
	}
}
