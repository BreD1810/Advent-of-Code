module aoc20.tests.Day2

open Xunit
open aoc20.Day2

[<Fact>]
let ``Task 1 Example``() =
    let input = ["1-3 a: abcde"; "1-3 b: cdefg"; "2-9 c: ccccccccc"]
    let expected = 2
    let actual = countValidPasswords input
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 474
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let input = ["1-3 a: abcde"; "1-3 b: cdefg"; "2-9 c: ccccccccc"]
    let expected = 1
    let actual = countValidPasswords2 input
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 745
    let actual = solveT2
    Assert.Equal(expected, actual)
