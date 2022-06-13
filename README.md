# About
This is a simple letter permutation application. It receives as an input a set of letters, the length of the desired results and it returns all the dictionary words that contain combinations of letters received as an input.

It also allows the user to update the dictionary with new words if needed.

# Usage
To run the application you need to have Golang installed.
1. Perm solver
- Clone the repository locally
- From PermSolver folder run <code>go run cmd/main.go mansion 4 5</code>

The result will be the list of words that have between 4 and 5 letters, ordered alphabetically, that contain letters from "mansion".

If only the initial word is passed as argument the application will return all the words starting from 3 letter up to the number of input letters.

2. Add to dictionary
- Clone the repository locally
- From PermSolver folder run <code>go run cmd/main.go add newWord</code>

Enjoy!
