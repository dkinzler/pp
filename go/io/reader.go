package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LineReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewLineReader(filename string) (*LineReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	return &LineReader{file: file, scanner: scanner}, nil
}

func (l *LineReader) ReadLine() (string, error) {
	if ok := l.scanner.Scan(); !ok {
		if l.scanner.Err() == nil {
			return "", ErrorEOF
		}
		return "", l.scanner.Err()
	}
	return l.scanner.Text(), nil
}

func (l *LineReader) Close() error {
	return l.file.Close()
}

var ErrorEOF = errors.New("EOF")

func (l *LineReader) ReadInt() (int, error) {
	line, err := l.ReadLine()
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (l *LineReader) ReadIntSlice() ([]int, error) {
	line, err := l.ReadLine()
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	elements := strings.Split(line, ",")
	result := make([]int, len(elements))
	for i, e := range elements {
		e = strings.TrimSpace(e)
		n, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}

func (l *LineReader) ReadStringSlice() ([]string, error) {
	line, err := l.ReadLine()
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	elements := strings.Split(line, ",")
	result := make([]string, len(elements))
	for i, e := range elements {
		e = strings.TrimSpace(e)
		e = strings.TrimPrefix(e, "\"")
		e = strings.TrimSuffix(e, "\"")
		result[i] = e
	}
	return result, nil
}
