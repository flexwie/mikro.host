package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/google/uuid"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var client *hcloud.Client

func main() {
	token := os.Getenv("hcloud_token")
	client = hcloud.NewClient(hcloud.WithToken(token))

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server, key, priv := deployVM("test")
	defer cleanupKey(key)
	defer cleanupServer(server)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			case <-ticker.C:
				out, err := remoteRun("root", server.PublicNet.IPv4.IP.String(), priv, "echo 'Welcome to this world'")
				if err == nil {
					fmt.Println(out)
					return
				}
			}
		}
	}()

	wg.Wait()
}

func deployVM(subscription string) (*hcloud.Server, *hcloud.SSHKey, string) {
	pub, priv, _ := MakeSSHKeyPair()
	hkey := createCloudKey(pub)

	fmt.Println(priv)

	create, _, err := client.Server.Create(context.Background(), hcloud.ServerCreateOpts{
		Name:       uuid.New().String(),
		Image:      &hcloud.Image{Name: "ubuntu-20.04"},
		Location:   &hcloud.Location{Name: "nbg1"},
		ServerType: &hcloud.ServerType{Name: "cx11"},
		SSHKeys:    []*hcloud.SSHKey{hkey},
		Labels:     map[string]string{"subscription": subscription},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("created server: ", create.Server.Name)

	return create.Server, hkey, priv
}

func createCloudKey(publicKey string) *hcloud.SSHKey {
	fmt.Println(publicKey)

	key, _, err := client.SSHKey.Create(context.Background(), hcloud.SSHKeyCreateOpts{
		Name:      "mikro-1234",
		PublicKey: publicKey,
		Labels:    map[string]string{"subscription": "123"},
	})

	if err != nil {
		panic(err)
	}

	return key
}

func MakeSSHKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}

	// generate and write private key as PEM
	var privKeyBuf strings.Builder

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(&privKeyBuf, privateKeyPEM); err != nil {
		return "", "", err
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	var pubKeyBuf strings.Builder
	pubKeyBuf.Write(ssh.MarshalAuthorizedKey(pub))

	return pubKeyBuf.String(), privKeyBuf.String(), nil
}

func remoteRun(user string, addr string, privateKey string, cmd string) (string, error) {
	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		// https://github.com/golang/go/issues/19767
		// as clientConfig is non-permissive by default
		// you can set ssh.InsercureIgnoreHostKey to allow any host
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		//alternatively, you could use a password
		/*
		   Auth: []ssh.AuthMethod{
		       ssh.Password("PASSWORD"),
		   },
		*/
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(cmd)
	return b.String(), err
}

func cleanupServer(server *hcloud.Server) {
	client.Server.Delete(context.Background(), server)
}

func cleanupKey(key *hcloud.SSHKey) {
	client.SSHKey.Delete(context.Background(), key)
}
