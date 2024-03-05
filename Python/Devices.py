B01 = {
    "Band" : "Parkway Drive",
    "Genre" : "Metalcore, Melodic Metalcore",
    "Region" : "Australia",
    "Status" : "Headline",
}
B02 = {
    "Band" : "Lorna Shore",
    "Genre" : "Deathcore",
    "Region" : "America",
    "Status" : "Headline",
}
B03 = {
    "Band" : "Sylosis",
    "Genre" : "Metalcore, Modern Metal",
    "Region" : "America",
    "Status" : "Starter",
}


Wacken_2024 = {
    "Entry" : B03,
    "HEADLINER Evening" : B02,
    "HEADLINER Midnight" : B01,
}

for key, value in Wacken_2024.items():
    print(f"{key:>16s} : {value}")
