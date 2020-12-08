module aoc20.tests.Day8

open Xunit
open aoc20.Day8

[<Fact>]
let ``Task 1 Example 1``() =
    let ins = [|"nop +0"
                "acc +1"
                "jmp +4"
                "acc +3"
                "jmp -3"
                "acc -99"
                "acc +1"
                "jmp -4"
                "acc +6"|]
    let expected = 5
    let actual = endLoopAcc ins |> fst
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 2025
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let ins = [|"nop +0"
                "acc +1"
                "jmp +4"
                "acc +3"
                "jmp -3"
                "acc -99"
                "acc +1"
                "jmp -4"
                "acc +6"|]
    let expected = 8
    let actual = changeToTerminate ins
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 2001
    let actual = solveT2
    Assert.Equal(expected, actual)
