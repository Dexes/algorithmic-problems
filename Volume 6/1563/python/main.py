result = 0
data = {}
for i in range(int(input())):
    row = input()
    if row in data:
        result += 1
    else:
        data[row] = True

print(result)
