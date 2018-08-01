package main

import (
	"bytes"
	"fmt"
	"strings"
)

const x = "hello"

func withPlus() string {
	return x + x + x + x + x + x + x + x + x + x
}

func withSprintf() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func withBuffer() string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func withStringBuilder() string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}
