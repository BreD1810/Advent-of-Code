module aoc20.Day12

open Common

type Instruction =
    | North of int
    | South of int
    | East of int
    | West of int
    | Left of int
    | Right of int
    | Forward of int

type Position =
    { north: int
      east: int
      direction: int }

type PositionWithWaypoint =
    { north: int
      east: int
      waypointNorth: int
      waypointEast: int }

let parseInstruction (ins:string) =
    let n = ins.Substring 1 |> int
    match ins.[0] with
    | 'N' -> North n
    | 'E' -> East n
    | 'S' -> South n
    | 'W' -> West n
    | 'L' -> Left n
    | 'R' -> Right n
    | 'F' -> Forward n
    | _ -> failwith "Invalid instruction provided"

let doInstruction (curPos:Position) ins =
    match ins with
    | North x -> { curPos with north = curPos.north + x }
    | East x -> { curPos with east = curPos.east + x }
    | South x -> { curPos with north = curPos.north - x }
    | West x -> { curPos with east = curPos.east - x }
    | Left x -> { curPos with direction = (curPos.direction + 360 - x) % 360 }
    | Right x -> { curPos with direction = (curPos.direction + x) % 360 }
    | Forward x ->
        match curPos.direction with
        | 0 -> { curPos with north = curPos.north + x }
        | 90 -> { curPos with east = curPos.east + x }
        | 180 -> { curPos with north = curPos.north - x }
        | 270 -> { curPos with east = curPos.east - x }
        | d -> failwithf "Invalid direction %i" d

let doInstructionWithWaypoint (curPos:PositionWithWaypoint) ins =
    let rotate90 pos =
        { pos with
              waypointEast = pos.waypointNorth
              waypointNorth = -pos.waypointEast }
    let rec rotate =
        function
        | 0 -> id
        | 90 -> rotate90
        | 180 -> rotate90 >> rotate90
        | 270 -> rotate90 >> rotate90 >> rotate90
        | d ->
            if d >= 360 then rotate (d % 360)
            elif d < 0 then rotate (360 + d)
            else failwith "Invalid direction"
    match ins with
    | North x -> { curPos with waypointNorth = curPos.waypointNorth + x }
    | South x -> { curPos with waypointNorth = curPos.waypointNorth - x }
    | East x -> { curPos with waypointEast = curPos.waypointEast + x }
    | West x -> { curPos with waypointEast = curPos.waypointEast - x }
    | Left x -> rotate -x curPos
    | Right x -> rotate x curPos
    | Forward x -> {curPos with
                        north = curPos.north + (curPos.waypointNorth * x)
                        east = curPos.east + (curPos.waypointEast * x)}

let manhattanDistance x y = abs x + abs y

let solveT1 =
    let finalPos =
        readLines("12.txt")
        |> List.map parseInstruction
        |> List.fold doInstruction {north = 0; east = 0; direction = 90;}
    manhattanDistance finalPos.east finalPos.north

let solveT2 =
    let finalPos =
        readLines("12.txt")
        |> List.map parseInstruction
        |> List.fold doInstructionWithWaypoint {north = 0; east = 0; waypointNorth = 1; waypointEast = 10;}
    manhattanDistance finalPos.east finalPos.north
