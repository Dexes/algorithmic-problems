import math


def get_sum(x, n):
    if x < 10:
        x += 0.05 - 1e-12

    return math.floor(x * n)


def get_voices_count(x, y, n):
    y += 0.05 - 1e-12

    if x <= y:
        return 0

    s = get_sum(x, n)
    return math.ceil((s - y * n) / (y - 1))


data = input().split(' ')
print(get_voices_count(float(data[0]), float(data[1]), int(data[2])))
