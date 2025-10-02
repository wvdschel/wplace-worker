package wasmergen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wasmerio/wasmer-go/wasmer"
)

type Generator struct {
	engine *wasmer.Engine
	store  *wasmer.Store
}

func New() *Generator {
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	return &Generator{
		engine: engine,
		store:  store,
	}
}

func (g *Generator) GenerateImportStubs(wasmModule []byte, packageDir string) error {
	module, err := wasmer.NewModule(g.store, wasmModule)
	if err != nil {
		return err
	}

	exports := module.Exports()
	imports := module.Imports()
	importedModules := make(map[string][]*wasmer.ImportType)
	for _, imp := range imports {
		moduleName := imp.Module()
		importedModules[moduleName] = append(importedModules[moduleName], imp)
	}

	os.MkdirAll(packageDir, 0755)
	pkgName := path.Base(packageDir)
	if pkgName == "." || pkgName == "/" {
		pkgName = "wasmergen"
	}

	filePath := filepath.Join(packageDir, pkgName+".go")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := g.generateStruct(file, pkgName); err != nil {
		return err
	}

	if err := g.generateConstructor(file, pkgName, importedModules); err != nil {
		return err
	}

	filePath = filepath.Join(packageDir, "exports.go")
	exportFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer exportFile.Close()

	if err := g.generateExports(exportFile, pkgName, exports); err != nil {
		return err
	}

	for moduleName, imports := range importedModules {
		filePath := filepath.Join(packageDir, moduleName+".go")
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		fmt.Fprintf(file, "package %s\n\n", pkgName)
		fmt.Fprintf(file, "import \"fmt\"\n\n")
		fmt.Fprintf(file, "import \"github.com/wasmerio/wasmer-go/wasmer\"\n\n")

		if err := g.generateRegisterFunction(file, pkgName, moduleName, imports); err != nil {
			return err
		}

		for _, importType := range imports {
			if err := g.generateImportFunction(file, pkgName, importType); err != nil {
				return err
			}
		}

		file.Close()
	}
	return nil
}

func mangleFunctionName(importType *wasmer.ImportType) string {
	return importType.Module() + "_" + importType.Name()
}

func mangleStructName(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[0:1]) + name[1:]
}

func mangleExportName(name string) string {
	if name == "" {
		return ""
	}
	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = mangleStructName(part)
	}
	return strings.Join(parts, "")
}

func stringifyValueTypes(params ...*wasmer.ValueType) string {
	res := []string{}

	for _, param := range params {
		switch param.Kind() {
		case wasmer.I32:
			res = append(res, "wasmer.I32")
		case wasmer.I64:
			res = append(res, "wasmer.I64")
		case wasmer.F32:
			res = append(res, "wasmer.F32")
		case wasmer.F64:
			res = append(res, "wasmer.F64")
		case wasmer.AnyRef:
			res = append(res, "wasmer.AnyRef")
		case wasmer.FuncRef:
			res = append(res, "wasmer.FuncRef")
		default:
			panic("Unknown value kind")
		}
	}

	return strings.Join(res, ", ")
}

func (g *Generator) generateStruct(file *os.File, pkgName string) error {
	fmt.Fprintf(file, "package %s\n\n", pkgName)
	fmt.Fprintf(file, "import \"github.com/wasmerio/wasmer-go/wasmer\"\n\n")

	fmt.Fprintf(file, "type %s struct {\n", mangleStructName(pkgName))
	fmt.Fprintf(file, "\tStore *wasmer.Store\n")
	fmt.Fprintf(file, "\tInstance *wasmer.Instance\n")
	fmt.Fprintf(file, "\tModule *wasmer.Module\n")
	fmt.Fprintf(file, "}\n\n")

	return nil
}

func (g *Generator) generateRegisterFunction(file *os.File, pkgName string, moduleName string, imports []*wasmer.ImportType) error {
	fmt.Fprintf(file, "func (o *%s) Register%s(importObj *wasmer.ImportObject) {\n", mangleStructName(pkgName), moduleName)
	fmt.Fprintf(file, "\t// Register the module\n")

	fmt.Fprintf(file, "\timportObj.Register(\"%s\", map[string]wasmer.IntoExtern{\n", moduleName)
	for _, importType := range imports {
		funcType := importType.Type().IntoFunctionType()

		if funcType == nil {
			fmt.Printf("skipping %s because it is not a function\n", importType.Name())
			continue
		}

		fmt.Fprintf(file, "\t\t\"%s\": wasmer.NewFunction(\n", importType.Name())
		fmt.Fprintf(file, "\t\t\to.Store,\n")
		fmt.Fprintf(file, "\t\t\twasmer.NewFunctionType(\n")
		fmt.Fprintf(file, "\t\t\t\twasmer.NewValueTypes(%s),\n", stringifyValueTypes(funcType.Params()...))
		fmt.Fprintf(file, "\t\t\t\twasmer.NewValueTypes(%s),\n", stringifyValueTypes(funcType.Results()...))
		fmt.Fprintf(file, "\t\t\t),\n")
		fmt.Fprintf(file, "\t\t\to.%s,\n", mangleFunctionName(importType))
		fmt.Fprintf(file, "\t\t),\n")
	}
	fmt.Fprintf(file, "\t})\n")

	fmt.Fprintf(file, "}\n\n")
	return nil
}

func (g *Generator) generateImportFunction(file *os.File, pkgName string, importType *wasmer.ImportType) error {
	funcType := importType.Type().IntoFunctionType()

	if funcType == nil {
		fmt.Printf("skipping %s because it is not a function\n", importType.Name())
		return nil
	}

	fmt.Fprintf(file, "func (o *%s) %s(args []wasmer.Value) ([]wasmer.Value, error) {\n", mangleStructName(pkgName), mangleFunctionName(importType))
	fmt.Fprintf(file, "\tfmt.Printf(\"> wasm import call: %s.%s(%%v)\\n\", args)\n", importType.Module(), importType.Name())
	fmt.Fprintf(file, "\t// TODO implement import '%s'\n", importType.Name())
	fmt.Fprintf(file, "\treturn []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil\n")
	fmt.Fprint(file, "}\n\n")

	return nil
}

func (g *Generator) generateConstructor(file *os.File, pkgName string, importedModules map[string][]*wasmer.ImportType) error {
	fmt.Fprintf(file, "func New(store *wasmer.Store, module *wasmer.Module) (*%s, error) {\n", mangleStructName(pkgName))
	fmt.Fprintln(file, "\tres := &Wasm{")
	fmt.Fprintln(file, "\t\tStore: store,")
	fmt.Fprintln(file, "\t\tModule: module,")
	fmt.Fprint(file, "\t}\n\n")

	fmt.Fprintf(file, "\timportObject := wasmer.NewImportObject()\n")
	for mod := range importedModules {
		fmt.Fprintf(file, "\tres.Register%s(importObject)\n", mod)
	}

	fmt.Fprintln(file, "\n\tinst, err := wasmer.NewInstance(module, importObject)")
	fmt.Fprintln(file, "\tif err != nil {")
	fmt.Fprintln(file, "\t\treturn nil, err")
	fmt.Fprintln(file, "\t}")
	fmt.Fprintln(file, "\n\tres.Instance = inst")
	fmt.Fprintln(file, "\n\treturn res, nil")
	fmt.Fprintln(file, "}")

	return nil
}

func (g *Generator) generateExports(file *os.File, pkgName string, exports []*wasmer.ExportType) error {
	fmt.Fprintf(file, "package %s\n\n", pkgName)
	fmt.Fprintf(file, "import \"github.com/wasmerio/wasmer-go/wasmer\"\n\n")
	for _, exp := range exports {
		switch exp.Type().Kind() {
		case wasmer.FUNCTION:
			g.generateExportFunction(file, pkgName, exp)
		case wasmer.GLOBAL:
			g.generateExportGlobal(file, pkgName, exp)
		case wasmer.MEMORY:
			g.generateExportMemory(file, pkgName, exp)
		case wasmer.TABLE:
			g.generateExportTable(file, pkgName, exp)
		default:
			return fmt.Errorf("unknown export type %s: %v", exp.Name(), exp.Type().Kind())
		}
	}
	return nil
}

func (g *Generator) generateExportFunction(file *os.File, pkgName string, export *wasmer.ExportType) {
	fmt.Fprintf(file, "func (o *%s) %s(args ...any) (any, error) {\n", mangleStructName(pkgName), mangleExportName(export.Name()))
	fmt.Fprintf(file, "\twasmFunc, err := o.Instance.Exports.GetFunction(\"%s\")\n", export.Name())
	fmt.Fprintln(file, "\tif err != nil { return nil, err }")
	fmt.Fprint(file, "\treturn wasmFunc(args...)\n}\n\n")
}

func (g *Generator) generateExportMemory(file *os.File, pkgName string, export *wasmer.ExportType) {
	fmt.Fprintf(file, "func (o *%s) %s() *wasmer.Memory {\n", mangleStructName(pkgName), mangleExportName(export.Name()))
	fmt.Fprintf(file, "\tmem, _ := o.Instance.Exports.GetMemory(\"%s\")\n", export.Name())
	fmt.Fprint(file, "\treturn mem\n}\n\n")
}

func (g *Generator) generateExportTable(file *os.File, pkgName string, export *wasmer.ExportType) {
	fmt.Fprintf(file, "func (o *%s) %s() *wasmer.Table {\n", mangleStructName(pkgName), mangleExportName(export.Name()))
	fmt.Fprintf(file, "\ttable, _ := o.Instance.Exports.GetTable(\"%s\")\n", export.Name())
	fmt.Fprint(file, "\treturn table\n}\n\n")
}

func (g *Generator) generateExportGlobal(file *os.File, pkgName string, export *wasmer.ExportType) {
	glob := export.Type().IntoGlobalType()

	if glob.Mutability() == wasmer.MUTABLE {
		// Generate setter
		ttype := glob.ValueType()

		fmt.Fprintf(file, "func (o *%s) Set%s(value any) error {\n", mangleStructName(pkgName), mangleExportName(export.Name()))
		fmt.Fprintf(file, "\tglob, err := o.Instance.Exports.GetGlobal(\"%s\")\n", export.Name())
		fmt.Fprintln(file, "\tif err != nil { return nil, err }")
		fmt.Fprintf(file, "\treturn glob.Set(value, %s)\n", stringifyValueTypes(ttype))
		fmt.Fprint(file, "}\n\n")
	}

	fmt.Fprintf(file, "func (o *%s) Get%s() (any, error) {\n", mangleStructName(pkgName), mangleExportName(export.Name()))
	fmt.Fprintf(file, "\tglob, err := o.Instance.Exports.GetGlobal(\"%s\")\n", export.Name())
	fmt.Fprintln(file, "\tif err != nil { return nil, err }")
	fmt.Fprintf(file, "\treturn glob.Get()\n")
	fmt.Fprint(file, "}\n\n")
}
