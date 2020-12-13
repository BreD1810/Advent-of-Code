module aoc20.tests.Day11

open Xunit
open aoc20.Day11
open aoc20.Common

[<Fact>]
let ``Task 1 Example``() =
    let inp = [|[|'L';'.';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'L';'L';'L';'L';'L';'L';'.';'L';'L'|];
              [|'L';'.';'L';'.';'L';'.';'.';'L';'.';'.'|];
              [|'L';'L';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'.';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'.';'L';'L';|];
              [|'.';'.';'L';'.';'L';'.';'.';'.';'.';'.';|];
              [|'L';'L';'L';'L';'L';'L';'L';'L';'L';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'L';'.';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'.';'L';'L';|]|]
              |> arraysToArray2D
    let expected = 37
    let actual = countOccupied (stabilise doRound) 4 inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 2126
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let inp = [|[|'L';'.';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'L';'L';'L';'L';'L';'L';'.';'L';'L'|];
              [|'L';'.';'L';'.';'L';'.';'.';'L';'.';'.'|];
              [|'L';'L';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'.';'L';'L';'.';'L';'L';'.';'L';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'.';'L';'L';|];
              [|'.';'.';'L';'.';'L';'.';'.';'.';'.';'.';|];
              [|'L';'L';'L';'L';'L';'L';'L';'L';'L';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'L';'.';'L';|];
              [|'L';'.';'L';'L';'L';'L';'L';'.';'L';'L';|]|]
              |> arraysToArray2D
    let expected = 26
    let actual = countOccupied (stabilise doInSight) 5 inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 1914
    let actual = solveT2
    Assert.Equal(expected, actual)
