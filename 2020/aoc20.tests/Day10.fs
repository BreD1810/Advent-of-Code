module aoc20.tests.Day10

open Xunit
open aoc20.Day10

[<Fact>]
let ``Task 1 Example 1``() =
    let inp = [16;10;15;5;1;11;7;19;6;12;4]
    let expected = 35
    let actual = productDiffs inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1 Example 2``() =
    let inp = [28;33;18;42;31;14;46;20;48;47;24;23;49;45;19;38;39;11;1;32;25;35;8;17;7;9;4;2;34;10;3]
    let expected = 220
    let actual = productDiffs inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 2760
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 1``() =
    let inp = [16;10;15;5;1;11;7;19;6;12;4]
    let expected = 8L
    let actual = countCombos inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 2``() =
    let inp = [28;33;18;42;31;14;46;20;48;47;24;23;49;45;19;38;39;11;1;32;25;35;8;17;7;9;4;2;34;10;3]
    let expected = 19208L
    let actual = countCombos inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 13816758796288L
    let actual = solveT2
    Assert.Equal(expected, actual)
