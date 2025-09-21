package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wvdschel/wplace-worker/pkg/wplace/bot"
)

var (
	cfgPath   = flag.String("config", "config.json", "config file path")
	genConfig = flag.Bool("generate-config", false, "generate sample config file")
)

func main() {
	flag.Parse()
	cfg := &bot.Config{}

	if *genConfig {
		if _, err := os.Stat(*cfgPath); !os.IsNotExist(err) {
			fmt.Printf("config file already exists - refusing to generate: %v\n", err)
			os.Exit(1)
		}
		cfg = bot.ExampleConfig()
		b, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(*cfgPath, b, 0644); err != nil {
			fmt.Printf("failed to save config file: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	f, err := os.Open(*cfgPath)
	if err != nil {
		fmt.Printf("failed to load config file: %v\n", err)
		os.Exit(1)
	}

	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		fmt.Printf("failed to parse config file: %v\n", err)
		os.Exit(1)
	}
	b, err := bot.New(cfg)
	if err != nil {
		fmt.Printf("failed to initialize bot: %v\n", err)
		os.Exit(1)
	}

	if err := b.Run(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("exiting")
}
