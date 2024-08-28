# Nyctal-X11 - Run Nyctal in an X11 Window

This directory contains code that demonstrates how to run Nyctal in an X11 window, allowing you to run Wayland applications inside a typical
X session.

Nyctal-X11 depends on minifb for creating and managing the X window, further it requires a patched version that:

 1) hides the X11 cursor within the window (so that Wayland cursors can be seen and used properly) and 
 2) sends linux scancodes in keyboard events instead of xkb mapped codes (so that keyboard handling functions as-intended).

In order to build libminifb.a with the patch you will need to do the following:

```
    git clone https://github.com/emoon/minifb
    cd minifb
    git apply ../nyctal-minifb-fixes.patch
    mkdir build
    cd build
    cmake .. -DUSE_OPENGL_API=ON -DUSE_WAYLAND_API=OFF
    make -j
```

You should then be able to run `go build nyctal-x11.go` from the nyctal-x11 directory