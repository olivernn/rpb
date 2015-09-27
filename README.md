# RPB

Remote `pbcopy` and `pbpaste` commands.

Simulates the OS X [`pbcopy`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbcopy.1.html) and [`pbpaste`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/pbpaste.1.html) commands on a remote host.

Requires that the OS X commands have been bound to listen on a port that is forwarded to the remote server, `launchd` configuration for both `pbcopy` and `pbpaste` are provided in this repository.

## Installation

### Remote Host

Just put the two commands on your path.

### Local (OS X) Host

Copy the launchd config files `local.pbcopy.plist` and `local.pbpaste.plist` into `~/Library/LaunchAgents` and then enable both of these services.

```
$ cp local.*.plist ~/Library/LaunchAgents
```

Load both of these services using `launchctl`

```
$ launchctl load ~/Libary/LaunchAgents/local.pbcopy.plist
$ launchctl load ~/Libary/LaunchAgents/local.pbpaste.plist
```

You can test that this is working by executing the following command, you should see the contents of your clipboard:

```
$ telnet localhost 2225
```

To forward the relevant ports to the remote server automatically when sshing to that host you can add the following to your `.ssh/config`:

```
RemoteForward 2224 localhost:2224
RemoteForward 2225 localhost:2225
```

