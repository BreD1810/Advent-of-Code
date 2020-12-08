module aoc20.Day8

open Common

let endLoopAcc (ins:string[]) =
    let rec execute lineNo (executed:int list) acc =
        if List.contains lineNo executed then
            acc, false
        else if lineNo = ins.Length then
            acc, true
        else
            let cIns = ins.[lineNo]
            let executed' = (executed@[lineNo])
            match cIns with
            | Regex @"nop\s[-,+]?\d+" [] -> execute (lineNo + 1) executed' acc
            | Regex @"acc\s([-,+]?\d+)" [n] -> execute (lineNo + 1) executed' (acc+(int n))
            | Regex @"jmp\s([-,+]?\d+)" [n]-> execute (lineNo + (int n)) executed' acc
            | _ -> failwithf "Unrecognised instruction: %s" cIns
    execute 0 [] 0

let swapInstruction (i:string) =
    let [|oc;oa|] = i.Split ' '
    match oc with
    | "nop" -> "jmp " + oa
    | "jmp" -> "nop " + oa
    | "acc" -> i
    | _ -> failwithf "Unrecognised instruction: %s" i

let changeToTerminate (ins:string[]) =
    let rec checkIns i =
        if i = ins.Length then 0
        else
            ins.[i] <- swapInstruction ins.[i]
            let acc, f = endLoopAcc ins
            if f then acc
            else
                ins.[i] <- swapInstruction ins.[i]
                checkIns (i+1)
    checkIns 0


let solveT1 =
    readLines("8.txt")
    |> Array.ofList
    |> endLoopAcc
    |> fst

let solveT2 =
    readLines("8.txt")
    |> Array.ofList
    |> changeToTerminate
