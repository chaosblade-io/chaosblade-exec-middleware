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
  Value NEWLINE? ((Value | regexp)  NEWLINE? )*
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
  '"' StringCharacters? '"'
  ;

fragment RegexpPrefix : [~][*]?;

fragment StringCharacters : (~["\\] | EscapeSequence)+;

fragment NON_ASCII :  '\u0080'..'\uFFFF';

fragment EscapeSequence
    :   '\\' [btnfr"'\\]?
    ;

SINGLE_QUOTED
:
'\'' ~['\\]* '\'';

WS
:
[ \t]+ -> skip;

NEWLINE
:[\n\r]+ ;
