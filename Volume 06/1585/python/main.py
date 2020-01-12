from collections import defaultdict

result_count = 0
result_name = ""
data = defaultdict(int)
for i in range(int(input())):
    row = input()
    data[row] += 1
    if data[row] > result_count:
        result_count = data[row]
        result_name = row

print(result_name)
