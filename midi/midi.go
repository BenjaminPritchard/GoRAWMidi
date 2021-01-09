package midi
/*
extern int c_open_midi_input();
extern int c_open_midi_output();
extern int c_read_one_byte();
extern unsigned char c_get_one_byte();
extern int c_write_byte(unsigned char b);
extern void c_close_midi();
*/
// #cgo LDFLAGS: -lasound
import "C"

// wrapper functions to perform type conversions
func Open_midi_input() int {
    return int(C.c_open_midi_input())
}

func Open_midi_output() int {
    return int(C.c_open_midi_output())
}

func Read_one_byte() int {
    return int(C.c_read_one_byte())
}

func Get_one_byte() byte {
	return byte(C.c_get_one_byte())
}

func  Write_byte(b byte) int {
	return int(C.c_write_byte(C.uchar(b)))
}

func Close_midi() {
    C.c_close_midi()
}