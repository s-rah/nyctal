# Nyctal Wayland Compositor 

This directory contains an example Nyctal compositor that makes use of direct reading of linux DRI and event devices. This means that
to run sucessfully this application needs to be run as a user with `input` and `video` group permissions -this is probably something you don't want to do.

As such, take this as in illustrative example of how to build a display server using Nyctal. In reality you will probably want to integate with something like libseat in order for an unpriviledged user to run the compositor such that it safely access input and output devices (while disallowing that right to other user processes).



