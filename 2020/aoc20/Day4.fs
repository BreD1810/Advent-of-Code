module aoc20.Day4

open System
open System.Text.RegularExpressions

let removeUnit (s : string) = s.Substring(0, s.Length - 2)

let fieldValidators =
    [ "byr", (fun bys -> let by = int bys in by >= 1920 && by <= 2002)
      "iyr", (fun iys -> let iy = int iys in iy >= 2010 && iy <= 2020)
      "eyr", (fun eys -> let ey = int eys in ey >= 2020 && ey <= 2030)
      "hgt", (fun (hgts : string) -> hgts.EndsWith("cm") && (let hgt = int (removeUnit hgts) in hgt >= 150 && hgt <= 193)
                                     || hgts.EndsWith("in") && (let hgt = int (removeUnit hgts) in hgt >= 59 && hgt <= 76))
      "hcl", (fun hcls -> Regex.IsMatch(hcls, "^#[0-9,a-f]{6}$"))
      "ecl", (fun ecls -> Regex.IsMatch(ecls, "amb|blu|brn|gry|grn|hzl|oth"))
      "pid", (fun pids -> Regex.IsMatch(pids, "^[0-9]{9}$"))]

let parsePassports (ps:string) =
    let ps = ps.Replace(Environment.NewLine, " ")
    let getFieldVal (s:string) =
        let [|f;v|] = s.Split(':')
        f,v
    ps.Split(' ')
    |> Array.map getFieldVal
    |> Map.ofArray

let countValidPassports (lines : string) =
    let requiredFields = List.map fst fieldValidators
    lines.Split("\n\n")
    |> Array.map parsePassports
    |> Array.filter (fun pfv -> List.forall (fun pv -> pfv.ContainsKey pv) requiredFields)
    |> Array.length

let applyRule (ps : Map<string,string>) (f,fnc) =
    match ps.TryFind f with
    | Some s -> fnc s
    | None -> failwith "Couldn't find"

let countValidPassportsFieldValidation (lines : string) =
    let requiredFields = List.map fst fieldValidators
    lines.Split("\n\n")
    |> Array.map parsePassports
    |> Array.filter (fun ps -> List.forall (fun f -> ps.ContainsKey f) requiredFields)
    |> Array.filter (fun ps -> List.forall (fun v -> applyRule ps v) fieldValidators)
    |> Array.length

let solveT1 =
    countValidPassports (Common.readAsString("4.txt"))

let solveT2 =
    countValidPassportsFieldValidation (Common.readAsString("4.txt"))
