module aoc20.Common

open System.IO
open System.Text.RegularExpressions

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

let readLinesAsInt64s (filePath:string) =
    seq {
        use sr = new StreamReader (filePath)
        while not sr.EndOfStream do
            yield sr.ReadLine ()
    } |> Seq.map int64 |> Seq.toList

let charToInt (c : char) = int c - int '0'

let readLinesAsCharArrays (filePath : string) =
    seq {
        use sr = new StreamReader (filePath)
        while not sr.EndOfStream do
            yield sr.ReadLine ()
    } |> Seq.map Seq.toArray |> Seq.toArray

let arraysToArray2D (arrays:'a[][]) =
    Array2D.init (Array.length arrays) (Array.length arrays.[0]) (fun i j -> arrays.[i].[j])

let array2DToSeq a =
    seq {
        for i = 0 to Array2D.length1 a - 1 do
            for j = 0 to Array2D.length2 a - 1 do
                yield a.[i,j]
    }

let isEqualArray2d a a' =
    Array2D.length1 a = Array2D.length1 a'
    && Array2D.length2 a = Array2D.length2 a'
    && seq {
        for i = 0 to Array2D.length1 a - 1 do
            for j = 0 to Array2D.length2 a - 1 do
                yield a.[i,j] = a'.[i,j]
    }
    |> Seq.reduce (fun b b' -> b && b')

let readAsString filepath = File.ReadAllText(filepath)

let (|Regex|_|) pattern input =
    let m = Regex.Match(input, pattern)
    if m.Success then Some(List.tail [for g in m.Groups -> g.Value])
    else None
