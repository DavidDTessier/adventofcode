#!/usr/bin/env python

import re

with open('./input', 'r') as input:
    passport_data = [l.strip() for l in input.readlines()]

def parse_passports(passports):
    passport = []
    for line in passports + ['']:
        if line == '' and len(passport) > 0:
            yield ' '.join(passport).strip()
            passport = []
        passport.append(line)

def part1():
    passports = parse_passports(passport_data)
    valid_count = 0
    for passport in passports:
        parts = passport.split()
        parts_count = len(parts)
        if ((parts_count == 8) or
           (parts_count == 7 and all(map(lambda s: not s.startswith('cid'), parts)))):
            valid_count = valid_count + 1
    print(f'Number of valid passports = {valid_count}')

def part2():
    valid_byr = lambda byr: (byr != None and int(byr) in range(1920, 2002+1)) 
    valid_iyr = lambda iyr: (iyr != None and int(iyr) in range(2010, 2020+1))
    valid_eyr = lambda eyr: (eyr != None and int(eyr) in range(2020, 2030+1))
    valid_hgt = lambda hgt: (hgt != None and re.match('^(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)$', hgt))
    valid_hcl = lambda hcl: (hcl != None and re.match('^#[0-9a-f]{6}$', hcl))
    valid_ecl = lambda ecl: (ecl != None and ecl in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'])
    valid_pid = lambda pid: (pid != None and re.match('^[0-9]{9}$', pid))

    def validate_parts(byr=None, iyr=None, eyr=None, hgt=None, hcl=None, ecl=None, pid=None, cid=None):
        return (
            valid_byr(byr) and
            valid_iyr(iyr) and
            valid_eyr(eyr) and
            valid_hgt(hgt) and
            valid_hcl(hcl) and
            valid_ecl(ecl) and
            valid_pid(pid)
        )
    passports = parse_passports(passport_data)
    valid_count = 0
    for passport in passports:
        fields = {x[0]: x[1] for field in passport.split() for x in [field.split(':')]}
        if validate_parts(**fields):
            valid_count = valid_count + 1
    print(f'Number of valid passports = {valid_count}')

part1()
part2()