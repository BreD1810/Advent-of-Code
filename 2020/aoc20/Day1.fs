module aoc20.Day1

let rec makeCombinations (entries:list<int>) = match entries with
    | [] -> []
    | x :: xs ->
        let xps = List.map (fun y -> (x,y)) xs
        xps @ (makeCombinations xs)

let rec make3Combinations (entries:list<int>) = match entries with
    | [] -> []
    | x :: xs ->
        let pairs = makeCombinations xs
        let xps = List.map (fun (y,z) -> (x,y,z)) pairs
        xps @ (make3Combinations xs)

let rec sum2020Multiply entries =
    makeCombinations entries |> List.find (fun (x,y) -> x + y = 2020) |> (fun (x,y) -> x * y)

let rec sum3Multiply entries =
    make3Combinations entries |> List.find (fun (x,y,z) -> x + y + z = 2020) |> (fun(x,y,z) -> x * y * z)

let solveT1 =
    sum2020Multiply (Common.readLinesAsInts("1.txt"))

let solveT2 =
    sum3Multiply (Common.readLinesAsInts("1.txt"))
