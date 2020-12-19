{
open Parser

let print_lexer_error c =
  Printf.eprintf "Unknown token: `%c`\n" c;
  exit 2
}

let integer = ['0'-'9']+
let white = [' ' '\t']
let op_or = "|"
let op_colon = ":"

rule parse_token = parse
| white+              { parse_token lexbuf }
| integer as lxm      { INT(int_of_string lxm) }
| op_or               { OR }
| op_colon            { COLON }
| '"'                 { parse_token lexbuf }
| "\n"+               { NEWLINE }
| ['a'-'z']+ as str   { STR str }
| eof                 { EOF }
| _ as c              { print_lexer_error c }
