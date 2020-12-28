import sys
import fileinput

X = [int(line) for line in fileinput.input("inputs/01.txt")]

done = False
for i, val1 in enumerate(X):
    for j, val2 in enumerate(X[i:]):
        if val1 + val2 == 2020:
            print('part1', val1*val2)
        for val3 in X[j:]:
            if val1 + val2 + val3 == 2020:
                print('part2 ', val1*val2*val3)
                done = True