type rule =
  | NIL
  | CHAR of char
  | STR of string
  | RULES of int list
  | OR of int list * int list

type input =
  {
    rules: rule array;
    candidates: string list;
  }
