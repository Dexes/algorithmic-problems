class BIT:
    _size = 0
    _data = None

    def __init__(self, size):
        self._size = size
        self._data = [0] * size

    @property
    def size(self) -> int:
        return self._size

    def update_range(self, left, right, delta):
        self._data[left] += delta
        print('left', left, delta)

        right += 1
        if right < self._size:
            self._data[right] -= delta
            print('right', right, delta)

    def get_element(self, index):
        result = 0
        for i in range(index + 1):
            result += self._data[i]

        return result

    def calculate_data(self):
        result = [0] * self._size
        result[0] = self._data[0]

        for i in range(1, self._size):
            result[i] += result[i - 1] + self._data[i]

        return result


bit = BIT(int(input()))
for i in range(bit.size + 1):
    left, right, coins = input().split()
    bit.update_range(int(left) - 1, int(right) - 1, int(coins))

print(' '.join(map(str, bit.calculate_data())))
