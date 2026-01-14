# A list of example strings to test
Input = [
    "racecar",
    "hello",
    "A man a plan a canal Panama"
]

# Loop through each string in the list
for s in Input:
    # Keep only letters and convert them to lowercase
    cleaned = "".join(c.lower() for c in s if c.isalpha())
    
    # Check if the cleaned string is a palindrome and print the result
    print(cleaned == cleaned[::-1], "|", cleaned)

