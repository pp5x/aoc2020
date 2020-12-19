%{
 let rules_array = Array.make 200 Ast.NIL

 type rule_def =
   {
     id: int;
     body: Ast.rule
   }

%}

%token <int> INT
%token <char> CHAR
%token <string> STR
%token EOF

%token OR
%left OR

%token COLON
%token NEWLINE
%left NEWLINE

%start root
%type <rule_def> rule
%type <Ast.rule> rule_body
%type <Ast.input> root
%%

ruleseq:
  | INT     { [$1] }
  | INT ruleseq { $1::$2 }

rule_body:
  | STR                 {
        let len = String.length $1 in
        if len = 1 then
          Ast.CHAR (String.get $1 0)
        else
          Ast.STR $1
      }
  | ruleseq             { Ast.RULES $1 }
  | ruleseq OR ruleseq  { Ast.OR ($1, $3) }

rule:
  | INT COLON rule_body {
          { id = $1;
            body = $3;
          }
        }

rules:
  | rule {
        let r = $1 in
        Array.set rules_array r.id r.body
      }
  | rules NEWLINE rules {}
  | rules NEWLINE {}

candidates:
  | STR { [ $1 ] }
  | STR NEWLINE { [ $1 ] }
  | STR NEWLINE candidates { $1::$3 }

root:
  | rules NEWLINE candidates EOF {
            {
              rules = rules_array;
              candidates = $3
            }
          }
