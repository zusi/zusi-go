package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"sort"

	"github.com/zusi/zusi-go/tcp/gen"
	"github.com/zusi/zusi-go/tcp/message"
)

var rootPath = flag.String("root", "", "root path of repository")

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	flag.Parse()

	if err := mainErr(); err != nil {
		slog.With("err", err).Error("running tcpgen failed")
		os.Exit(1)
	}
}

func mainErr() error {
	root := message.Message{}

	res := gen.FindStructsToReflect(root)

	var messages []gen.Message
	for _, v := range res {
		msg, err := gen.Reflect(reflect.New(v.Elem()).Interface())
		if err != nil {
			return fmt.Errorf("reflecting struct '%s' failed: %w", v.Name(), err)
		}

		messages = append(messages, *msg)
	}

	sort.Sort(Messages(messages))

	err := gen.WriteFile(messages, *rootPath)
	if err != nil {
		return err
	}

	return nil
}

type Messages []gen.Message

func (m Messages) Len() int {
	return len(m)
}

func (m Messages) Less(i, j int) bool {
	if m[i].PkgPath == m[j].PkgPath {
		return m[i].TypeName < m[j].TypeName
	} else {
		return m[i].PkgPath < m[j].PkgPath
	}
}

func (m Messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
