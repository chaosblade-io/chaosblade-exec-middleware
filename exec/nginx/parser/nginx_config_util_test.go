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

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		isFile bool
	}{
		{
			name:   "testConfigParser",
			input:  "test.conf",
			isFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lexer *NginxLexer
			if tt.isFile {
				input, err := antlr.NewFileStream(tt.input)
				if err != nil {
					t.Errorf(fmt.Sprintf("parser test err: %s", err))
				}
				lexer = NewNginxLexer(input)
			} else {
				input := antlr.NewInputStream(tt.input)
				lexer = NewNginxLexer(input)
			}

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := NewNginxParser(stream)
			//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Config()
			//fmt.Println(tree.ToStringTree(nil, p))
			visitor := newMappingVisitor()
			config := tree.Accept(visitor).(*Config)
			if config == nil {
				t.Errorf("LoadConfig() err")
			}
		})
	}
}

func TestConfig_FindBlock(t *testing.T) {
	tests := []struct {
		locator string
		err     bool
	}{
		{locator: "", err: true},
		{locator: "global", err: false},
		{locator: "http.global", err: true},
		{locator: "ccxx", err: true},
		{locator: "http[0]", err: true},
		{locator: "http.server[0]", err: false},
		{locator: "http.server[1]", err: true},
		{locator: "http.server[-1]", err: true},
		{locator: "http.server[0].location[0]", err: false},
	}

	for _, tt := range tests {
		config, _ := LoadConfig("test.conf")
		t.Run("dd", func(t *testing.T) {
			_, err := config.FindBlock(tt.locator)
			if (tt.err && err == nil) || (!tt.err && err != nil) {
				t.Errorf("test case err : %#v, %s", tt, err)
			} else if err == nil {
				err := config.SetStatement(tt.locator, "test", "value", true)
				if err != nil {
					t.Error(err)
				}
				statements, _ := config.FindBlock(tt.locator)
				has := false
				for _, s := range *statements {
					if s.Key == "test" && s.Value == "value" {
						has = true
						break
					}
				}
				if !has {
					t.Errorf("test case err, can't set statement : %#v", tt)
				}
			}
		})
	}
}
