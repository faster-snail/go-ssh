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

//此例是从https://godoc.org/golang.org/x/crypto/ssh 拷贝修改，将句柄创建封装
func New(addr string,username string,password string) mycli {
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
        //需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		//如果按照官方文档操作出现问题，可以如下操作 		
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
//对官方案例的命令执行过程封装成shell方法，此处为了使用cli句柄方法，在上面将ssh.Client类型封装为结构体
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
	//（1）初始化一个cli句柄
	//（2）执行shell
	//（3）打印输出
	cli := New("172.0.0.42:22","root","ybcsp.!QAZ")
	str := cli.Myshell("ls -l /")
	fmt.Println(str)

}
