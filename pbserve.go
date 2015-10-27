package main

import(
  "log"
  "fmt"
  "net"
  "os"
  "os/user"
  "os/exec"
  "os/signal"
  "golang.org/x/crypto/ssh"
  "golang.org/x/crypto/ssh/agent"
  "syscall"
  "flag"
)

func openSSHConnection(user string, hostname string, port int) (*ssh.Client, error) {
  agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
  if err != nil {
    return nil, err
  }

  agent := agent.NewClient(agentConn)
  auths := []ssh.AuthMethod{
    ssh.PublicKeysCallback(agent.Signers),
  }

  config := &ssh.ClientConfig{
    User: user,
    Auth: auths,
  }

  address := fmt.Sprintf("%s:%d", hostname, port)
  ssh, err := ssh.Dial("tcp", address, config)
  if err != nil {
    return nil, err
  }

  return ssh, nil
}

type PbConnection struct {
  CommandName string
  RemotePort int
}

func (c PbConnection) ListenAndServe(ssh *ssh.Client) {
  address := fmt.Sprintf("localhost:%d", c.RemotePort)
  remote, err := ssh.Listen("tcp", address)
  if err != nil {
    log.Fatalf("Unable to bind %s to %s: %v", c.CommandName, address, err)
  }
  log.Printf("%s waiting for connections on %s", c.CommandName, address)

  for {
    client, err := remote.Accept()
    if err != nil {
      log.Printf("Error %v accepting connection for %s on %s", err, c.CommandName, address)
    }
    log.Printf("%s connection", c.CommandName)

    cmd := exec.Command(c.CommandName)
    cmd.Stdin = client
    cmd.Stdout = client
    cmd.Run()

    client.Close()
  }
}

func currentUser() string {
  user, err := user.Current()
  if err != nil {
    log.Fatalf("unable to determine current user: %v", err)
  }

  return user.Username
}

func main() {
  hostname := flag.String("host", "", "host to expose paste board to")
  copyPort := flag.Int("copy-port", 2224, "remote pbcopy port")
  pastePort := flag.Int("paste-port", 2225, "remote pbpaste port")
  remoteUser := flag.String("user", currentUser(), "remote user")

  flag.Parse()

  if *hostname == "" {
    log.Fatalln("hostname is required")
  }

  ssh, err := openSSHConnection(*remoteUser, *hostname, 22)
  if err != nil {
    log.Fatalf("Unable to open ssh connection to %s: %v", *hostname, err)
  }
  log.Println("ssh connection opened to ", *hostname)

  pbcopy := PbConnection{
    CommandName: "pbcopy",
    RemotePort: *copyPort,
  }

  pbpaste := PbConnection{
    CommandName: "pbpaste",
    RemotePort: *pastePort,
  }

  go pbcopy.ListenAndServe(ssh)
  go pbpaste.ListenAndServe(ssh)

  sigs := make(chan os.Signal, 1)

  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  <-sigs
  fmt.Println("exiting")
  ssh.Close()
}
