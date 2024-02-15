length_ft = int(input("What is the length of the room in feet? "))
width_ft = int(input("What is the width of the room in feet? "))

print(
    f"You entered dimensions of {length_ft} feet by {width_ft} feet.",
    "The area is",
    f"{length_ft * width_ft} square feet",
    f"{round(length_ft * width_ft * 0.09290304, 3)} square meters",
    sep="\n"
)