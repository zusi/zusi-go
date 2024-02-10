package st3

import (
	"encoding/xml"
	"fmt"

	"github.com/zusi/zusi-go/files"
)

type Reader struct {
	file File
}

func Read(root, filename string) (*Reader, error) {
	bts, err := files.Read(root, filename)
	if err != nil {
		return nil, fmt.Errorf("error loading timetable: %w", err)
	}

	var file File
	err = xml.Unmarshal(bts, &file)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling timetable: %w", err)
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
