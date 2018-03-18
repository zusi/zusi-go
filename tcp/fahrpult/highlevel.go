package fahrpult

import (
	"github.com/zusi/zusi-go/tcp/message/fahrpult"
	"github.com/zusi/zusi-go/tcp/message/fst"
)

func (c *Client) SubscribeData(fstIds ...fst.FuehrerstandId) error {
	var ids = make([]uint16, len(fstIds))

	for i, id := range fstIds {
		ids[i] = uint16(id)
	}

	msg := &fahrpult.NeededDataMessage{
		Fuehrerstandsanzeigen: &fahrpult.Fuehrerstandsanzeigen{
			Anzeigen: ids,
		}}

	err := c.SendFp(fahrpult.FahrpultMessage{NeededData: msg})

	return err
}
