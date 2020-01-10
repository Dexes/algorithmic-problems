input()
data = list(map(int, input().split()))

candidate = data[0]
candidate_index = -1
candidate_count = 0

max_candidate_index = -1
max_count = 0

for index, pirate in enumerate(data[1:]):
    candidate_count += 1
    if pirate < candidate:
        if max_count < candidate_count:
            max_count = candidate_count
            max_candidate_index = candidate_index

        candidate = pirate
        candidate_index = index
        candidate_count = 1

if max_count < candidate_count:
    max_candidate_index = candidate_index

print(max_candidate_index + 2)
