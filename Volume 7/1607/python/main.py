a, b, c, d = input().split()
a = int(a)
b = int(b)
c = int(c)
d = int(d)

if a > c:
    print(a)
    exit()

while True:
    if a + b <= c:
        a += b
    else:
        print(c)
        exit()

    if c - d >= a:
        c -= d
    else:
        print(a)
        exit()
