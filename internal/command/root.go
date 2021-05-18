package command

import (
	"io/ioutil"
	"os"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
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

	request := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, request); err != nil {
		panic(err)
	}

	generatorService := generator.GetService()
	response := generatorService.GenerateCode(request)

	marshalled, err := proto.Marshal(response)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
