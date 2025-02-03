package command

import (
	"fmt"
	"time"

	"github.com/fenrirunbound/bambulab-bridge/pkg/client"
	"github.com/fenrirunbound/bambulab-bridge/pkg/device"
	"github.com/fenrirunbound/bambulab-bridge/pkg/payload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DEFAULT_DEVICE_NAME      = "my_printer"
	DEFAULT_INTERVAL_SECONDS = 5
	DEFAULT_SLICER_IP        = "127.0.0.1"

	flagUsageDeviceModel = `Device model (required).
Can be x1c, x1, p1p, p1s, x1e, a1m (A1 Mini), a1.`
)

func NewMainCommand() *cobra.Command {
	debugMode := false
	deviceModelLabel := ""
	deviceName := DEFAULT_DEVICE_NAME
	intervalSeconds := DEFAULT_INTERVAL_SECONDS
	serialNumber := ""
	slicerIP := DEFAULT_SLICER_IP
	printerIP := ""

	rootCmd := &cobra.Command{
		Use: "bambulab-bridge -m your_printer_model -p your_printer_ip -n your_printer_serial_number",
		RunE: func(cmd *cobra.Command, args []string) error {
			deviceLabel := cmd.Flag("device-model").Value.String()
			model, err := device.FromLabel(deviceLabel)
			if err != nil {
				return err
			}

			deviceName := cmd.Flag("device-name").Value.String()
			if deviceName == "" {
				deviceName = deviceLabel
			}

			udpClient, err := client.New(slicerIP)
			if err != nil {
				return err
			}
			defer udpClient.Close()

			for {
				now := time.Now().Format(time.RFC3339)

				p := &payload.Param{
					Timestamp:    now,
					PrinterIP:    printerIP,
					SerialNumber: serialNumber,
					DeviceModel:  model,
					DeviceName:   deviceName,
				}
				buf, err := payload.New(p)
				if err != nil {
					return err
				}

				_, err = udpClient.Write(buf)
				if err != nil {
					return err
				}

				if debugMode {
					fmt.Printf("%s", buf)
				}
				time.Sleep(time.Duration(intervalSeconds) * time.Second)
			}
		},
	}

	rootCmd.Flags().BoolVar(&debugMode, "debug", false, "Flag to enable debug mode.")
	rootCmd.Flags().StringVarP(&deviceModelLabel, "device-model", "m", "", flagUsageDeviceModel)
	rootCmd.Flags().StringVarP(&serialNumber, "serial-number", "n", "", "Device serial number (required).")
	rootCmd.Flags().StringVarP(&printerIP, "printer-ip", "p", "", "IP address of the printer (required).")
	rootCmd.Flags().StringVarP(&slicerIP, "slicer-ip", "s", DEFAULT_SLICER_IP, "IP address of the slicer machine.")
	rootCmd.Flags().StringVarP(&deviceName, "device-name", "d", "", "Device name. Defaults to device model.")
	rootCmd.Flags().IntVarP(&intervalSeconds, "interval", "i", 5, "Time interval in seconds between pings.")

	_ = viper.BindPFlag("serialNumber", rootCmd.Flags().Lookup("serial-number"))
	_ = viper.BindPFlag("slicer-ip", rootCmd.Flags().Lookup("slicer-ip"))
	_ = viper.BindPFlag("printer-ip", rootCmd.Flags().Lookup("printer-ip"))
	_ = viper.BindPFlag("device-model", rootCmd.Flags().Lookup("device-model"))

	_ = rootCmd.MarkFlagRequired("device-model")
	_ = rootCmd.MarkFlagRequired("serial-number")
	_ = rootCmd.MarkFlagRequired("printer-ip")

	return rootCmd
}
