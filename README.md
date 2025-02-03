# Bambu Lab Bridge

Single-purpose tool that informs your slicer (Bambu Studio or Orca Slicer) on how to discover your Bambu Lab 3D printer on another subnet.

## Execution

The tool requires 3 pieces of information to run:

1.  Device model.
1.  IP address of the 3D printer.
1.  Serial number of the 3D printer.

Once you have these, you can pass it into the tool with either Golang or Docker

### Run with Docker

Depending on your networking & Docker configuration, you may need to input the IP address of the machine where you're running your slicer.

```shell
$ docker run --rm -ti slikshooz/bambulab-bridge:latest \
    -m your_model \
    -p your_printer_ip \
    -n your_printer_serial_number \
    -s yourS-licer_ip
```


### Run with Golang

Shortest method to run the binary with Go is to pass.

```shell
$ go run ./cmd/... -m your_model -p your_printer_ip -n your_printer_serial_number
```