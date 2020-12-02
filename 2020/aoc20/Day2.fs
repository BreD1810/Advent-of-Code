module aoc20.Day2

open Common

type entry = {min:int; max:int; c:string; s:string}

let getSingleInputInfo (entry : string) =
    let [| rule; password |] = entry.Split(": ")
    let [| range; c |] = rule.Split(' ')
    { min = range.Split('-').[0] |> int;
      max = range.Split('-').[1] |> int;
      c = c.TrimEnd(':');
      s = password}

let checkCountInRange entry =
    let count = entry.s |> Seq.filter (fun c' -> c' = char entry.c) |> Seq.length
    count >= entry.min && count <= entry.max

let checkPositions entry =
    let x = entry.s.[entry.min - 1] = char entry.c
    let y = entry.s.[entry.max - 1] = char entry.c
    (x || y) && not(x && y)

let countValidPasswords entries =
    List.map getSingleInputInfo entries
    |> List.filter checkCountInRange
    |> List.length

let countValidPasswords2 entries =
    List.map getSingleInputInfo entries
    |> List.filter checkPositions
    |> List.length

let solveT1 =
    countValidPasswords (readLines("2.txt"))

let solveT2 =
    countValidPasswords2 (readLines("2.txt"))
