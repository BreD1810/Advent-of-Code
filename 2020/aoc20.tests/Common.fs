module aoc20.tests.Common

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
