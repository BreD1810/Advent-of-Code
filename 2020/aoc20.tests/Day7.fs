module aoc20.tests.Day7

open Xunit
open aoc20.Day7

[<Fact>]
let ``Task 1 Example``() =
    let exampleRules = ["light red bags contain 1 bright white bag, 2 muted yellow bags."
                        "dark orange bags contain 3 bright white bags, 4 muted yellow bags."
                        "bright white bags contain 1 shiny gold bag."
                        "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags."
                        "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags."
                        "dark olive bags contain 3 faded blue bags, 4 dotted black bags."
                        "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."
                        "faded blue bags contain no other bags."
                        "dotted black bags contain no other bags."]
    let expected = 4
    let actual = noColoursContainingShinyGold exampleRules
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 1``() =
    let expected = 101
    let actual = solveT1
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2 Example``() =
    let exampleRules = ["shiny gold bags contain 2 dark red bags."
                        "dark red bags contain 2 dark orange bags."
                        "dark orange bags contain 2 dark yellow bags."
                        "dark yellow bags contain 2 dark green bags."
                        "dark green bags contain 2 dark blue bags."
                        "dark blue bags contain 2 dark violet bags."
                        "dark violet bags contain no other bags."]
    let expected = 126
    let actual = noBagsShinyGoldContains exampleRules
    Assert.Equal(expected, actual)

[<Fact>]
let ``Task 2``() =
    let expected = 108636
    let actual = solveT2
    Assert.Equal(expected, actual)
