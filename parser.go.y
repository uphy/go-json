%{
package json

import (
    "strconv"
)
%}

%union{
    object Object
    array Array
    member Member
    members []Member
    value Value
    values []Value
    token Token
    string string
}
%type<object> object
%type<members> members
%type<member> member
%type<value> value
%type<array> array
%type<values> values
%type<string> string

%token<token> INT FLOAT STRING TRUE FALSE NULL

%%

json :
    object
    {
        yylex.(*Lexer).result = &$1
    }
    | array
    {
        yylex.(*Lexer).result = $1
    }

object :
   '{' members '}'
    {
        $$ = Object{$2}
    }
    | '{' '}'
    {
        $$ = Object{[]Member{}}
    }

members :
    member
    {
        $$ = []Member{$1}
    }
    | members ',' member
    {
        $$ = append($1, $3)
    }

member :
    string ':' value
    {
        $$ = Member{$1, $3}
    }

value :
    INT
    {
        v, _ := strconv.ParseInt($1.literal, 10, 64)
        $$ = Value{v}
    }
    | FLOAT
    {
        v, _ := strconv.ParseFloat($1.literal, 64)
        $$ = Value{v}
    }
    | TRUE
    {
        $$ = Value{true}
    }
    | FALSE
    {
        $$ = Value{false}
    }
    | STRING
    {
        $$ = Value{$1.literal[1:len($1.literal)-1]}
    }
    | array
    {
        $$ = Value{$1}
    }
    | NULL
    {
        $$ = Value{nil}
    }

array :
    '[' values ']'
    {
        $$ = $2
    }
    | '[' ']'
    {
        $$ = Array{}
    }

values :
    value
    {
        $$ = []Value{$1}
    }
    | values ',' value
    {
        $$ = append($1, $3)
    }

string :
    STRING
    {
        $$ = $1.literal[1:len($1.literal)-1]
    }

%%
