/*
 * Copyright 2025 The ChaosBlade Authors
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

// Code generated from Nginx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nginx

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by NginxParser.
type NginxVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NginxParser#config.
	VisitConfig(ctx *ConfigContext) interface{}

	// Visit a parse tree produced by NginxParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by NginxParser#genericStatement.
	VisitGenericStatement(ctx *GenericStatementContext) interface{}

	// Visit a parse tree produced by NginxParser#regexHeaderStatement.
	VisitRegexHeaderStatement(ctx *RegexHeaderStatementContext) interface{}

	// Visit a parse tree produced by NginxParser#luaBlock.
	VisitLuaBlock(ctx *LuaBlockContext) interface{}

	// Visit a parse tree produced by NginxParser#luaStatement.
	VisitLuaStatement(ctx *LuaStatementContext) interface{}

	// Visit a parse tree produced by NginxParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by NginxParser#genericBlockHeader.
	VisitGenericBlockHeader(ctx *GenericBlockHeaderContext) interface{}

	// Visit a parse tree produced by NginxParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by NginxParser#ifBody.
	VisitIfBody(ctx *IfBodyContext) interface{}

	// Visit a parse tree produced by NginxParser#regexp.
	VisitRegexp(ctx *RegexpContext) interface{}

	// Visit a parse tree produced by NginxParser#locationBlockHeader.
	VisitLocationBlockHeader(ctx *LocationBlockHeaderContext) interface{}

	// Visit a parse tree produced by NginxParser#rewriteStatement.
	VisitRewriteStatement(ctx *RewriteStatementContext) interface{}
}
