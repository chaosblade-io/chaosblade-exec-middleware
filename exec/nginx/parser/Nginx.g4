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

grammar Nginx;

config
  :
  (statement | block | luaBlock)+
  ;

statement
   :
   NEWLINE? ( rewriteStatement | genericStatement | regexHeaderStatement) ';' NEWLINE?
   ;

genericStatement
  :
  Value ( Value | regexp)*
  ;

regexHeaderStatement
  :
  // Use token definition for regexp-driven parameter name in Nginx config
  // See: http://nginx.org/en/docs/http/ngx_http_map_module.html
  REGEXP_PREFIXED Value
  ;

LUA_HEADER
    :
    [a-z_]+ 'by_lua_block'
    ;

luaBlock
   :
   NEWLINE?
   LUA_HEADER NEWLINE?
   '{'  NEWLINE?
   luaStatement *
   '}'  NEWLINE?
    ;

luaStatement
   :
   (
   ( (Value | 'if' | ')' | '(')+ ';'? )
   |
   ';'
   ) NEWLINE?
   ;

block
  :
  NEWLINE?
  ( locationBlockHeader | genericBlockHeader ) NEWLINE?
  '{' NEWLINE?
  (( statement | block | ifStatement | luaBlock ) NEWLINE? ) *
  '}' NEWLINE?
  ;

genericBlockHeader
  :
  Value ( Value | regexp )*
  ;

ifStatement
  :
  'if' NEWLINE?
  ifBody NEWLINE?
  '{' NEWLINE?
    (statement NEWLINE?)*
  '}' NEWLINE?
  ;

ifBody
  :
  '(' NEWLINE?
  Value + NEWLINE?
  ( Value | regexp)? NEWLINE?
  ')'
  ;

regexp
:
(
  '\\.'
  | '^'
  | Value
  | '(' regexp ')'
)+;

locationBlockHeader
  :
  'location'
  (Value)?
  ( Value | regexp )
;

rewriteStatement
  :
  'rewrite'
  (Value | regexp) Value
  ('last' | 'break' | 'redirect' | 'permanent')?
  ;

Value: STR_EXT | QUOTED_STRING | SINGLE_QUOTED
;

STR_EXT
  :
  ([()a-zA-Z0-9_/.,\-:=~+!?$&^*[\]@|#] | NON_ASCII)+
  ;

LINE_COMMENT: (
        ('-- ' | '#') ~[\r\n]* ('\r'? '\n' | EOF)
        | '--' ('\r'? '\n' | EOF)
    ) NEWLINE? -> skip;

REGEXP_PREFIXED
  : (RegexpPrefix)[a-zA-Z0-9_/.,\-:=~+!?$&^*[\]@|#)(]+
  ;

QUOTED_STRING
  :
  '"' (~["] | EscapeSequence)* '"'
  ;

fragment RegexpPrefix : [~][*]?;

fragment NON_ASCII :  '\u0080'..'\uFFFF';

fragment EscapeSequence
    :   '\\' [btnfr"'\\]?
    ;

SINGLE_QUOTED
:
'\'' ~[']* '\'';

WS
: [ \t]+ -> skip;

NEWLINE
:[\n\r]+ ;
