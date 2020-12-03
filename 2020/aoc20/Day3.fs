module aoc20.Day3

let rec right3Down1 x y w h (lines : char[][]) =
    if y = h - 1 then
        0
    else
        let x' = (x + 3) % w
        let y' = y + 1
        if lines.[y'].[x'] = '#' then
            1 + right3Down1 x' y' w h lines
        else
            right3Down1 x' y' w h lines

let rec rightrDownd x y w h (lines : char[][]) r d =
    if y >= h - d then
        0
    else
        let x' = (x + r) % w
        let y' = y + d
        if lines.[y'].[x'] = '#' then
            1 + rightrDownd x' y' w h lines r d
        else
            rightrDownd x' y' w h lines r d

let solveT1 =
    let ls = Common.readLinesAsCharArrays "3.txt"
    let w = ls.[0].Length
    let h = ls.Length
    right3Down1 0 0 w h ls

let solveT2 =
    let ls = Common.readLinesAsCharArrays "3.txt"
    let w = ls.[0].Length
    let h = ls.Length
    let s1 = rightrDownd 0 0 w h ls 1 1 |> bigint
    let s2 = rightrDownd 0 0 w h ls 3 1 |> bigint
    let s3 = rightrDownd 0 0 w h ls 5 1 |> bigint
    let s4 = rightrDownd 0 0 w h ls 7 1 |> bigint
    let s5 = rightrDownd 0 0 w h ls 1 2 |> bigint
    List.fold (*) (bigint 1) [ s1; s2; s3; s4; s5 ]
