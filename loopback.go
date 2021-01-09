package main

import (
    "bufio"
    "fmt"
	"os"
	"rawmidi/midi"
)

// read one byte at a time from midi input
// echo that byte to midi output
func DoLoopBack() {

	for {
		result := midi.Read_one_byte();
		if (result > 0) {
			b := midi.Get_one_byte()
			midi.Write_byte(b)
		}
	}
	
}

func main() {
	
	var status  int	

	status = midi.Open_midi_input()
	if status != 0 {
		panic(fmt.Sprintf("error opening MIDI input (hw:1,0,0); open_midi_input() returned: %d\n", status))
	}


	status = midi.Open_midi_output()
	if status != 0 {
		panic(fmt.Sprintf("error opening MIDI output (hw:1,0,0); open_midi_input() returned: %d\n", status))
	}

	go DoLoopBack()

	fmt.Print("Echoing... Press enter to exit...\n")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	midi.Close_midi()

}