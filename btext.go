package btext

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BParseFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%s nofile %s \n", err, path)
		return make([]byte, 0)
	}

	defer f.Close()
	lines := make([]byte, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buffer := BParseLine(scanner.Text())
		lines = append(lines, buffer...)
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "err %s scan error: %v\n", path, err)
	}
	return lines
}

func BParseLine(text string) []byte {
	textData := eraceComment(text)
	if len(textData) == 0 {
		return nil
	}
	var readBuffer []byte
	for _, v := range textData {
		n, err := strconv.ParseInt(v, 16, 16)
		if err == nil {
			readBuffer = append(readBuffer, byte(n))
		}
	}
	return readBuffer
}

func eraceComment(text string) []string {
	idx := strings.Index(text, ";")
	if idx == 0 {
		return make([]string, 0)
	} else if idx != -1 {
		text = text[0:idx]
	}
	return strings.Split(text, " ")
}

const padMax = 8

func TParseAry(binary []byte) string {
	text := addHeader()
	for i, v := range binary {
		val := fmt.Sprintf("%02x", v)
		if ((i + 1) % 16) == 0 {
			text += val + "\n"
		} else if (i % 16) == 0 {
			idxStr := fmt.Sprintf("0x%02x", i)
			padSize := 8 - len(idxStr)
			text += idxStr + strings.Repeat(" ", padSize) + val + " "
		} else {
			text += val + " "
		}
	}
	return text
}

func addHeader() string {
	heads := make([]string, 16)
	for i := 0; i < 16; i++ {
		heads[i] = fmt.Sprintf("%02x", i)
	}
	return strings.Repeat(" ", padMax) + strings.Join(heads, "-") + "\n"
}
