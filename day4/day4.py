#!/usr/bin/env python3

import fileinput
import re
import os

prefix = r"/Users/benwanless/Projects/AoC2020/"

def main():
    passCount = 0
    data = {}
    reHgt = re.compile('(^[\d]{2,3})[ic][nm]$')
    reHcl = re.compile('^#[0-9a-f]{6}$')
    colors = ("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
    rePid = re.compile('^[0-9]{9}$')

    for line in fileinput.input(prefix + "inputs/04.txt"):
        line = line.strip()
        if line:
            for description in line.split():
                    k, v = description.split(':')
                    data[k]= v
        else:
            if checkdata(data):
                passCount +=1
            data.clear()
    if checkdata(data):
        #  need to call this one more time for the EOF
        passCount +=1
    print(passCount)

def checkdata(data):
    # cid (Country ID) - ignored, missing or not.
    if 'cid' in data:
        del data['cid']
    
    if len(data) == 7:
        return True
    return False

if __name__ == '__main__':
    main()
    