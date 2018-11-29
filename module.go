package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"github.com/"
)

type rightsGen struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*rightsGen) Name() string {
	return "zap"
}

func (m *rightsGen) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (m *rightsGen) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {

	modulePath := "github.com/ayushbpl10/protoc-gen-rights/example/"

	for _, f := range targets {


		name := m.Context.OutputPath(f).SetExt(".rights.go").String()
		fm := fileModel{PackageName: m.Context.PackageName(f).String(), }
		for _,im := range f.Imports() {
			fm.Imports = append(fm.Imports, im.Descriptor().Options.GetGoPackage())
		}

		fm.Imports = append(fm.Imports, modulePath+f.Descriptor().Options.GetGoPackage())


		for _,srv := range f.Services() {

			service := serviceModel{}
			service.ServiceName = srv.Name().String()
			service.PackageName = m.Context.PackageName(f).String()

			for _, rpc := range srv.Methods() {

				opt := rpc.Descriptor().GetOptions()
				option, err := proto.GetExtension(opt, scopepb.E_Scope)
				if err != nil {
					panic(err)
				}
				byteData, err := json.Marshal(option)
				if err != nil {
					panic(err)
				}
				right := scopepb.MyScopes{}
				err = json.Unmarshal(byteData, &right)
				if err != nil {
					panic(err)
				}

				//m.Log(rpc.Output().Package().ProtoName())

				rpcModel := rpcModel{RpcName: rpc.Name().UpperCamelCase().String(), Input: rpc.Input().Name().UpperCamelCase().String(), Output: rpc.Output().Name().UpperCamelCase().String(), Option: right, PackageName: m.Context.PackageName(f).String()}

				}

				fm.Services = append(fm.Services, service)
			}

			m.OverwriteGeneratorTemplateFile(
				name,
				T.Lookup("File"),
				&fm,
			)
		}

		return m.Artifacts()
}

type rpcModel struct {
	PackageName string
	RpcName     string
	Input       string
	Output      string
	Option      scopepb.MyScopes
	Resources   []Resource
}

type Resource struct {
	IsRepeated  					bool
	GetStrings   					[]map[string]bool
	ResourceStringWithCurlyBraces 	string
	ResourceStringWithFormatter     string
	ForLoop     					[]ForLoop
}

type ForLoop struct {
	RangeKey 	   string
	ValueKey string
	Level      int
}

type serviceModel struct {
	ServiceName   string
	PackageName   string
	Rpcs          []rpcModel
}

type fileModel struct {
	PackageName string
	Imports     []string
	Services    []serviceModel
}