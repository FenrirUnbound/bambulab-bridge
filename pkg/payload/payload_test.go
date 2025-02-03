package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	timestamp := "2025-02-18T18:20:00Z"
	param := &Param{
		Timestamp:    timestamp,
		PrinterIP:    "192.168.4.20",
		SerialNumber: "FakeSerialNumber",
		DeviceModel:  "FakeDeviceModel",
		DeviceName:   "FakeDeviceName",
		DeviceSignal: "-88",
	}

	result, err := New(param)
	assert.NoError(t, err)

	expected := "HTTP/1.1 200 OK\r\nServer: Buildroot/2018.02-rc3 UPnP/1.0 ssdpd/1.8\r\nDate: 2025-02-18 18:20:00 +0000 UTC\r\nLocation: 192.168.4.20\r\nST: urn:bambulab-com:device:3dprinter:1\r\nEXT:\r\nUSN: FakeSerialNumber\r\nCache-Control: max-age=1800\r\nDevModel.bambu.com: FakeDeviceModel\r\nDevName.bambu.com: FakeDeviceName\r\nDevSignal.bambu.com: -88\r\nDevConnect.bambu.com: lan\r\nDevBind.bambu.com: free\r\n\r\n"

	assert.Equal(t, []byte(expected), result)
}
