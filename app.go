package main

import (
	"context"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Encrypt(plaintext, key string) string {
	text := strings.ToUpper(strings.ReplaceAll(plaintext, " ", ""))
	key = strings.ToUpper(key)
	textRunes := []rune(text)
	cols := len(key)
	rows := len(text) / cols
	if len(text)%cols != 0 {
		rows++
	}

	matrix := make([][]rune, rows)
	for i := range matrix {
		matrix[i] = make([]rune, cols)
	}
	index := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if index < len(text) {
				matrix[i][j] = textRunes[index]
				index++
			} else {
				matrix[i][j] = 'X'
			}
		}
	}

	order := getColOrder(key)
	var result strings.Builder
	result.Grow(rows * cols)
	for _, col := range order {
		for i := 0; i < rows; i++ {
			result.WriteRune(matrix[i][col])
		}
	}

	return result.String()
}

func (a *App) Decrypt(text, keyword string) string {
	text = strings.ToUpper(text)
	keyword = strings.ToUpper(keyword)

	cols := len(keyword)
	rows := len(text) / cols
	if len(text)%cols != 0 {
		rows++
	}

	order := getColOrder(keyword)

	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	index := 0
	for _, col := range order {
		for i := 0; i < rows; i++ {
			if index < len(text) {
				matrix[i][col] = text[index]
				index++
			}
		}
	}

	var result strings.Builder
	result.Grow(rows * cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result.WriteByte(matrix[i][j])
		}
	}

	return strings.TrimRight(result.String(), "X")
}

func getColOrder(key string) []int {
	keyRunes := []rune(key)
	type keyChar struct {
		pos  int
		char rune
	}

	pairs := make([]keyChar, len(key))
	for i := 0; i < len(key); i++ {
		pairs[i] = keyChar{
			pos:  i,
			char: keyRunes[i],
		}
	}

	for i := 0; i < len(pairs)-1; i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i].char > pairs[j].char {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}

	order := make([]int, len(key))
	for i, ch := range pairs {
		order[i] = ch.pos
	}
	return order
}

func (a *App) OpenFile() string {
	file, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if file == "" {
		return ""
	}
	data, _ := os.ReadFile(file)
	return string(data)
}

func (a *App) SaveFile(content string) string {
	file, _ := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{})
	if file == "" {
		return ""
	}
	os.WriteFile(file, []byte(content), 0644)
	return "OK"
}
