type params = {
  str: string;
  rules: Ast.rule array;
}

let rec display = function
  | [] -> ()
  | e::t -> Printf.printf "%d" e; (display t)

let candidate_match rules candidate =
  let candidate_len = String.length candidate in
  let rec seq_rule i = function
        | [] -> 0
        | e::t ->
          let rule = Array.get rules e in
          let consumed = match_rule i rule in
              if consumed > 0 then
                let res = seq_rule (i + consumed) t in
                if res < 0 then
                  -1
                else
                  consumed + res
              else
                -1
  and
    match_rule i = function
    | Ast.CHAR chr ->
      if i >= candidate_len then
        -1
      else
        let c_chr = String.get candidate i in
        if c_chr = chr then
          1
        else
          -1
    | Ast.RULES seq -> seq_rule i seq
    | Ast.OR (r1, r2) -> begin
        let r1_i = seq_rule i r1 in
        if r1_i > 0 then
          r1_i
        else
          let r2_i = seq_rule i r2 in
          if r2_i > 0 then
            r2_i
          else
            -1
      end
    | _ -> begin
        Printf.printf "Unexpected case";
        exit 2
      end
  in
  let rule = Array.get rules 0 in
  match_rule 0 rule

let main () =
  let input = Parser.root Lexer.parse_token (Lexing.from_channel stdin) in
  let count = ref 0 in
  let rec eval = function
    | [] -> ()
    | e::t -> begin
        let consumed = candidate_match input.rules e in
        let len = String.length e in
        Printf.printf "%s %d\n" e consumed;
        if consumed = len then
          count := !count + 1;
        eval t
      end
  in
  eval input.candidates;
  Printf.printf "%d\n" !count

let _ = main ()
