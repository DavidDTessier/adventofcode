import re

with open('/Users/dtessier/Documents/GitHub/adventofcode/2020/Excerise 7 - Baggage/input/input.dat') as f:
    bagRegulations = [line.rstrip() for line in f]
	
# input:
# ------------
#['light red bags contain 1 bright white bag, 2 muted yellow bags.',
# 'dark orange bags contain 3 bright white bags, 4 muted yellow bags.',
# 'bright white bags contain 1 shiny gold bag.',
# 'muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.',
# 'shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.',
# 'dark olive bags contain 3 faded blue bags, 4 dotted black bags.',
# 'vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.',
# 'faded blue bags contain no other bags.',
# 'dotted black bags contain no other bags.']

# output:
# --------
# {'light red': [('1', 'bright white'), ('2', 'muted yellow')],
# 'dark orange': [('3', 'bright white'), ('4', 'muted yellow')],
# 'bright white': [('1', 'shiny gold')],
# 'muted yellow': [('2', 'shiny gold'), ('9', 'faded blue')],
# 'shiny gold': [('1', 'dark olive'), ('2', 'vibrant plum')],
# 'dark olive': [('3', 'faded blue'), ('4', 'dotted black')],
# 'vibrant plum': [('5', 'faded blue'), ('6', 'dotted black')],
# 'faded blue': [],
# 'dotted black': []}

def parseBagRegulations(bagRegulations):
    bagRules = {}
    for regulation in bagRegulations:
        regulation = re.sub('bag[s]?[\.]?', '', regulation)
        split = regulation.split("contain")
        tuples = []
        if(split[1].strip() == "no other"):
            bagRules[split[0].strip()] = []
            continue
        for node in split[1].strip().replace(' , ',',').split(','):
            tuples.append((node[:1],node[2:]))
        bagRules[split[0].strip()] = tuples

    return bagRules


def bagContainsColor(bag,coloredBag,parseBagRegulations):
    for subBag in parseBagRegulations[bag]:
        subBagCount = subBag[0]
        subBagColor = subBag[1]
        if subBagColor == coloredBag:
            return True
        elif bagContainsColor(subBagColor, coloredBag,parseBagRegulations):
            return True
    
    return False


def bagCount(bag,parsedBagRegulations):
    count = 0
    for subBag in parsedBagRegulations[bag]:
        subBagCount = int(subBag[0])
        subBagColor = subBag[1] 

        count += subBagCount + subBagCount*bagCount(subBagColor,parsedBagRegulations)

    return count    
    
def part1():
    parsedBagRegulations = parseBagRegulations(bagRegulations)
    desiredBag = 'shiny gold'
    count = 0
    for bag in parsedBagRegulations:
        if bagContainsColor(bag, desiredBag, parsedBagRegulations):
            count += 1

    print("Number of bags containing '{}' bags = {}".format(desiredBag,count))

def part2():
    parsedBagRegulations = parseBagRegulations(bagRegulations)  
    desiredBag = 'shiny gold'
    count = bagCount(desiredBag,parsedBagRegulations)
    
    print('{} bag must contain {} other bags'.format(desiredBag,count))
	

part1()
#part2()

# OUTPUT
# -------
# Number of bags containing 'shiny gold' bags = 128
# shiny gold bag must contain 20189 other bags