package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_iphone.h>
import "C"

/*
// Multitasking on iOS is different than on other platforms. When an
// application receives an ALLEGRO_DISPLAY_SWITCH_OUT or ALLEGRO_DISPLAY_CLOSE
// event on a multitasking-capable device, it should cease all activity and do
// nothing but check for an ALLEGRO_DISPLAY_SWITCH_IN event. To let the iPhone
// driver know that you've ceased all activity, call this function. You should
// call this function very soon after receiving the event telling you it's time
// to switch out (within a couple milliseconds). Certain operations, if done,
// will crash the program after this call, most notably any function which uses
// OpenGL. This function is needed because the "switch out" handler on iPhone
// can't return until these operations have stopped, or a crash as described
// before can happen.
func IPhoneProgramHasHalted() {
    C.al_iphone_program_has_halted()
}

// Original iPhones and iPod Touches had a screen resolution of 320x480 (in
// Portrait mode). When the iPhone 4 and iPod Touch 4th generation devices came
// out, they were backwards compatible with all old iPhone apps. This means
// that they assume a 320x480 screen resolution by default, while they actually
// have a 640x960 pixel screen (exactly 2x on each dimension). An API was added
// to allow access to the full (or in fact any fraction of the) resolution of
// the new devices. This function is normally not needed, as in the case when
// you want a scale of 2.0 for "retina display" resolution (640x960). In that
// case you would just call al_create_display with the larger width and height
// parameters. It is not limited to 2.0 scaling factors however. You can use
// 1.5 or 0.5 or other values in between, however if it's not an exact multiple
// of the original iPhone resolution, linear filtering will be applied to the
// final image.
func IPhoneOverrideScreenScale(scale float32) {
    C.al_iphone_override_screen_scale(C.float(scale))
}
*/
