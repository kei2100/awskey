package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/pflag"
)

var (
	flag string
	help bool
)

func init() {
	pflag.StringVarP(&flag, "flag", "f", "", "flag string. e.g. if '-f=-e', then prints -e AWS_ACCESS_KEY_ID=... -e AWS_SECRET_ACCESS_KEY=...")
	pflag.BoolVarP(&help, "help", "h", false, "show this help")
}

func main() {
	pflag.Parse()
	if help {
		pflag.Usage()
		os.Exit(2)
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("failed to new session: %v", err)
	}
	cred, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("failed to get credentials: %v", err)
	}

	b := &strings.Builder{}
	writeWithFlag(b, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", cred.AccessKeyID))
	writeWithFlag(b, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", cred.SecretAccessKey))
	fmt.Print(b.String())
}

func writeWithFlag(b *strings.Builder, v string) {
	if len(v) == 0 {
		return
	}
	if b.Len() > 0 {
		b.WriteString(" ")
	}
	if len(flag) > 0 {
		b.WriteString(flag + " ")
	}
	b.WriteString(v)
}
