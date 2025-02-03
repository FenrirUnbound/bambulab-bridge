package payload

import (
	"fmt"
	"strings"
	"time"
)

const (
	DEFAULT_DEVICE_BIND    = "free"
	DEFAULT_DEVICE_CONNECT = "lan"
	DEFAULT_DEVICE_SIGNAL  = "-44"
)

type Param struct {
	Timestamp     string
	PrinterIP     string
	SerialNumber  string
	DeviceModel   string
	DeviceName    string
	DeviceSignal  string
	DeviceConnect string
	DeviceBind    string
}

func New(data *Param) ([]byte, error) {
	timestamp, err := time.Parse(time.RFC3339, data.Timestamp)
	if err != nil {
		return nil, err
	}

	deviceSignal := DEFAULT_DEVICE_SIGNAL
	if data.DeviceSignal != "" {
		deviceSignal = data.DeviceSignal
	}

	deviceConnect := DEFAULT_DEVICE_CONNECT
	if data.DeviceConnect != "" {
		deviceConnect = data.DeviceConnect
	}

	deviceBind := DEFAULT_DEVICE_BIND
	if data.DeviceBind != "" {
		deviceBind = data.DeviceBind
	}

	parts := []string{
		"HTTP/1.1 200 OK",
		"Server: Buildroot/2018.02-rc3 UPnP/1.0 ssdpd/1.8",
		fmt.Sprintf("Date: %s", timestamp),
		fmt.Sprintf("Location: %s", data.PrinterIP),
		"ST: urn:bambulab-com:device:3dprinter:1",
		"EXT:",
		fmt.Sprintf("USN: %s", data.SerialNumber),
		"Cache-Control: max-age=1800",
		fmt.Sprintf("DevModel.bambu.com: %s", data.DeviceModel),
		fmt.Sprintf("DevName.bambu.com: %s", data.DeviceName),
		fmt.Sprintf("DevSignal.bambu.com: %s", deviceSignal),
		fmt.Sprintf("DevConnect.bambu.com: %s", deviceConnect),
		fmt.Sprintf("DevBind.bambu.com: %s", deviceBind),
		"\r\n",
	}

	result := strings.Join(parts, "\r\n")

	return []byte(result), nil
}
