grammar Nginx;

config
  :
  (statement | block)+
  ;

statement
   :
   ( rewriteStatement | genericStatement | regexHeaderStatement) ';'
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

block
  :
  ( locationBlockHeader | genericBlockHeader )
  '{'
  ( statement | block | ifStatement )*
  '}'
  ;

genericBlockHeader
  :
  Value ( Value | regexp )*
  ;

ifStatement
  :
  'if'
  ifBody
  '{'
    (statement)*
  '}'
  ;

ifBody
  :
  '('
  Value
  (Value)?
  (
    Value
    |
    regexp
  )?
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

//QUOTED_STRING
//: '"' (~('"' | '\\' | '\r' | '\n') | '\\' ('"' | '\\'))* '"';


Value: STR_EXT | QUOTED_STRING | SINGLE_QUOTED
;

STR_EXT
  :
 ([a-zA-Z0-9_/.,\-:=~+!?$&^*[\]@|#] | NON_ASCII)+
  ;

LINE_COMMENT: (
        ('-- ' | '#') ~[\r\n]* ('\r'? '\n' | EOF)
        | '--' ('\r'? '\n' | EOF)
    ) -> skip;

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
[ \t\n\r]+ -> skip ;
