# Nyctal - a tiny, zero dependency, Wayland Compositor 

(or about 3000 lines of code you would have to write anyway....)

## What and Why

Nyctal is a **prototype** wayland compositor, written in Go with zero[^1] dependencies. If you have only a Go compiler you
should be able to compile Nyctal and direct applications to use it as a display server.

([^1]: Ok, *technically* there is a dependency on `golang.org/x/sys`, which isn't part of the go standard library, but it's close enough that I'm going to claim mission success.)

 Nyctal is **not a drop-in replacement** for any of the well-engineered, feature-full display servers out there. You will probably want to use something based on  [wl_roots](https://gitlab.freedesktop.org/wlroots/wlroots).
 
 Nyctal was written to be used as a base for building your own display server, probably to be deployed on custom kernels or in very-opinionated setups.
 The base code of Nyctal is not aware of systemd, or dbus, or seats or another other modern linux "standard" services.

 To put it plainly: You probably don't want to use Nyctal, and even if you do you almost certainly will want to fork the code and build your own window management logic.

 I built Nyctal because I like building operating systems as a hobby and have been wanting a small, easily maintainable display server for those projects and has the potential to be compatible with more mainstream applications.

## Trying out Nyctal

There are currently two example applications that can be used to try out Nyctal:

1. [cmd/nyctal-x11](cmd/nyctal-x11) - hosts a nyctal compositor in an X11 window (depends on building minifb, see instructions in the linked folder).
2. [cmd/nyctal-dri](cmd/nyctal-dri) - requires direct access to `input` and `video` devices (see instructions in the linked folder).

Ensure that any apps are run in an environment setting `XDG_RUNTIME_DIR=/tmp/nyctal/` and `WAYLAND_DISPLAY=nyctal-0`


### Applications that Work with Nyctal

Nyctal can currently interact and render a wide range of applications, some that I've tested include:

- Flutter-based apps
- Zathura (for pdf rendering)
- Basic gtk apps (like xfce-terminal)

Nyctal will currently fail to interact with anything that expects to be able to use the dmabuf extension (so flags like `LIBGL_ALWAYS_SOFTWARE=1` and `__GLX_VENDOR_LIBRARY_NAME=mesa` are needed when running applications to force software rendering)

Nyctal also doesn't support extensions that most video apps (like mpv) need.
 

## Does it support <protocol>?

Nyctal aims to supports the bare-minimum of wayland needed to run most basic applications. In practice that means that Nyctal currently supports:

- The **Core Wayland Protocol** (Except `wl_shell` which is deprecated)
    - [X] wl_compositor
    - [ ] wl_subcompositor (partial)
        - [ ] wl_subsurface 
    - [X] wl_shm 
    - [X] wl_shmpool
    - [X] wl_region (partial)
    - [X] wl_seat
        - [X] wl_pointer
        - [X] wl_keyboard
            - [X] No keyboard map (raw scancodes)
            - [ ] xkb keyboard maps
        - [ ] wl_touch
    - [ ] wl_data_device_manager (partial)
        - [ ] wl_data_device (partial)
        - [ ] wl_data_source (partial, offers are supported which means in-app copying/pasing works, but cross-app sharing currently does not)
- The **XDG Shell Protocol** Extension (Necessary because `wl_shell` is deprecated)
    - [X] xdg_wm_base
    - [X] xdg_surface
        - [X] xdg_toplevel (requrest like set_title / set_app_id / set_min_size are currently ignored)
        - [X] xdg_popup
            - [X] xdg_positioner
 
Given time, Nyctal also aims to support:

- The **Viewporter** extension 
- The **linux dmabuf** extenion
- The **Presentation Time** extension

It is incredibly unlikely that Nyctal will ever have X11 protocol support, or support for the more niche wayland protocols (e.g. those designed to directly integrate with a given compositor) - but if you think you have a way of introducing such support in a way that aligns with [the goals of the project](CONTRIBUTING.md) then I await your pull request!

### What about keymappings?

Nyctal only supports sending linux scancodes to applications, with all the inherent limitations of that approach. PRs that introduce support for xkbcommon keymaps (while aligning with the goals of the project) are welcome.
