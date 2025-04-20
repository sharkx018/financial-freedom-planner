package logger

import (
	"context"
	"fmt"
)

func LogInfo(ctx context.Context, message string, args ...interface{}) {
	fmt.Println(message)
}

func LogError(ctx context.Context, message string, args ...interface{}) {
	fmt.Println(message)
}
