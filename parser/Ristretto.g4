grammar Ristretto;

start: doctype? tag*;

doctype: DOCTYPE;

tag: NUMBERSIGN LSBRACKET tagWithAttr RSBRACKET ';' # closedTag
    | NUMBERSIGN LSBRACKET tagWithAttr RSBRACKET LCBRACKET (tag | content)* RCBRACKET # openedTag
    ;

tagname: IDENTITY;

content: STRING;

tagWithAttr: tagname attrs?;

attrs: attr (',' attr)*;

attr: IDENTITY '=' attrValue;

attrValue: IDENTITY | STRING;

IDENTITY: (DIGIT | LETTER)+;

STRING: '"' ~('"' | '\r' | '\n' | '\t')* '"';

DOCTYPE: '![DOCTYPE]';

NUMBERSIGN: '#';

LSBRACKET: '[';
RSBRACKET: ']';

LPAREN: '(';
RPAREN: ')';

LCBRACKET: '{';
RCBRACKET: '}';

WS: (' ' | '\r' | '\n' | '\t')+ -> skip;

DIGIT: [0-9];

fragment LETTER: [a-zA-Z];
