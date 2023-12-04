import fileinput

text_digits = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

in_ = """
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
"""

def main():
    sum = 0
    # for line in fileinput.input():
    for line in in_.split("\n"):
        line = line.strip()
        if line == "":
            continue

        original = line
        first_idx = len(line)
        first_sub = ""
        last_idx = -1
        last_sub = ""
        for k, v in text_digits.items():
            lidx = line.find(k)
            ridx = line.rfind(k)
            if lidx != -1 and lidx < first_idx:
                first_idx = lidx
                first_sub = k
            if ridx != -1 and ridx > last_idx:
                last_idx = ridx
                last_sub = k

        if first_sub != "":
            new_line = list(line)
            new_line[first_idx] = str(text_digits[first_sub])
            line = "".join(new_line)

        if last_sub != "":
            new_line = list(line)
            new_line[last_idx] = str(text_digits[last_sub])
            line = "".join(new_line)

        for k, v in text_digits.items():
            line = line.replace(k, str(v))

        only_digits = ""
        for c in line:
            if not c.isdigit():
                continue
            only_digits += c

        value = int(only_digits[0]+only_digits[-1])
        print(original, only_digits, value)

        sum += value

    print(sum)


if __name__ == "__main__":
    main()
