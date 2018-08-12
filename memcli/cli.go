package memcli

import (
	"net"
	"errors"
	"fmt"
	"time"
	"bufio"
)

const (
	DefaultTimeout = 10000 * time.Millisecond
);

type Memcli struct {
	/*network access point*/
	address net.Addr
}

func (this *Memcli) Network() (string) {
	return this.address.Network()
}

func (this *Memcli) String() (string) {
	return this.address.String()
}

func (this *Memcli) Get(key string) ([]byte, error) {
	conn, e := net.DialTimeout(this.Network(), this.String(), DefaultTimeout)
	readWriter := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	if e != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to network: %s server: %s", this.Network(), this.String()))
	}
	conn.SetDeadline(time.Now().Add(DefaultTimeout))
	reader := readWriter.Reader
	line, err := reader.ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	fmt.Println(line)
	return nil, nil
}

func New(server string) (*Memcli, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to memcache on: %s", server))
	}
	return &Memcli{address:tcpAddr}, nil
}