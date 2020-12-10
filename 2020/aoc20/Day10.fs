module aoc20.Day10

open Common

let productDiffs adapters =
    List.append adapters [0;(List.max adapters) + 3]
    |> List.sort
    |> List.pairwise
    |> List.countBy (fun (x,y) -> y - x)
    |> List.fold (fun acc (x,v) -> if x = 3 || x = 1 then acc * v else acc) 1

let countCombos adapters =
    let diffList = List.append adapters [0;(List.max adapters) + 3]
                   |> List.sort
                   |> List.pairwise
                   |> List.map (fun (x,y) -> y - x)
                   |> List.map string
                   |> String.concat ""
    let groups = diffList.Split('3')
    let groupsOf2 = groups |> Array.filter (fun g -> g.Length = 2) |> Array.length |> float
    let groupsOf3 = groups |> Array.filter (fun g -> g.Length = 3) |> Array.length |> float
    let groupsOf4 = groups |> Array.filter (fun g -> g.Length = 4) |> Array.length |> float
    (2.0 ** groupsOf2 * 4.0 ** groupsOf3 * 7.0 ** groupsOf4) |> int64

let solveT1 =
    readLinesAsInts("10.txt")
    |> productDiffs

let solveT2 =
    readLinesAsInts("10.txt")
    |> countCombos
