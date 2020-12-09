module aoc20.Day9

open Common

let makeSums ns =
    [for i in ns do
     for j in ns do
     if i <> j then yield i+j]
    |> List.distinct

let findNotSum preambleSz (inp:int64 list) =
    let isValid i =
        let possibleSums = inp.[i - preambleSz .. i - 1] |> makeSums
        if possibleSums |> List.contains inp.[i] then None else Some(inp.[i])
    [preambleSz .. inp.Length - 1]
    |> List.choose isValid
    |> List.head

let rec getWeakness invNum (inp:int64 list) = match inp with
    | [] -> failwith "Sum not found"
    | is ->
        let rec sumVals acc curs =
            if acc = invNum then
                let window = [for i=0 to curs - 1 do yield is.[i]]
                (List.max window, List.min window)
            else if acc > invNum then (0L,0L)
            else sumVals (acc + is.[curs]) (curs+1)
        match sumVals 0L 0 with
        | (0L,0L) -> getWeakness invNum is.Tail
        | (r,r') -> r + r'

let solveT1 =
    readLinesAsInt64s("9.txt")
    |> findNotSum 25

let solveT2 =
    let inp = readLinesAsInt64s("9.txt")
    let invNum = findNotSum 25 inp
    getWeakness invNum inp
