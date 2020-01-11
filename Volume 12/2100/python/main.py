result = 2
for i in range(int(input())):
    if input()[-4:] == '+one':
        result += 2
    else:
        result += 1

if result == 13:
    print('1400')
else:
    print(result * 100)
