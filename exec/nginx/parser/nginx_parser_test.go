/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func Test1(t *testing.T) {
	//input, err := antlr.NewFileStream("test.conf")
	//if err != nil {
	//	panic(err)
	//}
	input := antlr.NewInputStream(`rewrite_by_lua_block {
		local uri=ngx.var.uri;
		if uri == "/tt"
		then
			ngx.say(uri);
			ngx.exit(200);
		end
	}`)
	lexer := NewNginxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewNginxParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Config()
	fmt.Println(tree.ToStringTree(nil, p))
	//visitor := newMappingVisitor()
	//config := tree.Accept(visitor).(*Config)
	//fmt.Println(config)
	//
	//file, err := os.OpenFile("out.conf", os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//config.EasyDumpToFile(file)
}

func Test2(t *testing.T) {
	config, _ := LoadConfig("test.conf")
	config.EasyDumpToFile("out.conf")
}
