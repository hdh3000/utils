package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"utils/sendjobmail"
)

var homeDir = os.Getenv("~")
var configPath = path.Join(homeDir, "/Users/hdh/.config/sendjobmail/config.json")

type sendMailConfig struct {
	// No pointers, makes it easy to write out an empty file with fields.
	TemplateDir string
	SenderInfo  sendjobmail.BasicInfo
}

func main() {
	flag.Parse()
	config, err := ParseConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Now parse the arguments out of the
	tmplName, toInfo, err := ParseArgs(flag.Args())
	if err != nil {
		log.Fatalf("%s", err)
	}

	tmplArgs := &sendjobmail.TmplArgs{
		To:   toInfo,
		From: &config.SenderInfo,
	}

	out, err := sendjobmail.ParseTmpl(tmplName, config.TemplateDir, tmplArgs)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(out)

}

func ParseArgs(args []string) (string, *sendjobmail.BasicInfo, error) {
	// Expect the args to be in a strict order (the order they are in the struct), with the name of the template first.
	expectedArgs := reflect.TypeOf(sendjobmail.BasicInfo{}).NumField() + 1
	if len(args) != expectedArgs {
		return "", nil, fmt.Errorf("expected %v args, but got %v", expectedArgs, len(args))
	}

	return args[0], &sendjobmail.BasicInfo{
		First:       args[1],
		Last:        args[2],
		CompanyName: args[3],
		Title:       args[4],
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
