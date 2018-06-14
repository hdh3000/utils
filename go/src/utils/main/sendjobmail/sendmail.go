package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path"
	"reflect"
	"utils/sendjobmail"

	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
)

var homeDir = os.Getenv("HOME")
var configPath = path.Join(homeDir, ".config/sendjobmail/config.json")

type sendMailConfig struct {
	// No pointers, makes it easy to write out an empty file with fields.
	TemplateDir string
	SenderInfo  sendjobmail.BasicInfo
}

func main() {
	ctx := context.Background()
	flag.Parse()
	config, err := ParseConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Now parse the arguments body of the
	tmplName, toInfo, err := ParseArgs(flag.Args())
	if err != nil {
		log.Fatalf("%s", err)
	}

	tmplArgs := &sendjobmail.TmplArgs{
		To:   toInfo,
		From: &config.SenderInfo,
	}

	body, err := sendjobmail.ParseTmpl(tmplName, config.TemplateDir, tmplArgs)
	if err != nil {
		log.Fatalf("%s", err)
	}

	b, _ := json.MarshalIndent(toInfo, "", " ")

	fmt.Println("I am going to send this, you sure you want to?")
	fmt.Println("--------------------------------")
	fmt.Println("")
	fmt.Printf("%s\n", string(b))
	fmt.Println("")
	fmt.Println("--------------------------------")
	fmt.Println(body)
	fmt.Println("--------------------------------")
	fmt.Println("")

	gmailSvc, err := gmail.New(sendjobmail.GetClient())
	if err != nil {
		log.Fatalf("failed to get gmail client %s", err)
	}

	to := mail.Address{Name: "", Address: toInfo.Email}
	from := mail.Address{Name: fmt.Sprintf("%s %s", config.SenderInfo.First, config.SenderInfo.Last), Address: config.SenderInfo.Email}

	if err := sendjobmail.SendEmail(ctx, gmailSvc, &from, &to, "TEST", body); err != nil {
		log.Fatalf("failed to send email %s", err)
	}

}

func ParseArgs(args []string) (string, *sendjobmail.BasicInfo, error) {
	// Expect the args to be in a strict order (the order they are in the struct), with the name of the template first.
	expectedArgs := reflect.TypeOf(sendjobmail.BasicInfo{}).NumField() + 1
	if len(args) != expectedArgs {
		return "", nil, fmt.Errorf("expected %v args, but got %v", expectedArgs, len(args))
	}

	return args[0], &sendjobmail.BasicInfo{
		Email:       args[1],
		First:       args[2],
		Last:        args[3],
		CompanyName: args[4],
		Title:       args[5],
	}, nil
}

func ParseConfig() (*sendMailConfig, error) {
	f, err := os.Open(configPath)
	if os.IsNotExist(err) {
		os.Mkdir(path.Dir(configPath), 0700)
		// If it doesn't exist
		b, _ := json.MarshalIndent(sendMailConfig{}, "", " ")
		f, err := os.Create(configPath)
		if err != nil {
			return nil, fmt.Errorf("err creating config file %s", err)
		}
		defer f.Close()
		f.Write(b)

		return nil, fmt.Errorf("your config file didn't exist, I made an empty one at %s, please fill it out", configPath)

	} else if err != nil {
		return nil, fmt.Errorf("err opening config file, please go and fill it out (%s) using the sendjobmail config json format\n %s", configPath, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	var config sendMailConfig

	return &config, decoder.Decode(&config)

}
