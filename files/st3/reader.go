package st3

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/zusi/zusi-go/files"
)

type Reader struct {
	file File
}

func Read(root, filename string) (*Reader, error) {
	bts, err := files.Read(root, filename)
	if err != nil {
		return nil, errors.Wrap(err, "error loading timetable")
	}

	var file File
	err = xml.Unmarshal(bts, &file)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling timetable")
	}

	return &Reader{
		file: file,
	}, nil
}

func (r *Reader) GetBetriebsstellen() []string {
	btrmap := map[string]bool{}

	for _, e := range r.file.Strecke.StrElement {
		if e.InfoNormRichtung.Signal != nil {
			btrmap[e.InfoNormRichtung.Signal.NameBetriebsstelle] = true
		}
		if e.InfoGegenRichtung.Signal != nil {
			btrmap[e.InfoGegenRichtung.Signal.NameBetriebsstelle] = true
		}
	}

	var btr []string
	for name, v := range btrmap {
		if v && name != "" {
			btr = append(btr, name)
		}
	}

	return btr
}
