package exprint

import (
	"bufio"
	"fmt"
	"io/fs"
	"strings"

	"github.com/lithammer/dedent"
)

type EX struct {
	fs           fs.FS
	commentStart string
	commentEnd   string
}

func New(fs fs.FS, commentStart, commentEnd string) *EX {
	return &EX{
		fs:           fs,
		commentStart: commentStart,
		commentEnd:   commentEnd,
	}
}

func (e *EX) Print(filePath, name string) (string, error) {
	file, err := e.fs.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		started bool
		ended   bool
		output  []string
	)

	startComment := fmt.Sprintf("%sex:start:%s%s", e.commentStart, name, e.commentEnd)
	endComment := fmt.Sprintf("%sex:end:%s%s", e.commentStart, name, e.commentEnd)

	for scanner.Scan() {
		if !started {
			if strings.TrimSpace(scanner.Text()) == startComment {
				started = true
			}
			continue
		}

		if strings.TrimSpace(scanner.Text()) == endComment {
			ended = true
			break
		}

		output = append(output, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return "", fmt.Errorf("failed to scan file: %w", err)
	}

	if !started {
		return "", fmt.Errorf("%s not found", startComment)
	}
	if !ended {
		return "", fmt.Errorf("%s not found", endComment)
	}

	contents := strings.Join(output, "\n")
	return dedent.Dedent(contents), nil
}

func (e *EX) PrintOrErr(filePath, name string) string {
	contents, err := e.Print(filePath, name)
	if err != nil {
		return err.Error()
	}

	return contents
}
