#include <alsa/asoundlib.h>     /* Interface to the ALSA system */

snd_rawmidi_t* midiin = NULL;
snd_rawmidi_t* midiout = NULL;

unsigned char buffer[1];        // Storage for input buffer received

int c_open_midi_input() {
    int status;
    int mode = SND_RAWMIDI_NONBLOCK;
    const char* portname = "hw:1,0,0";
    return snd_rawmidi_open(&midiin, NULL, portname, mode);
}

int c_open_midi_output() {
    int status;
    int mode = SND_RAWMIDI_SYNC;
    const char* portname = "hw:1,0,0";
    return snd_rawmidi_open(NULL, &midiout, portname, mode);
}

// try to eat the keep alive messages...
int c_read_one_byte() {
    int retval = 254;
    while (retval == 254) {
        retval = snd_rawmidi_read(midiin, buffer, 1);
    }
    return retval;
}

unsigned char c_get_one_byte() {
    return buffer[0];  
}

int c_write_byte(unsigned char b) {

   char tmpbuffer[0];
   tmpbuffer[0] = b;

   return snd_rawmidi_write(midiout, tmpbuffer, 1);
}

void c_close_midi() {
    snd_rawmidi_close(midiin);
}