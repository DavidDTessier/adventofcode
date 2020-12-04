#!/usr/bin/env python

import operator
from functools import reduce

with open('./Excercise 3 - Toboggan Trajectory/input/day3input.txt', 'r') as input:
    forest = [l.strip() for l in input.readlines()]

num_columns = len(forest[0])

def part1():
    slope = 3
    pos = 0
    tree_count = 0
    for row in forest:
        if row[pos] == '#':
            tree_count = tree_count + 1
        pos = (pos + slope) % num_columns
    print(f'Number of trees = {tree_count}')

def part2():
    slopes = (
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2)
    )
    def tree_count(run, rise):
        pos = 0
        tree_count = 0
        for row in forest[0::rise]:
            if row[pos] == '#':
                tree_count = tree_count + 1
            pos = (pos + run) % num_columns
        return tree_count
    result = reduce(operator.mul, (tree_count(*slope) for slope in slopes))
    print(f'Part 2 Result = {result}')

part1()
part2()