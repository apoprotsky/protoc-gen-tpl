package command

import (
	"io/ioutil"
	"os"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/golang"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func rootCommand(cmd *cobra.Command, args []string) {
	commandService := GetService()
	if commandService.NeedHelp() {
		cmd.Help()
		return
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	request := &plugin.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, request); err != nil {
		panic(err)
	}

	golangService := golang.GetService()
	responce := golangService.GenerateCode(request)

	marshalled, err := proto.Marshal(responce)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
