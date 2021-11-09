data = input().split()
n, k = int(data[0]), len(data[1])

result = n
factor = n - k
iteration = 1

while factor > 0:
    result *= factor

    iteration += 1
    factor = n - k * iteration

print(result)
