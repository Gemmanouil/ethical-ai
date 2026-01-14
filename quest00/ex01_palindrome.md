Objective: Experience the difference between using AI as a shortcut and using it as a learning tool.
Step 1 - Do it yourself

    Write pseudocode for a function that checks if a string is a palindrome.

    first it will check a string and only , it will skip spaces and everything else that it is not a letter
    second it will check if the first letter of the string is the same with the last and every other letter of the string until it reaches the middle and has checked every position of the string 
    last but not least it will print true if the string is correct and it will print the string too  

    Implement your solution in Python.

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


    Test with examples like "racecar", "hello", and "A man a plan a canal Panama".

    Add comments explaining your reasoning.

    Check pseudocode

Step 2 - Use AI to learn

Now that your function works, use AI to go deeper:
"I wrote a palindrome function. Here's my code:
[insert your code]

What's the time complexity?

 the total time complexity is O(n)

What edge cases might I miss?

Empty strings : fix after cleaning empty string must leave 
Strings with no letters : 
Input: "12345!!!"

Cleaned: ""

Output: True | ""  

Very long strings : fix : For extremely large inputs, a two‑pointer approach avoids making a reversed copy.

Are there better approaches? 
Two‑pointer method (no reversed copy) 
Use .casefold() for better Unicode handling
Normalize Unicode

Step 3 - Reflection

    What did you learn from solving it before asking AI? 
    That i can learn from ai and i need its help and not use to do my work for me  

    How is your understanding different now? 
    That i can produse code and and make it better with the help of ai 

    Could you now write similar functions (e.g., reverse a string) without help ? not the best code for sure but i can surely try 