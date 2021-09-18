package client

import (
	"fmt"
	"io"
)

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
	}
}

func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))

	return err
}

func (w *CommandWriter) Write(command interface{}) error {
	var err error
	switch v := command.(type) {
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("tag  %v %v\n", v.Name, v.Message))
	case SendCommand:
		err = w.writeString(fmt.Sprintf("broad %v\n", v.Message))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("name %v\n", v.Name))
	default:
		err = ErrC
	}

	return err
}
