package handlers

import (
	"encoding/json"
	"net"
	"net/http"
)

type client struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func ClientRequestIp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	clientAddr, clientPort, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	clientIP := net.ParseIP(clientAddr)

	c := client{
		IP:   clientIP.String(),
		Port: clientPort,
	}

	j, _ := json.Marshal(c)

	res.Write(j)
}
