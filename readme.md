# Monkey Interpreter With Go

Following the instructions in `Writing an Interpreter in Go` with some of my own understandings

## Interpreter Prelim

A number of different categories of interpreters exist

- Simple, no parsing
    - brainfuck interpreters
- Compiled
    - compiles into internal representations called bytecode before evaluation
    - More advanced are JIT compilers, compiles just-in-time into machine codes tob be executed
- Tree walking
    - build and walk the AST(abstract syntax tree) 
    - To be implemented in this tutorial

This interpreter is made up of the following major components

- lexer
- parser
- AST
- internal object system
- evaluator

## Monkey Programming Language

Then Monkey language has a number of features

- C-like syntax
- Variable Binding, Integer, String and Booleans
    ```js
    let age = 1;
    let name = 'Monkey'; 
    ```
- Arithmetic expressions
    ```js
    let result = 10 * (20 / 2);
    ```
- Built-in functions
- First class and higher order functions
    - **First class functions** are just values like integers and strings that can be used as input in another function
    - Functions that takes other functions as input are **higher order functions**
    - A classical example is the twice function
        ```js
        let addOne = fn(a) { return a + 1; };
        // return statement can be implicit
        // let add = fn(a, b) { a + b; };

        let twice = fn(f, a) {
            f(f(a));
        }

        twice(addOne(2)); // => 4
        ```
- Closures
- Array and hash table
    ```js
    let myArray = [1, 2, 3, 4, 5];
    let thorsten = {"name": "Thorsten", "age": 28};

    // accessing is done using the index expressions
    myArray[0] // => 1
    thorsten["name"] // => "Thorsten"
    ```