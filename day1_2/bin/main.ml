let maybe_read_line () = try Some (read_line ()) with End_of_file -> None

let rec read_stdin_aux lines =
  match maybe_read_line () with
  | Some line -> read_stdin_aux (line :: lines)
  | None -> lines

let read_stdin = 
    read_stdin_aux []

let get_digit_char c = match c with
  | '0' .. '9' -> String.make 1 c
  | _ -> ""


let text_digits = [("one", 1); ("two", 2); ("three", 3); ("four", 4); ("five", 5); ("six", 6); ("seven", 7); ("eight", 8); ("nine", 9)]

(*
let sub_text_to_digit line =
    List.concat_map
*)
    



let find_cal_value line =
    let digit_str = String.fold_left (fun acc el -> acc ^ (get_digit_char el)) "" line in
    let cal_value = Format.sprintf "%c%c" digit_str.[0] digit_str.[String.length digit_str - 1] in
    let out = int_of_string @@ cal_value in
    out

let () = 
    read_stdin |> List.map find_cal_value |> List.fold_left (fun acc el -> acc + el) 0 |> print_int;
    print_newline ()
