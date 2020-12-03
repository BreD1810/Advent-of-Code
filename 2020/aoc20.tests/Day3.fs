module aoc20.tests.Day3

open System.Numerics
open Xunit
open aoc20.Day3

[<Fact>]
let ``Task 1 Example``() =
    let ls = [| "..##......."
                "#...#...#.."
                ".#....#..#."
                "..#.#...#.#"
                ".#...##..#."
                "..#.##....."
                ".#.#.#....#"
                ".#........#"
                "#.##...#..."
                "#...##....#"
                ".#..#...#.#" |]
    let input = Array.map (fun l -> Seq.toArray l) ls
    let expected = 7
    let w = 11
    let h = 11
    let actual = right3Down1 0 0 w h input
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 247
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let ls = [| "..##......."
                "#...#...#.."
                ".#....#..#."
                "..#.#...#.#"
                ".#...##..#."
                "..#.##....."
                ".#.#.#....#"
                ".#........#"
                "#.##...#..."
                "#...##....#"
                ".#..#...#.#" |]
    let input = Array.map (fun l -> Seq.toArray l) ls
    let expected = 336 |> bigint
    let w = 11
    let h = 11
    let s1 = rightrDownd 0 0 w h input 1 1 |> bigint
    let s2 = rightrDownd 0 0 w h input 3 1 |> bigint
    let s3 = rightrDownd 0 0 w h input 5 1 |> bigint
    let s4 = rightrDownd 0 0 w h input 7 1 |> bigint
    let s5 = rightrDownd 0 0 w h input 1 2 |> bigint
    let result = List.fold (*) (bigint 1) [ s1; s2; s3; s4; s5 ]
    Assert.Equal(expected, result)

[<Fact>]
let ``Task 2``() =
    let expected = BigInteger(2983070376L)
    let actual = solveT2
    Assert.Equal(expected, actual)
