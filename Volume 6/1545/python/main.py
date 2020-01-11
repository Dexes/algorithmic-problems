from collections import defaultdict

data = defaultdict(list)
for i in range(int(input())):
    h = input()
    data[h[0]].append(h)

for h in data[input()]:
    print(h)
