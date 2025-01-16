package formats

import (
	"fmt"
	"os"
	"time"
)

func LongListFormat(results []string) []string {
	var longList []string
	for _, file := range results {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		permissions := info.Mode().String()
		modTime := info.ModTime().Format(time.RFC822)
		size := info.Size()

		longList = append(longList, fmt.Sprintf("%c %s %10d %s %s", permissions[0], permissions[1:], size, modTime, file))
	}
	return longList
}
