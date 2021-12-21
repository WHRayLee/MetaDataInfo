> Yacc reads the grammar descriptions in bas.y and generates a syntax analyzer (parser), that
includes function yyparse, in file y.tab.c. Included in file bas.y are token declarations. The –d
option causes yacc to generate definitions for tokens and place them in file y.tab.h. Lex reads the
pattern descriptions in bas.l, includes file y.tab.h, and generates a lexical analyzer, that includes
function yylex, in file lex.yy.c.

> From main we call yyparse to run the compiler. Function yyparse automatically calls yylex to
obtain each token.

> During the first phase the compiler reads the input and converts strings in the source to tokens.
With regular expressions we can specify patterns to lex so it can generate code that will allow it to
scan and match strings in the input. Each pattern specified in the input to lex has an associated
action.

> Typically an action returns a token that represents the matched string for subsequent use
by the parser. Initially we will simply print the matched string rather than return a token value.

> Any regular expression expressions may be expressed as a finite state automaton (FSA). We can
represent an FSA using states, and transitions between states. There is one start state and one
or more final or accepting states.

> ECHO is a macro that writes code matched by the pattern. This is
the default action for any unmatched strings. Typically, ECHO is defined as:

```c 
#define ECHO fwrite(yytext, yyleng, 1, yyout)
```

- Variable yytext is a pointer to the matched string (NULL-terminated)
- yyleng is the length of the matched string
- Variable yyout is the output file and defaults to stdout
- Function yywrap is called by lex when input is exhausted. Return 1 if you are done or 0 if more processing is required
- yylex(void) call to invoke lexer, returns token
- yylval value associated with token
- FILE *yyout output file
- FILE *yyin input file
- INITIAL initial start condition
- BEGIN condition switch start condition
- ECHO write matched string

```c
%{ 
 int yylineno; 
 %} 
 %% 
 ^(.*)\n printf("%4d\t%s", ++yylineno, yytext); 
 %% 
 int main(int argc, char *argv[]) { 
    yyin = fopen(argv[1], "r"); 
    yylex(); 
    fclose(yyin); 
 }
 ```

```c 
digit [0-9] 
letter [A-Za-z] 
%{ 
    int count; 
%} 
%% 
 /* match identifier */ 
{letter}({letter}|{digit})* count++; 
%% 
int main(void) { 
    yylex(); 
    printf("number of identifiers = %d\n", count); 
    return 0; 
}
```

## Yacc
```c
E -> E + E 
E -> E * E 
E -> id 
```
>Terms that appear on the left-hand side (lhs) of a production, such as E (expression) are nonterminals. Terms such as id (identifier) are terminals (tokens returned by lex) and only appear on the right-hand side (rhs) of a production

> At each step we expanded a term and replace the lhs of a production with the corresponding rhs.

>  Instead of starting with a single nonterminal (start symbol) and generating an expression from a grammar we need to reduce an expression to a single nonterminal. This is known as bottom-up or shift-reduce parsing and uses a stack for storing terms.

> Terms to the left of the dot are on the stack while remaining input is to the right of the dot. We start by shifting tokens onto the stack. When the top of the stack matches the rhs of a production we replace the matched tokens on the stack with the lhs of the production. In other words the matched tokens of the rhs are popped off the stack, and the lhs of the production is pushed on the stack. The matched tokens are known as a handle and we are reducing the handle to the lhs of the production. This process continues until we have shifted all input to the stack and only the starting nonterminal remains on the stack.

>  The definitions section consists of token declarations and C code bracketed by “%{“ and “%}”. The BNF grammar is placed in the rules section and user subroutines are added in the subroutines section

## The Definition Section
> The definition section includes declarations of the tokens used in the grammar, the types
of values used on the parser stack, and other odds and ends.

## The Rules Section
> we use a colon between the
left- and right-hand sides of a rule, and we put a semicolon at the end of each rule

> The symbol on the left-hand side of the first rule in the grammar is normally the start
symbol, though you can use a %start declaration in the definition section to override
that.

## Symbol Values and Actions
> Every symbol in a yacc parser has a value. The value gives additional information about a particular instance of a symbol. If a symbol represents a number, the value would be the particular number. If it represents a literal text string, the value would probably be a pointer to a copy of the string. If it represents a variable in a program, the value would be a pointer to a symbol table entry describing the variable.

> Non-terminal symbols can have any values you want, created by code in the parser. 

> Often the action code builds a parse tree corresponding to the input, so that later code can process a whole statement or even a whole program at a time.

> If you have multiple value types, you have to list all the value types used in a parser so that yacc can create a C union typedef called YYSTYPE to contain them

> The action code can refer to the values of the right-hand side symbols as $1, $2, ..., and can set the value of the left-hand side by setting $$

```c
%token NAME NUMBER
%%
statement: NAME '=' expression
| expression { printf("= %d\n", $1); }
;
expression: expression '+' NUMBER { $$ = $1 + $3; }
| expression '−' NUMBER { $$ = $1 − $3; }
| NUMBER { $$ = $1; }
;
```
> The action on the last rule is not strictly necessary, since the default action that yacc performs after every reduction, before running any explicit action code, assigns the value $1 to $$.

> Whenever the lexer returns a token to the parser, if the token has an associated value, the lexer must store the value in yylval before returning.

> yacc defines yylval as a union

> Precedence controls which operators to execute first in an expression,multiplication and division take precedence over addition and subtraction

> In any expression grammar, operators are grouped into levels of precedence from lowest to highest.

```go
expression: expression '+' mulexp
 | expression '−' mulexp
 | mulexp
 ;
 mulexp: mulexp '*' primary
 | mulexp '/' primary
 | primary
 ;
 primary: '(' expression ')'
 | '-' primary
 | NUMBER
 ;
```

```lex
%left '+' '-'
%left '*' '/'
%nonassoc UMINUS
```

> They tell yacc that “+” and “−” are left associative and at the lowest precedence level, “*” and “/” are left associative and at a higher precedence level, and UMINUS, a pseudo-token standing for unary minus, has no associativity and is at the highest precedence.

> 它们告诉yacc，`+`和`-`是左联动的，优先级最低，`*`和`/`是左联动的，优先级更高。处于较高的优先级，而`UMINUS`，一个代表单数减号的伪标记。没有关联性，优先级最高。

> Yacc assigns each rule the precedence of the rightmost token on the right-hand side; if the rule contains no tokens with precedence assigned, the rule has no precedence of its own.

> When yacc encounters a `shift/reduce` conflict due to an ambiguous grammar, it consults the table of precedences, and if all of the rules involved in the conflict include a token which appears in a precedence declaration, it uses precedence to resolve the conflict.

> This parser using precedences is slightly smaller and faster than the one with the extra rules for implicit precedence, since it has fewer rules to reduce

```
%token NAME NUMBER
%left '-' '+'
%left '*' '/'
%nonassoc UMINUS
%%
statement: NAME '=' expression
 | expression { printf("= %d\n", $1); }
 ;
expression: expression '+' expression { $$ = $1 + $3; }
 | expression '−' expression { $$ = $1 − $3; }
 | expression '*' expression { $$ = $1 * $3; }
 | expression '/' expression
     { if($3 == 0)
        yyerror("divide by zero");
     else
        $$ = $1 / $3;
     }
 | '−' expression %prec UMINUS { $$ = −$2; }
 | '(' expression ')' { $$ = $2; }
 | NUMBER { $$ = $1; }
 ;
%%
```

> The %prec tells yacc to use the precedence of UMINUS for this rule.

> We recommend that you use precedence in only two situations: in expression grammars, and to resolve the “dangling else” conflict in grammars for if-then-else language constructs.

## Variables and Typed Tokens
```
%{
double vbltable[26];
%{
%union {
 double dval;
 int vblno;
}
%token <vblno> NAME
%token <dval> NUMBER
%left '-' '+'
%left '*' '/'
%nonassoc UMINUS
%type <dval> expression
%%
statement_list: statement '\n'
 | statement_list statement '\n'
 ;
statement: NAME '=' expression { vbltable[$1] = $3; }
 | expression { printf("= %g\n", $1); }
 ;
expression: expression '+' expression { $$ = $1 + $3; }
 | expression '−' expression { $$ = $1 − $3; }
 | expression '*' expression { $$ = $1 * $3; }
 | expression '/' expression
 { if($3 == 0.0)
 yyerror("divide by zero");
 else
 $$ = $1 / $3;
 }
 | '−' expression %prec UMINUS { $$ = −$2; }
 | '(' expression ')' { $$ = $2; }
 | NUMBER
 | NAME { $$ = vbltable[$1]; }
 ;
%%
```

### Symbol Values and %union
> Expressions have double values, while the value for variable references and NAME symbols are integers from 0 to 25 corresponding to the slot in vbltable.

```c
 #define NAME 257
 #define NUMBER 258
 #define UMINUS 259
 typedef union {
    double dval;
    int vblno;
 } YYSTYPE;
 extern YYSTYPE yylval;
```

> Now we have to tell the parser which symbols use which type of value. We do that by putting the appropriate field name from the union in angle brackets in the lines in the definition section that defines the symbol:

> The new declaration %type sets the type for non-terminals which otherwise need no declaration.

>  In action code, yacc automatically qualifies symbol value references with the appropriate field name

> We’ve also added an action for the rule that sets a variable, and a new rule at the end that turns a NAME into an expression by fetching the value of the variable.

> The lexer doesn’t have any automatic way to associate types with tokens, so you have to put in explicit field references when you set yylval.

### Symbol Tables






