module aoc20.tests.Day5

open Xunit
open Xunit.Abstractions
open aoc20.Day5

[<Fact>]
let ``Task 1 Example 0``() =
    let expected = 357
    let bp = "FBFBBFFRLR".ToCharArray(0,10)
    let actual = getSeatNoFromBP bp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1 Example 1``() =
    let expected = 567
    let bp = "BFFFBBFRRR".ToCharArray(0,10)
    let actual = getSeatNoFromBP bp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1 Example 2``() =
    let expected = 119
    let bp = "FFFBBBFRRR".ToCharArray(0,10)
    let actual = getSeatNoFromBP bp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1 Example 3``() =
    let expected = 820
    let bp = "BBFFBBFRLL".ToCharArray(0,10)
    let actual = getSeatNoFromBP bp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 976
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 685
    let actual = solveT2
    Assert.Equal(expected, actual)
