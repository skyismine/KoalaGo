####测试命名
#####功能和性能测试文件名以_test结尾，如:errors_test.go
#####功能测试函数名以Test开头，如: func TestError(t *testing.T)
#####性能测试函数名以Benchmark开头，如：func BenchmarkError(b *testing.B)

####测试运行
#####切换到测试文件所在的目录执行 go test 或 go test -test.bench=".*"