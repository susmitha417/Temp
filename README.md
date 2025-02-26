# Temperature Converter

## Description
This program converts temperatures between Celsius and Fahrenheit using a `Temperature` struct and pointer-based methods in Golang.

## Features
- Converts Celsius to Fahrenheit.
- Converts Fahrenheit to Celsius.
- Uses pointers to modify temperature values directly within methods.

## Formula
- °F = (°C × 9/5) + 32
- °C = (°F - 32) × 5/9

## Code Explanation

### 1. Define the `Temperature` Struct
The `Temperature` struct contains:
- `Temp` (float64): Stores the temperature value.

```go
// Temperature struct
type Temperature struct {
    Temp float64
}
```

### 2. Define Methods to Convert Temperature
We define two methods using pointers so that the original `Temp` value gets updated.

#### Convert Celsius to Fahrenheit
This method modifies the `Temp` value using the conversion formula.

```go
func (t *Temperature) ToFahrenheit() {
    t.Temp = (t.Temp * 9 / 5) + 32
}
```

#### Convert Fahrenheit to Celsius
This method modifies the `Temp` value using the conversion formula.

```go
func (t *Temperature) ToCelsius() {
    t.Temp = (t.Temp - 32) * 5 / 9
}
```

### 3. Main Function
The `main` function demonstrates how to use the `Temperature` struct and its methods.

#### Convert 35°C to Fahrenheit
1. Create an instance of `Temperature` with `Temp = 35`.
2. Call `ToFahrenheit()` to convert it to Fahrenheit.
3. Print the result.

```go
celsiusTemp := Temperature{Temp: 35} // Create a Temperature object
celsiusTemp.ToFahrenheit() // Convert to Fahrenheit
fmt.Printf("Converted: %.2f°F\n", celsiusTemp.Temp) // Output: Converted: 95.00°F
```

#### Convert 95°F to Celsius
1. Create an instance of `Temperature` with `Temp = 95`.
2. Call `ToCelsius()` to convert it to Celsius.
3. Print the result.

```go
fahrenheitTemp := Temperature{Temp: 95} // Create a Temperature object
fahrenheitTemp.ToCelsius() // Convert to Celsius
fmt.Printf("Converted: %.2f°C\n", fahrenheitTemp.Temp) // Output: Converted: 35.00°C
```

The program will output:
   ```
   Converted: 95.00°F
   Converted: 35.00°C
   ```

