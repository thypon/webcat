WebCat
======

`webcat` is command line utility that permits to interface simple unix commands to WebSocket channels. WebSocket API can be consumed easily.

-@thypon

About
-----

Upon startup, `webcat` connects to a websocket endpoint.

Upon connection, it will redirect server WebSocket messages to `STDOUT` separated by a newline `\n`. Meanwhile permits you to redirect `STDIN` lines to server with a WebSocket message.