# hugenum

**hugenum** is a Go package that provides a convenient and efficient way to work with extremely large numbers, beyond the limitations of the built-in numeric types. This package implements a `HugeNumber` struct that can represent and perform arithmetic operations on numbers of virtually any magnitude, limited only by the available memory.

## Features

- **Arbitrary Precision**: Represent and perform calculations on numbers of virtually any size, limited only by available memory.
- **Engineering Notation**: Numbers are represented in engineering notation, with scaling and naming of exponents (e.g., thousand, million, billion, etc.).
- **Efficient Arithmetic**: Optimized algorithms for performing arithmetic operations on large numbers.
- **Normalization and Alignment**: Methods for normalizing and aligning numbers to a consistent representation.
- **Formatting and Conversion**: Convert `HugeNumber` instances to and from string representations with customizable precision.

## Installation
```sh
go get github.com/TrungKenbi/hugenum
```

## Usage

```go
import "github.com/TrungKenbi/hugenum"

func main() {
    // Create a new HugeNumber
    num1 := hugenum.New(1.23, 6) // 1.23 million

    // Perform arithmetic operations
    num2 := hugenum.New(4.56, 3) // 4.56 thousand
    num1.Add(num2)                         // num1 is now 1.234560 million

    // Convert to string representation
    fmt.Println(num1.String()) // Output: 1.235 million
}
```

## Documentation
Detailed documentation and examples can be found in the GoDoc reference.
## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any bugs, improvements, or new features.
## License
This project is licensed under the MIT License.
