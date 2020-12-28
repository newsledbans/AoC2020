import fileinput
import re
import string

goodpass = 0
lines = list(fileinput.input("inputs/02.txt"))
for line in lines:
    minmax = re.findall(r'\d+', line)
    password = re.findall('[a-z]+', line)
    count = password[1].count(password[0])
    if  int(minmax[1]) >= count >= int(minmax[0]):
        goodpass += 1

print("Part1:",goodpass)