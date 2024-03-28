// # tree
// . snow
const getInput = (): string => {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

enum Thing {
  Tree,
  Snow
}

const things = getInput().
  split("\n").
  map(x => x.split("").
    map(x => x === "." ? Thing.Snow : Thing.Tree))
//
const colLen = things[0].length;
//
let treeCount = 0;
//

for (let i = 0; i < things.length; i++) {
  const row = things[i];
  if (row[i * 3 % colLen] === Thing.Tree) {// over 3 from i 
    treeCount++;
  }

}

console.log("tree count: ", treeCount)
