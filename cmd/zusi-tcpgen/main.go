package main

import (
	"flag"
	"reflect"
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/zusi/zusi-go/tcp/gen"
	"github.com/zusi/zusi-go/tcp/message"
)

var rootPath = flag.String("root", "", "root path of repository")

func main() {
	flag.Parse()

	root := message.Message{}

	res := gen.FindStructsToReflect(root)

	var messages []gen.Message
	for _, v := range res {
		msg, err := gen.Reflect(reflect.New(v.Elem()).Interface())
		if err != nil {
			log.WithError(err).WithField("f", v).Error()
		}

		messages = append(messages, *msg)
	}

	sort.Sort(Messages(messages))

	err := gen.WriteFile(messages, *rootPath)
	if err != nil {
		log.WithError(err).Error()
	}
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
