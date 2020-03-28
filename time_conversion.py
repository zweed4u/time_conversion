#!/usr/bin/python3


def timeConversion(s: str) -> str:
    meridiem = s[-2:]
    hour = int(s[:2])
    if meridiem == "AM":
        if hour == 12:  # 12AM
            return f"00{s[2:-2]}"
        return s[:-2]  # ie. 3AM
    if hour == 12:  # 12PM
        return s[:-2]
    return f"{hour+12}{s[2:-2]}"  # ie. 3PM


timeConversion("12:24:32AM")
