module aoc20.Day5

let getRowNo (ri:char[]) =
    let rec getRowNo' ri' n n' = match ri' with
    | d::[] -> if d = 'F' then n else n'
    | d::ds -> if d = 'F' then getRowNo' ds n (n + (n'-n)/2)
               else getRowNo' ds (n + (n'-n)/2 + 1) n'
    in getRowNo' (Array.toList ri) 0 127

let getSeatNo (si:char[]) =
    let rec getSeatNo' si' n n' = match si' with
    | d::[] -> if d = 'L' then n else n'
    | d::ds -> if d = 'L' then getSeatNo' ds n (n + (n'-n)/2)
               else getSeatNo' ds (n + (n'-n)/2 + 1) n'
    in getSeatNo' (Array.toList si) 0 7

let getSeatNoFromBP (bp:char[]) =
    let rowInfo = Array.take 7 bp
    let seatInfo = Array.rev bp |> Array.take 3 |> Array.rev
    let rowNo = getRowNo rowInfo
    let seatNo = getSeatNo seatInfo
    rowNo * 8 + seatNo

let getMaxSeatNo bps =
    Array.map getSeatNoFromBP bps
    |> Array.max

let getMinSeatNo bps =
    Array.map getSeatNoFromBP bps
    |> Array.min

let solveT1 =
    getMaxSeatNo (Common.readLinesAsCharArrays("5.txt"))

let getSeatIdFromNo (seatNo:int) (seats:int list) (bps:char[][]) =
    match List.tryFindIndex (fun s -> s = seatNo) seats with
    | Some i -> i
    | None -> failwith "Couldn't find seat"

let solveT2=
    let bps = (Common.readLinesAsCharArrays("5.txt"))
    let max = getMaxSeatNo bps
    let min = getMinSeatNo bps
    let seats = [for bp in bps do getSeatNoFromBP bp]
    let possibleSeats = [min..max]
    let [seatNo] = List.except seats possibleSeats
    seatNo

