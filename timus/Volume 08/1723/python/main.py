data = [0] * 26
result_count = 0
result_symbol = ""
for symbol in input():
    index = ord(symbol) - 97
    data[index] += 1
    if data[index] > result_count:
        result_count = data[index]
        result_symbol = symbol

print(result_symbol)
