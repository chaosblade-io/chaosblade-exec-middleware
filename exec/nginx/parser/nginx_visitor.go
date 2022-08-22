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
