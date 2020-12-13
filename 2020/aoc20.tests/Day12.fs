module aoc20.tests.Day12

open Xunit
open aoc20.Day12

[<Fact>]
let ``Task 1 Example``() =
    let finalPos =
        ["F10";"N3";"F7";"R90";"F11"]
        |> List.map parseInstruction
        |> List.fold doInstruction {north = 0; east = 0; direction = 90;}
    let actual = manhattanDistance finalPos.east finalPos.north
    let expected = 25
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 1601
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let expected = 286
    let finalPos =
        ["F10";"N3";"F7";"R90";"F11"]
        |> List.map parseInstruction
        |> List.fold doInstructionWithWaypoint {north = 0; east = 0; waypointNorth = 1; waypointEast = 10;}
    let actual = manhattanDistance finalPos.east finalPos.north
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 13340
    let actual = solveT2
    Assert.Equal(expected, actual)
