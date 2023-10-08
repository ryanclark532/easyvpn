package vpn_dtos

import "time"

type ServerConnection struct {
	CommonName     string
	Address        string
	BytesRec       string
	BytesSent      string
	ConnectedSince time.Time
}
