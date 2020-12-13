module aoc20.Day11

open Common

let checkInSeating seats x y =
    x >= 0
    && y >= 0
    && x < Array2D.length1 seats
    && y < Array2D.length2 seats

let checkSeatAround seats i j =
    [ for x in i - 1 .. i + 1 do
          for y in j - 1 .. j + 1 do x, y ]
    |> List.filter (fun (x,y) -> checkInSeating seats x y && (x <> i || y <> j))

let checkSeatInSight seats i j =
    let dirs =
        [ for x in -1 .. 1 do
            for y in -1 .. 1 do x,y ]
        |> List.filter (fun (x,y) -> not (x = 0 && y = 0))
    [ for (xdir, ydir) in dirs do
        let mutable x, y = i + xdir, j + ydir
        while checkInSeating seats x y && seats.[x,y] = '.' do
            x <- x + xdir
            y <- y + ydir
        if checkInSeating seats x y then yield x,y]

let checkIfSit validSeatFunc (seats:char[,]) i j =
    validSeatFunc seats i j
    |> List.exists (fun (x,y) -> seats.[x,y] = '#')
    |> not

let checkIfMove validSeatFunc (seats:char[,]) tol i j =
    validSeatFunc seats i j
    |> List.filter (fun (x,y) -> seats.[x,y] = '#')
    |> List.length >= tol

let doRound seats tol =
    Array2D.init (Array2D.length1 seats) (Array2D.length2 seats) (fun i j ->
        match seats.[i,j] with
        | '.' -> '.'
        | '#' -> if checkIfMove checkSeatAround seats tol i j then 'L' else '#'
        | 'L' -> if checkIfSit checkSeatAround seats i j then '#' else 'L'
        | _ -> failwith "Invalid character in seating description"
        )

let doInSight seats tol =
    Array2D.init (Array2D.length1 seats) (Array2D.length2 seats) (fun i j ->
        match seats.[i,j] with
        | '.' -> '.'
        | '#' -> if checkIfMove checkSeatInSight seats tol i j then 'L' else '#'
        | 'L' -> if checkIfSit checkSeatInSight seats i j then '#' else 'L'
        | _ -> failwith "Invalid character in seating description"
        )

let rec stabilise stableFunc seats tol =
    let seats' = stableFunc seats tol
    if isEqualArray2d seats seats' then seats else stabilise stableFunc seats' tol

let countSitters seats =
    array2DToSeq seats
    |> Seq.filter (fun x -> x = '#')
    |> Seq.length

let countOccupied stableFunc tol seats =
    stableFunc seats tol
    |> countSitters

let solveT1 =
    readLinesAsCharArrays("11.txt")
    |> arraysToArray2D
    |> countOccupied (stabilise doRound) 4

let solveT2 =
    readLinesAsCharArrays("11.txt")
    |> arraysToArray2D
    |> countOccupied (stabilise doInSight) 5
