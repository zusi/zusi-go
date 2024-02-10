package gen

import (
	"bytes"
	"go/format"
	"io"
	"log/slog"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"
)

const (
	PkgPath    = "github.com/zusi/zusi-go"
	TcpPath    = PkgPath + "/tcp"
	GenPath    = TcpPath + "/gen"
	CommonPath = TcpPath + "/common"
	PkgErrs    = "errors"
)

func Title(s string) string {
	if len(s) == 0 {
		return ""
	}

	var output []rune
	output = append(output, unicode.ToUpper(rune(s[0])))
	output = append(output, []rune(s[1:])...)

	return string(output)
}

func Generate(msg Message, writer io.Writer) error {
	t := template.New("").Funcs(template.FuncMap{
		"title": Title,
		"pkg": func(m string) string {
			pkgs := strings.Split(m, "/")

			return pkgs[len(pkgs)-1]
		},
	})

	_, err := t.Parse(clientTemplate())
	if err != nil {
		return err
	}

	return t.Execute(writer, msg)
}

type header struct {
	Imports map[string]struct{}
}

func generateHeader(msgs []Message, writer io.Writer) error {
	t := template.New("")

	_, err := t.Parse(headerTemplate())
	if err != nil {
		return err
	}

	imports := map[string]struct{}{}

	for _, msg := range msgs {
		imports[msg.PkgPath] = struct{}{}
	}

	he := header{Imports: imports}

	return t.Execute(writer, he)
}

func WriteFile(messages []Message, rootPath string) error {
	outBuffer := &bytes.Buffer{}
	err := GenerateReaderFile(messages, outBuffer)
	if err != nil {
		return err
	}

	result, err := format.Source(outBuffer.Bytes())
	if err != nil {
		return err
	}

	fileName := path.Join(rootPath, "tcp", "reader_gen.go")
	writer, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = writer.Write(result)
	if err != nil {
		writer.Close()
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	slog.With("filename", fileName).Info("wrote file")

	return nil
}

func GenerateReaderFile(messages []Message, writer io.Writer) error {
	err := generateHeader(messages, writer)
	if err != nil {
		return err
	}

	for _, v := range messages {
		err = Generate(v, writer)
		if err != nil {
			return err
		}
	}

	return nil
}
