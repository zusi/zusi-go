package gen

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"go/format"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	PkgPath    = "github.com/zusi/zusi-go"
	TcpPath    = PkgPath + "/tcp"
	GenPath    = TcpPath + "/gen"
	CommonPath = TcpPath + "/common"
	PkgErrs    = "github.com/pkg/errors"
)

func Generate(msg Message, writer io.Writer) error {
	t := template.New("").Funcs(template.FuncMap{
		"title": strings.Title,
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

	log.WithField("filename", fileName).Info("successfully wrote file")

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
