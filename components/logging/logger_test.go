package logging

import "testing"

func TestLogger(t *testing.T) {
	logger := New("/Users/erikwu/Desktop/Project/KoalaGo/components/logging/app.log")
	logger.Info("info message")
	logger.Warnning("warnning message")
	logger.Debug("debug message")
	logger.Error("error message")
}

func BenchmarkLogger(b *testing.B) {
	logger := New("/Users/erikwu/Desktop/Project/KoalaGo/components/logging/app.log")
	for i := 0; i < b.N; i++ {
		logger.Info("info message")
		logger.Warnning("warnning message")
		logger.Debug("debug message")
		logger.Error("error message")
	}
}