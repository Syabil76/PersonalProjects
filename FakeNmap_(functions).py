#create devices through fake networks project
import random
from datetime import datetime
import string
#list

#functions
devices = list()
def create_devices(num_device):
    if num_device > 254:
        print("Too many devices! (254<)")
        return devices
    print(f"-----scanning started at {datetime.now()}-----")
    for i in range(0,num_device):
        device = dict()
        device["name"] = (
        random.choice(['2F','57','PQ','SD','D0'])
        + random.choice(['2F','57','PQ','SD','D0'])
        + random.choice(['2F','57','PQ','SD','D0'])
        )
        device["IP"] = (
        f"10.0.{str(random.randint(0,255))}.{str(random.randint(0,255))}"
        )
        mac = list()
        for x in range(12):
            mac.append(
            random.choice(["A",'B','C','D','E','F','1','2','3','4','5','6','7','8','9','0'])
            )
        s = ''.join(map(str, mac))
        device["MAC"] = s
        devices.append(device)
    return devices

num_device = int(input("create how many devices?"))

create_devices(num_device)
for i in range(len(devices)):
    print(devices[i])
