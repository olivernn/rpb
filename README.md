# RPB

Remote `pbcopy` and `pbpaste` commands.

Simulates the OS X [`pbcopy`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbcopy.1.html) and [`pbpaste`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbpaste.1.html) commands on a remote host.

## Installation

### Remote Host

Just put the `pbcopy` and `pbpaste` commands on your path.

### Local (OS X) Host

The provided `pbserve` command will connect to your remote host and forward connections to the local `pbcopy` and `pbserve` builtins.

    $ pbserve --host remote.host.name

By default the current user will be used to access the remote machine, this can be changed using the `--user` switch

    $ pbserve --host remote.host.name --user bob
