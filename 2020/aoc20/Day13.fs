module aoc20.Day13

open Common

let getEarliestTimesWait (inp:string[]) =
    let t = inp.[0] |> int
    let ids =
        inp.[1].Split(",")
        |> Array.filter (fun id -> id <> "x")
        |> Array.map (int)
    let waitTimes =
        ids
        |> Array.map (fun id -> id - (t % id))
    let sortedByWait = Array.zip waitTimes ids
                       |> Array.sortBy (fun (t',_) -> t')
    fst sortedByWait.[0] * snd sortedByWait.[0]

// Spent too long on this problem
// Shamelessly stole a better solution from https://www.reddit.com/r/adventofcode/comments/kc4njx/2020_day_13_solutions/gfny3uc
let matchingOffsetDeparture (depTime, period) (offset, id) =
    Seq.initInfinite (fun n -> (int64 n) * period + depTime)
    |> Seq.find (fun depTime -> (depTime + offset) % id = 0L)
    |> fun depTime -> depTime, period * id

let getEarliestTimestamp (inp:string) =
    let idDiffs =
        inp.Split(",")
        |> Array.indexed
        |> Array.filter (fun (_,v) -> v <> "x")
        |> Array.map (fun (i,v) -> int64 i, int64 v)
    Array.fold matchingOffsetDeparture (0L, 1L) idDiffs |> fst

let solveT1 =
    readLines("13.txt")
    |> Array.ofList
    |> getEarliestTimesWait

let solveT2 =
    let inp = readLines("13.txt")
              |> Array.ofList
    getEarliestTimestamp inp.[1]
