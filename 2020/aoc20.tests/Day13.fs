module aoc20.tests.Day13

open Xunit
open aoc20.Day13

[<Fact>]
let ``Task 1 Example``() =
    let inp = [|"939";"7,13,x,x,59,x,31,19"|]
    let actual = getEarliestTimesWait inp
    let expected = 295
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 119
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 1``() =
    let inp = "17,x,13,19"
    let expected = 3417L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 2``() =
    let inp = "67,7,59,61"
    let expected = 754018L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 3``() =
    let inp = "67,x,7,59,61"
    let expected = 779210L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 4``() =
    let inp = "67,7,x,59,61"
    let expected = 1261476L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 5``() =
    let inp = "1789,37,47,1889"
    let expected = 1202161486L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example 6``() =
    let inp = "7,13,x,x,59,x,31,19"
    let expected = 1068781L
    let actual = getEarliestTimestamp inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 1106724616194525L
    let actual = solveT2
    Assert.Equal(expected, actual)
