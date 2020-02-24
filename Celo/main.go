package main

//                                                  jim380 <admin@cyphercore.io>
//  ============================================================================
//
//  Copyright (C) 2020 jim380
//
//  Permission is hereby granted, free of charge, to any person obtaining
//  a copy of this software and associated documentation files (the
//  "Software"), to deal in the Software without restriction, including
//  without limitation the rights to use, copy, modify, merge, publish,
//  distribute, sublicense, and/or sell copies of the Software, and to
//  permit persons to whom the Software is furnished to do so, subject to
//  the following conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
//  CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN setup OF CONTRACT,
//  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
//  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//  ============================================================================

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/node_tooling/Celo/bot"
	"github.com/node_tooling/Celo/cmd"
	"github.com/node_tooling/Celo/setup"
	"github.com/node_tooling/Celo/util"
)

func main() {
	var machine string
	var cmdInput bool
	var teleBot bool

	flag.BoolVar(&cmdInput, "cmd", false, "Show election details")
	flag.BoolVar(&teleBot, "bot", false, "Run the telegram bot")
	flag.Parse()

	if cmdInput {
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		util.SetEnv()
		cmd.OptionsAll()
	} else if teleBot {
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		util.SetEnv()
		bot.Run()
	} else if !cmdInput {
		//fmt.Println("Invalid flag value. flag.Args() is:", flag.Args())
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		util.SetEnv()
		message := "Which machine are you on:\n\n1) Local\n2) Validator\n3)" +
			" " + "Proxy\n4) Attestation\n\nEnter down below (e.g. \"1\" or \"Local\"): "
		machine = util.InputReader(message, machine)
		setup.NodeStop(machine)
		setup.KeyCheck(machine)
		setup.ChainDataDel(machine)
		setup.NodeRun(machine)
	}
}
