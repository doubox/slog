package handler

import (
	"os"

	"github.com/dounetio/slog"
	"github.com/gookit/color"
)

/********************************************************************************
 * console log handler
 ********************************************************************************/

// ConsoleHandler definition
type ConsoleHandler struct {
	IOWriterHandler
}

// NewConsole create new ConsoleHandler
func NewConsole(levels []slog.Level) *ConsoleHandler {
	return NewConsoleHandler(levels)
}

// NewConsoleHandler create new ConsoleHandler
func NewConsoleHandler(levels []slog.Level) *ConsoleHandler {
	h := &ConsoleHandler{
		IOWriterHandler: *NewIOWriterHandler(os.Stderr, levels),
	}

	// create new formatter
	f := slog.NewTextFormatter()
	// enable color on console
	f.EnableColor = color.SupportColor()

	h.SetFormatter(f)
	return h
}

// TextFormatter get the formatter
func (h *ConsoleHandler) TextFormatter() *slog.TextFormatter {
	return h.Formatter().(*slog.TextFormatter)
}
