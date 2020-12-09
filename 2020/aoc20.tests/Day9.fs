module aoc20.tests.Day9

open Xunit
open aoc20.Day9

[<Fact>]
let ``Task 1 Example``() =
    let inp = [35;20;15;25;47;40;62;55;65;95;102;117;150;182;127;219;299;277;309;576] |> List.map (int64)
    let expected = 127L
    let actual = findNotSum 5 inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 29221323L
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let inp = [35;20;15;25;47;40;62;55;65;95;102;117;150;182;127;219;299;277;309;576] |> List.map (int64)
    let expected = 62 |> int64
    let invNum = findNotSum 5 inp
    let actual = getWeakness invNum inp
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 4389369L
    let actual = solveT2
    Assert.Equal(expected, actual)
