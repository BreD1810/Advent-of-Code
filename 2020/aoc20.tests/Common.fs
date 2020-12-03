module aoc20.tests.Common

open System
open aoc20.Common
open Xunit

[<Fact>]
let ``read first line``() =
    let fileContents = readLines "1.txt"
    Assert.Equal("1470", Seq.head fileContents)

let ``read first line as int``() =
    let fileContents = readLinesAsInts "1.txt"
    Assert.Equal(1470, Seq.head fileContents)

[<Fact>]
let ``file length test`` () =
    let fileContents = readLines "1.txt"
    Assert.Equal(200, Seq.length fileContents)

[<Fact>]
let ``read first line as char array``() =
    let line1 = (readLinesAsCharArrays "1.txt").[0]
    Assert.Equal('1', line1.[0])
    Assert.Equal('4', line1.[1])
    Assert.Equal('7', line1.[2])
    Assert.Equal('0', line1.[3])
    Assert.Throws<IndexOutOfRangeException>(fun () -> line1.[4] |> ignore)
