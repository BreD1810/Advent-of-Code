module aoc20.tests.Day1

open aoc20.Day1
open Xunit

[<Fact>]
let ``Task 1 Example``() =
    let expected = 514579
    let input = [1721;979;366;299;675;1456]
    let actual = sum2020Multiply input
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let expected = 241861950
    let input = [1721;979;366;299;675;1456]
    let actual = sum3Multiply input
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 618144
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 173538720
    let actual = solveT2
    Assert.Equal(expected, actual)
