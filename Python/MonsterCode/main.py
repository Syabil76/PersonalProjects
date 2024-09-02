monsters = dict()
stats = dict{
    
}
with open("test.txt", 'r') as file:
    lines = file.readlines()
    for line in lines:
        data = line.strip().split(',')
        monsters[data[0]] = [data[1], *data[2:]]

for key, value in monsters.items():
    print(f"Monster: {key}\n{value}")

