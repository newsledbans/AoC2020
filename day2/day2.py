import fileinput
import re
import string

part1, part2 = 0, 0
lines = list(fileinput.input("inputs/02.txt"))
for line in lines:
    minmax = re.findall(r'\d+', line)
    password = re.findall('[a-z]+', line)
    count = password[1].count(password[0])
    if  int(minmax[1]) >= count >= int(minmax[0]):
        part1 += 1
    pos1 = password[1][int(minmax[0])-1] == password[0]
    pos2 = password[1][int(minmax[1])-1] == password[0]
    if pos1 ^ pos2: #(pos1 and not pos2) or (pos2 and not pos1):
        part2 += 1


print("Part1:", part1)
print("Part2:", part2)