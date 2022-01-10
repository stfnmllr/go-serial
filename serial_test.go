//
// Copyright 2014-2021 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package serial_test

import (
	"testing"

	"go.bug.st/serial"
)

// TestResetBuffer tests resetting the serial port's input,- and output buffer.
func TestResetBuffer(t *testing.T) {

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		t.Fatal(err)
	}
	if len(ports) == 0 {
		t.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		t.Logf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		t.Fatal(err)
	}
	defer port.Close()

	// Reset input buffer
	if err := port.ResetInputBuffer(); err != nil {
		t.Fatal(err)
	}
	// reset output buffer
	if err := port.ResetOutputBuffer(); err != nil {
		t.Fatal(err)
	}
}
