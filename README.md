# Assignment

Given:
A polynomial of degree m where k = m + 1 points are required to reconstruct it.
The (x, y) values are provided in a JSON format where:
Key = x-coordinate

Value = encoded y in a certain base (e.g., base 2, base 10, base 16, etc.)
Task is to:
Parse the JSON input file
Decode y-values from their base to decimal using Go
Apply Lagrange Interpolation to find f(0) (the constant term c)

Features
JSON parsing using encoding/json
Base conversion using math/big's SetString (for arbitrary base and precision)
Lagrange interpolation using big integers (no floating-point errors)
Easy to extend for modular arithmetic if needed
