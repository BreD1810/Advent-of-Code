module aoc20.Day7

open System.Text.RegularExpressions

let parseRule (r:string) =
    let [|container;contents|] = r.Replace("bag","").Replace("bags","").Split("contain")
    let containerRe = Regex(@"^(?<bag>\w+\s\w+)",RegexOptions.Compiled)
    let bag = containerRe.Match(container).Groups.["bag"].Value
    let contentsSplits = contents.Split(",")
    let contentsRe = Regex(@"(?<count>\d+)\s(?<bag>\w+\s\w+)",RegexOptions.Compiled)
    seq {
        for contentsSplit in contentsSplits do
            if contentsRe.IsMatch(contentsSplit) then
                let m = contentsRe.Match(contentsSplit)
                let contentsBag = m.Groups.["bag"].Value
                let contentsCount = m.Groups.["count"].Value |> int
                match (contentsBag, bag) with
                | ("no other",_) -> ()
                | (_,"shiny gold") -> ()
                | _ -> yield (contentsBag, (bag,contentsCount))
    }

let getRules s =
    let mutable rules = Map.empty
    for (container, contents) in s do
        let inside =
            match rules.TryFind container with
            | Some bags -> bags
            | None -> Set.empty
        rules <- Map.add container (inside.Add contents) rules
    rules

let parseRule' (r:string) =
    let [|container;contents|] = r.Replace("bag","").Replace("bags","").Split("contain")
    let containerRe = Regex(@"^(?<bag>\w+\s\w+)",RegexOptions.Compiled)
    let bag = containerRe.Match(container).Groups.["bag"].Value
    let contentsSplits = contents.Split(",")
    let contentsRe = Regex(@"(?<count>\d+)\s(?<bag>\w+\s\w+)",RegexOptions.Compiled)
    bag, seq {
        for contentsSplit in contentsSplits do
            if contentsRe.IsMatch(contentsSplit) then
                let m = contentsRe.Match(contentsSplit)
                let contentsBag = m.Groups.["bag"].Value
                let contentsCount = m.Groups.["count"].Value |> int
                (contentsBag, contentsCount)
    }

let rec findAllBags (rules: Map<string, Set<string * int>>) bag =
    seq {
        let bags =
            match rules.TryFind(bag) with
            | None -> Set.empty
            | Some b -> b
        yield! bags
        for (bag,_) in bags do
            yield! findAllBags rules bag
    }

let rec countBags (rules: Map<string, seq<string * int>>) bag =
    let innerBags = Map.find bag rules
    match Seq.length innerBags with
    | 0 -> 0
    | _ ->
        innerBags
        |> Seq.map (fun (desc,n) -> n * ((countBags rules desc) + 1))
        |> Seq.sum

let noColoursContainingShinyGold rs =
    let rules = List.map parseRule rs
                |> Seq.concat
                |> getRules
    findAllBags rules "shiny gold" |> Seq.distinctBy fst |> Seq.length

let noBagsShinyGoldContains rs =
    let rules = List.map parseRule' rs |> Map.ofSeq
    countBags rules "shiny gold"

let solveT1 =
    noColoursContainingShinyGold (Common.readLines("7.txt"))

let solveT2 =
    noBagsShinyGoldContains (Common.readLines("7.txt"))

