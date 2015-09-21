# RPB

Remote `pbcopy` and `pbpaste` commands.

Simulates the OS X [`pbcopy`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbcopy.1.html) and [`pbpaste`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbpaste.1.html) commands on a remote host.

Requires that the OS X commands have been bound to listen on a port that is forwarded to the remote server, an example `launchd` config for setting this up can be found [here](https://gist.github.com/olivernn/e298c3a5c32f33c550cb).

By default `pbcopy` writes to localhost:2224 and `pbcopy` reads from localhost:2225. The port can be configured using the `-port` switch.
