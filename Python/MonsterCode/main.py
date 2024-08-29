
monsters = dict()
with open("index.txt", 'r') as file:
    lines = file.readlines()
    for line in lines:
        data = line.strip().split(',')
        monsters[data[1]] = [data[0], *data[2:]]
print(monsters)

for key, value in monsters.items():
    print(tabulate(key, value))
