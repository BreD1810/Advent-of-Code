module aoc20.tests.Day6

open Xunit
open aoc20.Day6

[<Fact>]
let ``Task 1 Example``() =
    let ls = """abc

a
b
c

ab
ac

a
a
a
a

b"""
    let expected = 11
    let actual = sumAnswerCounts ls
    Assert.Equal(expected,actual)

[<Fact>]
let ``Task 1``() =
    let expected = 6273
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let ls = """abc

a
b
c

ab
ac

a
a
a
a

b"""
    let expected = 6
    let actual = sumAnswerCountsAllSame ls
    Assert.Equal(expected,actual)

[<Fact>]
let ``Task 2``() =
    let expected = 3254
    let actual = solveT2
    Assert.Equal(expected, actual)
