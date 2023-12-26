package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(input_fn string) []string {
    fd, err := os.Open(input_fn)

    if err != nil {
        log.Fatal(err)
    }

    defer fd.Close()

    reader := bufio.NewReader(fd)

    var lines []string
    for {
        line, err := reader.ReadString('\n')

        if err != nil && err.Error() == "EOF" {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        lines = append(lines, strings.TrimRight(line, "\n"))
    }

    return lines 
}
