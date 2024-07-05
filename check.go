package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := 5 * time.Second
	start := time.Now()

	// DNS Resolution
	ipRecords, err := net.LookupIP(destination)
	if err != nil {
		return fmt.Sprintf("[ERROR] Failed to resolve DNS for %v,\n Error: %v", destination, err)
	}
	ipAddresses := ""
	for _, ip := range ipRecords {
		ipAddresses += ip.String() + " "
	}

	// TCP Connectivity
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Sprintf("[DOWN] %v is unreachable at %v,\n DNS: %v,\n Error: %v", destination, address, ipAddresses, err)
	}
	defer conn.Close()

	duration := time.Since(start)

	// HTTP Check (Optional)
	httpStatus := ""
	if port == "80" || port == "443" {
		protocol := "http"
		if port == "443" {
			protocol = "https"
		}
		httpURL := fmt.Sprintf("%s://%s", protocol, destination)
		resp, err := http.Get(httpURL)
		if err != nil {
			httpStatus = fmt.Sprintf("[ERROR] HTTP request to %v failed,\n Error: %v", httpURL, err)
		} else {
			httpStatus = fmt.Sprintf("[HTTP] %v responded with status %v", httpURL, resp.Status)
			resp.Body.Close()
		}
	}

	// Detailed Status
	status := fmt.Sprintf(
		"[UP] %v is reachable at %v,\n DNS: %v,\n Response Time: %v,\n From: %v\n To: %v\n %v",
		destination, address, ipAddresses, duration, conn.LocalAddr(), conn.RemoteAddr(), httpStatus,
	)
	return status
}
