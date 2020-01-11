import math

v, a, k = input().split()
v = int(v)
a = int(a)
k = float(k)

result = v * v * k * math.sin(a * math.pi / 90) / (10 * (k - 1))
print('%.2f' % result)
