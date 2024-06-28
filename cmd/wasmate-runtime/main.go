package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/NethermindEth/juno/core/felt"
	snutils "github.com/NethermindEth/starknet.go/utils"
	extism "github.com/extism/go-sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
	"github.com/urfave/cli"
	"github.com/wasmate/wasmate-runtime/core/wmr"
	"github.com/wasmate/wasmate-runtime/pkg/confer"
	"github.com/wasmate/wasmate-runtime/pkg/ipfs"
	"github.com/wasmate/wasmate-runtime/pkg/starknetutils"
	utils "github.com/wasmate/wasmate-runtime/utils"
)

var WMR *wmr.WMR

func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

func main() {
	app := cli.NewApp()

	app.Name = APP_NAME_WASM_WORKER

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "config, c",
			Usage:    "init config file",
			EnvVar:   APP_CONFIG_ENV_NAME,
			Required: false,
		},
	}

	app.Before = func(ctx *cli.Context) error {
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.LstdFlags)
		printBanner()
		return nil
	}

	app.Action = func(ctx *cli.Context) error {
		os.Args = os.Args[:1]
		defer func() {
			if e := recover(); e != nil {
				log.Printf("app.Action recover %v \n", e)
			}
		}()

		configFile := ctx.String("config")
		if len(configFile) != 0 && !utils.IsFileExist(configFile) {
			return fmt.Errorf("%s does not exist", configFile)
		}

		if len(configFile) == 0 {
			switch app.Name {
			case APP_NAME_WASM_WORKER:
				configFile = "config/wmr_worker.yaml"
			default:
				configFile = "config/config.yaml"
			}
		}

		//Initialize configuration
		conf, err := confer.InitConfig(app.Name, configFile)
		if err != nil {
			log.Panicln(err)
		}

		//Whether to start golang pprof debugging
		if conf.Opts.DebugConf.Enable {
			utils.GoWithRecover(func() {
				startDebug(conf.Opts.DebugConf.PprofBindAddr)
			}, nil)
		}

		c, cancel := context.WithCancel(context.Background())
		offlineCtx, offlineCancel := context.WithCancel(c)
		WMR = wmr.NewWMR(c, offlineCtx, conf)

		initFunc := []func() error{
			WMR.LoadBasics,
		}

		switch conf.Opts.ApptypeConf {
		case confer.APP_TYPE_WASM_WORKER:
			//initFunc = append(initFunc, nil)
		default:
			panic("run APP type error:" + conf.Opts.ApptypeConf)

		}

		for k := range initFunc {
			err := initFunc[k]()
			if err != nil {
				log.Println("init function run fall:", k, err)
				panic(err)
			}
		}

		//wasm engine : create API for create wasm plugin into WasmMeta
		ctxWasm := context.Background()
		config := extism.PluginConfig{
			ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
			EnableWasi:   true,
		}

		manifest := extism.Manifest{Wasm: []extism.Wasm{}}

		//WASM Modules files path
		if conf.Opts.WASMModulesFiles.Enable {
			for _, wasmPath := range conf.Opts.WASMModulesFiles.WASMFilePaths {
				if _, err := os.Stat(wasmPath); err != nil {
					log.Panicln(err)

				}

				manifest.Wasm = append(manifest.Wasm, extism.WasmFile{Path: wasmPath}) // if file exist then put wasm file to manifest
			}
		}

		//WASM Modules From IPFS CID
		if conf.Opts.WASMModulesIPFS.Enable {
			ipfsC := ipfs.NewIPFSClient(conf.Opts.WASMModulesIPFS.LassieNet.Scheme,
				conf.Opts.WASMModulesIPFS.LassieNet.Host,
				conf.Opts.WASMModulesIPFS.LassieNet.Port)

			for _, wasmCID := range conf.Opts.WASMModulesIPFS.CIDS {

				D, err := ipfs.GetDATAFromIPFSCID(ipfsC, wasmCID)
				if err != nil {
					log.Panicln(err)
				}

				for _, wasmData := range D {
					manifest.Wasm = append(manifest.Wasm, extism.WasmData{Data: wasmData})
				}

			}

		}

		//WASM Modules From StarkNet
		if conf.Opts.WASMModulesStarknet.Enable {
			ipfsC := ipfs.NewIPFSClient(conf.Opts.WASMModulesStarknet.LassieNet.Scheme,
				conf.Opts.WASMModulesStarknet.LassieNet.Host,
				conf.Opts.WASMModulesStarknet.LassieNet.Port)

			for _, wasmFuncName := range conf.Opts.WASMModulesStarknet.WasmFuncNames {

				fcnamefelt, err := snutils.HexToFelt(hex.EncodeToString([]byte(wasmFuncName)))
				if err != nil {
					log.Panicln(err)
				}

				callResp, err := starknetutils.CallReadFunction(
					conf.Opts.WASMModulesStarknet.RpcAddress,
					conf.Opts.WASMModulesStarknet.SmartContract,
					conf.Opts.WASMModulesStarknet.ContractMethod,
					[]*felt.Felt{fcnamefelt})

				if err != nil {
					log.Panicln(err)
				}

				cidDataBigIntArr := snutils.FeltArrToBigIntArr(callResp)

				if len(cidDataBigIntArr) != 2 || cidDataBigIntArr[0].Int64() == 0 || cidDataBigIntArr[1].Int64() == 0 {
					log.Panicln("get cid data from starknet is wrong")
				}

				cidAddressPart1, err := hex.DecodeString(snutils.BigToHex(cidDataBigIntArr[0])[2:])
				if err != nil {
					log.Panicln(err)
				}
				cidAddressPart2, err := hex.DecodeString(snutils.BigToHex(cidDataBigIntArr[1])[2:])
				if err != nil {
					log.Panicln(err)
				}
				cidAddress := string(cidAddressPart1) + string(cidAddressPart2)

				D, err := ipfs.GetDATAFromIPFSCID(ipfsC, cidAddress)
				if err != nil {
					log.Panicln(err)
				}

				for _, wasmData := range D {
					manifest.Wasm = append(manifest.Wasm, extism.WasmData{Data: wasmData})
				}

				if wasmFuncName != "" && cidAddress != "" {
					log.Printf("Successfully loaded WASM function '%s' from StarkNet smart contract at CID '%s'\n", wasmFuncName, cidAddress)
				} else {
					log.Printf("Failed to load WASM function. Invalid inputs: wasmFunctionName='%s', contractCidAddress='%s'\n", wasmFuncName, cidAddress)
				}

			}

		}

		// Create an instance of the plugin
		pluginInst, err := WMR.WMRPlugin.NewWMRPlugin(ctxWasm, manifest, config, nil)
		if err != nil {
			log.Panicln(err)
		}

		//add wasm plugin by name: such as "StarkFAAS-hello"
		err = WMR.WMRPlugin.AddPlugin("StarkFAAS-hello", *pluginInst, true)

		if err != nil {
			log.Panicln(err)
		}

		//
		FASSFrame := WMR.NewHTTPServer(fiber.Config{DisableStartupMessage: true})

		FASSFrame.Post("/", sayHelloFiber)

		fmt.Println("🌍 http server is listening on:", conf.Opts.NetWork.BindAddress)
		FASSFrame.Listen(conf.Opts.NetWork.BindAddress)

		_ = cancel
		_ = offlineCtx
		_ = offlineCancel
		_ = initFunc

		return nil
	}

	//app run
	err := app.Run(os.Args)
	if err != nil {
		log.Printf("app.Run error: %v\n", err)
	}
}
