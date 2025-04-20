package Utility

import (
	"fmt"
	"os"
)

var FileInstance *os.File

func InitiateLogs() {
	newFileInstance, err := os.Create("DurationLog.txt")
	if err != nil {
		panic(err)
	}
	FileInstance = newFileInstance
}

func LogDuartion(TimeLog string) {
	TimeLog += "/n"
	numbytesWritten, err := FileInstance.WriteString(TimeLog)
	if err != nil {
		panic(err)
	}
	fmt.Println(numbytesWritten)
}
