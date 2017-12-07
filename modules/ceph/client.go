package ceph

import (
	"github.com/ceph/go-ceph/rados"
)

type Ceph struct {
	Conn rados.Conn
}

func (c *Ceph) Client() {
	conn, err := rados.NewCon()
	if err != nil {
		return nil, err
	}
	c.Conn = conn
}
