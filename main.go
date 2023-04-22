package main

import (
	"bufio"
	"os"
	"strings"
	"time"
)

// "02:42:03.565772 /root/lib/engine/compare.go:29: Comparing victim to suspect"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	pathTimes := make(map[string]time.Time)

	for scanner.Scan() {
		line := scanner.Text()
		timestr := line[:15]

		logTime, err := time.Parse("15:04:05.999999", timestr)
		if err == nil {
			lastColon := strings.LastIndex(line[15:], ":")
			logPath := line[15:lastColon]

			lastTime, ok := pathTimes[logPath]
			if ok {
				diff := logTime.Sub(lastTime)
				os.Stdout.WriteString(line + " (" + diff.String() + ")\n")
			}

			pathTimes[logPath] = logTime
		} else {
			os.Stdout.WriteString(line + "\n")
		}

	}
}
