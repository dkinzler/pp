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
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)
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
	return l.parseIntSlice(line)
}

func (l *LineReader) parseIntSlice(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if len(s) == 0 {
		return nil, nil
	}
	elements := strings.Split(s, ",")
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

func (l *LineReader) Read2DimIntSlice() ([][]int, error) {
	line, err := l.ReadLine()
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	if len(line) == 0 {
		return nil, nil
	}
	elements := strings.Split(line, "],")
	result := make([][]int, len(elements))
	for i, e := range elements {
		x, err := l.parseIntSlice(e)
		if err != nil {
			return nil, err
		}
		result[i] = x
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
	if len(line) == 0 {
		return nil, nil
	}
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
