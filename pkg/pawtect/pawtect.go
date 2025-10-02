//go:generate go run ./wasmer-gen -outdir wasm -wasmfile pawtect.wasm

package pawtect

import (
	_ "embed"
	"log"

	"github.com/wasmerio/wasmer-go/wasmer"
	"github.com/wvdschel/wplace-worker/pkg/pawtect/wasm"
)

//go:embed pawtect.wasm
var pawtectWasmModule []byte

type Pawtect struct {
	*wasm.Wasm
	engine *wasmer.Engine
}

func Load() (*Pawtect, error) {
	p := &Pawtect{}

	p.engine = wasmer.NewEngine()
	store := wasmer.NewStore(p.engine)

	module, err := wasmer.NewModule(store, pawtectWasmModule)
	if err != nil {
		return nil, err
	}

	p.Wasm, err = wasm.New(store, module)
	if err != nil {
		return nil, err
	}

	importObject := wasmer.NewImportObject()
	p.Wasm.Registerwbg(importObject)

	p.Wasm.Instance, err = wasmer.NewInstance(module, importObject)
	if err != nil {
		return nil, err
	}
	_, err = p.WbindgenStart()
	if err != nil {
		log.Printf("failed to call __wbindgen_start: %s", err.Error())
		return nil, err
	}

	return p, nil
}

func (p *Pawtect) SetGlobal(value any) error {
	glob, err := p.Wasm.Instance.Exports.GetGlobal("global")
	if err != nil {
		return nil
	}
	glob.Set(value, wasmer.F32)
	return err
}

func (p *Pawtect) GetGlobal(value any) (any, error) {
	glob, err := p.Wasm.Instance.Exports.GetGlobal("global")
	if err != nil {
		return nil, nil
	}
	return glob.Get()
}

// Load procedure:
// m.set_user_id(userid)
// X = m.get_load_payload()
// GET /load with body:
//    {"pawtectMe":X,"paint-the":"world","but-not":"using-bots","security":"/.well-known/security.txt"}

// Draw procedure:
// m.set_user_id(userid)
// m.request_url(url-utf8, urllen)
// X = m.get_pawtected_endpoint_payload(payload-data-utf8, payload-data-len);
// POST with body:
//   {"colors":[19],"coords":[349,606],"t": cf-token,"fp":???}
// and headers:
// 	 x-pawtect-token: X
// Values for fp (browser fingerprint):
//    firefox: 8cd6529df58ff24cc4bf4abe31db9ae0
//    chromium: 20c5d37a4996e3a28c486ea6eecef3c3
