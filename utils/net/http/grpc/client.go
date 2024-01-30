package grpci

import (
	"crypto/tls"
	"github.com/hopeio/tiga/utils/errors/multierr"
	"github.com/hopeio/tiga/utils/net/http/grpc/stats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

var ClientConns = make(clientConns)

func ClientConnsClose() error {
	if ClientConns != nil {
		return ClientConns.Close()
	}
	return nil
}

type clientConns map[string]*grpc.ClientConn

func (cs clientConns) Close() error {
	var multiErr multierr.MultiError
	for _, conn := range cs {
		err := conn.Close()
		if err != nil {
			multiErr.Append(err)
		}
	}
	if multiErr.HasErrors() {
		return &multiErr
	}
	return nil
}

func GetDefaultClient(addr string) (*grpc.ClientConn, error) {
	if conn, ok := ClientConns[addr]; ok {
		return conn, nil
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(&stats.ClientHandler{}))
	if err != nil {
		return nil, err
	}

	ClientConns[addr] = conn
	return conn, nil
}

func GetTlsClient(addr string) (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: strings.Split(addr, ":")[0], InsecureSkipVerify: true})),
		grpc.WithStatsHandler(&stats.ClientHandler{}))
	if err != nil {
		return nil, err
	}
	if oldConn, ok := ClientConns[addr]; ok {
		oldConn.Close()
	}
	ClientConns[addr] = conn
	return conn, nil
}
