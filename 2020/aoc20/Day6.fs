module aoc20.Day6

let countAnswers (ans:string) =
    ans.ToCharArray(0, ans.Length)
    |> Array.distinct
    |> Array.length

let sumAnswerCounts (anss:string) =
    anss.Split("\n\n")
    |> Array.map(fun s -> s.Replace("\n", ""))
    |> Array.map countAnswers
    |> Array.fold (+) 0

let countAnswersAllSame (ans:string) =
    let noPeople = ans.Split('\n').Length
    let ans = ans.Replace("\n", "")
    ans.ToCharArray(0,ans.Length)
    |> Array.countBy(fun c -> c)
    |> Array.filter (fun (c,n) -> n = noPeople)
    |> Array.length

let sumAnswerCountsAllSame (anss:string) =
    anss.Split("\n\n")
    |> Array.map countAnswersAllSame
    |> Array.fold (+) 0

let solveT1 =
    sumAnswerCounts (Common.readAsString("6.txt"))

let solveT2 =
    sumAnswerCountsAllSame (Common.readAsString("6.txt"))
