package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

type mycli struct {
	cli *ssh.Client
}

func New(addr string,username string,password string) mycli {
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	cli := mycli{cli:client}
	return cli
}

func (client mycli)Myshell(cmd string) string {
	session, err := client.cli.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	return b.String()
}
func main() {
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	cli := New("172.0.0.42:22","root","ybcsp.!QAZ")
	str := cli.Myshell("uptime")
	fmt.Println(str)

}
