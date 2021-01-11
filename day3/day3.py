#!/usr/bin/env python3

import fileinput
import os

prefix = r"/Users/benwanless/Projects/AoC2020/"

slopes = [(1,1),(3,1),(5,1),(7,1),(1,2)]
treetotals = []

lines = [line.strip() for line in fileinput.input(prefix + "inputs/03.txt")]
for slope in slopes:
    hits, xpos = 0,0
    for row, line in enumerate(lines):
        if row % slope[1] != 0 :
            continue
        if line[xpos] == '#':
            hits += 1
        xpos = (xpos + slope[0])%len(line)
    treetotals.append(hits)

print("Part1:", treetotals[1])
part2 = 1
for tree in treetotals:
    part2 *= tree
print("Part2:", part2)
