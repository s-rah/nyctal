# Contributing to Nyctal

The following is a set of guidelines for contrbuting to Nyctal.

## Tenets

- Zero Dependencies: One of the primary goals of Nyctal is to place a hard limit on the number of external dependencies. That number is zero. Prefer to delegate non-trivial integration (e.g. systemd / dbus / eudev) to downstream forks.

- Limit Code Size: Nyctal should be understandable by a single developer, and a single developer should feel comfortable forking Nyctal and using it as the base for thier own display server.

- Broadest compatibility set: the core Nyctal codebase is not the place to introduce desktop-environment specific wayland extensions. With few exceptions, the only protocols supported by Nyctal directly should be those designated "core" and "stable".


## Code Paths Likely to Change

It is highly likely that the entire rendering pipeline from surface caching to output rendering will change in the neat future (to support the linux dmabuf extension) and before embarking on changes to related code you should consult with Sarah.

It is likely that libseat support will be added in the future. Code in the `drm` and `evdev` directories are experimental, and for example use only and while unlikely to change, should not be used as the basis for future support.


# A Brief Overview of the Code


- `cmd` - a set of example standalone executables 
    - `cmd/nyctal-dri` - a compositor directly accessing linux direct rendering interface and event devices
    - `cmd/nyctal-x11` - a compositor that outputs to an X11 window. Useful for testing in X11 desktop environments.
- `model` - shared interfaces and structures related to compositing and workspaces e.g. keyboard, image formats, clients
- `specs` - Cached reference documents for various wayland protocols (not used by any other part of the code)
- `utils` - Small data structures used throughout the code e.g. queue, stack and logging
- `wayland` - All wayland protocol code including code for handling unix domain sockets, packet parsing and sending routing.
- `workspace` - A very basic tiling wayland compositor

