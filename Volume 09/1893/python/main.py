def premium(letter):
    if letter == 'A' or letter == 'D':
        return 'window'

    return 'aisle'


def business(letter):
    if letter == 'A' or letter == 'F':
        return 'window'

    return 'aisle'


def economy(letter):
    if letter == 'A' or letter == 'K':
        return 'window'

    if letter == 'B' or letter == 'E' or letter == 'F' or letter == 'J':
        return 'neither'

    return 'aisle'


place = input()
row, letter = int(place[:-1]), place[-1]

if row < 3:
    print(premium(letter))
    exit()

if row < 21:
    print(business(letter))
    exit()

print(economy(letter))
