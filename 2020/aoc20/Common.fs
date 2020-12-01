module aoc20.Common

open System.IO

let readLines (filePath:string) =
    seq {
        use sr = new StreamReader (filePath)
        while not sr.EndOfStream do
            yield sr.ReadLine ()
    } |> Seq.toList

let readLinesAsInts (filePath:string) =
    seq {
        use sr = new StreamReader (filePath)
        while not sr.EndOfStream do
            yield sr.ReadLine ()
    } |> Seq.map int |> Seq.toList
