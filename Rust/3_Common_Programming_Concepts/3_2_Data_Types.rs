/*
3.2 Data Types

• Rust statically typed hai: har variable ka type compile time par pata hona chahiye.
• parse() jaisi jagah type annotation dena zaroori hai.

Floating-Point:
• Do types: f32 (32 bit) aur f64 (64 bit). Default f64 hai (lagbhag f32 jitna fast, zyada precise).
• Sab signed hote hain.

Numeric Operations:
• +, -, *, /, %. Integer division zero ki taraf truncate hota hai (-5 / 3 = -1).

Boolean:
• bool: true ya false, size 1 byte.

Character:
• char: single quote mein ek character, emoji bhi chalta hai.

Tuple:
• Alag-alag types ki values ka group, length fixed.
• Destructuring ya .index se access karte hain.

Array:
• Sab elements ka type same, length fixed, data stack par.
• Size badalna ho toh Vector use karo.
*/

/*
Is lesson mein hum ye cover karenge:

1. Type annotation aur parse
2. Floating-point types
3. Numeric operations
4. Boolean type
5. Character type
6. Tuple type
7. Array type
*/

fn main() {

    // 1. Type annotation aur parse
    // parse() ko pata hona chahiye kaunsa number banana hai, isliye ": u32" likha hai.
    // Text number nahi hua toh expect() program rok kar message dikhata hai.
    let guess: u32 = "42".parse().expect("Not a number!");
    println!("{}", guess);

    // 2. Floating-point types
    let x = 2.0; // f64 (default)
    let y: f32 = 3.0; // f32 (annotation se)
    println!("Output of x: {}", x);
    println!("Output of y: {}", y);

    // 3. Numeric operations
    let sum = 5 + 10; // addition
    println!("Addition: {}", sum);

    let differance = 98.5 - 4.3; // subtraction
    println!("Subtraction: {}", differance);

    let product = 4 * 30; // multiplication
    println!("Multiplication: {}", product);

    let quatient = 56.7 / 32.2; // float division
    let truncated = -5 / 3; // integer division: decimal part kat jata hai, result -1
    println!("Division quatient: {}", quatient);
    println!("Division truncated: {}", truncated);

    let remainder = 43 % 5; // bacha hua hissa: 43 = 8 * 5 + 3
    println!("Remainder: {}", remainder);

    // 4. Boolean type
    let t = true; // type Rust khud samajh leta hai
    let f: bool = false; // type likh kar diya
    println!("Result of t: {}", t);
    println!("Result of f: {}", f);

    // 5. Character type
    // char single quote ' ' mein hota hai (double quote " " wala string hota hai).
    let c = 'z';
    let z: char = 'z';
    let heart_eyed_cat = '😻';
    println!("Result of c: {}", c);
    println!("Result of z: {}", z);
    println!("Result of heart eyed cat: {}", heart_eyed_cat);

    // 6. Tuple type
    let tup = (500, 6.4, 1);

    // Tarika 1: destructuring, ek line mein alag variables.
    // Ye naye x, y, z upar ke x aur y ko shadow kar dete hain.
    let (x, y, z) = tup;
    println!("The value of x is: {x}");
    println!("The value of y is: {y}");
    println!("The value of z is: {z}");

    // Tarika 2: dot aur index (0 se shuru).
    let x: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = x.0;
    let six_point_four = x.1;
    let one = x.2;

    println!("Result of five hundered: {five_hundred}");
    println!("Result of five six point four: {six_point_four}");
    println!("Result of five one: {one}");

    // 7. Array type
    // Sab elements same type ke, length fixed. Index 0 se shuru hota hai.
    let a = [1, 2, 3, 4, 5];
    println!("Output of a index 3: {}", a[3]); // 4

    // Mahine hamesha 12 hote hain, isliye array sahi choice hai (Vector ki zaroorat nahi).
    let months = ["January", "February", "March", "April", "May", "June", "July",
                  "August", "September", "October", "November", "December"];
    println!("Output of moths 6th index value: {}", months[6]); // July

    // Array ka type [type; length] ki tarah likhte hain.
    let a: [i32; 5] = [1, 2, 3, 4, 5];
} // end of main function

/*  Terminal output:

42
Output of x: 2
Output of y: 3
Addition: 15
Subtraction: 94.2
Multiplication: 120
Division quatient: 1.7608695652173914
Division truncated: -1
Remainder: 3
Result of t: true
Result of f: false
Result of c: z
Result of z: z
Result of heart eyed cat: 😻
The value of x is: 500
The value of y is: 6.4
The value of z is: 1
Result of five hundered: 500
Result of five six point four: 6.4
Result of five one: 1
Output of a index 3: 4
Output of moths 6th index value: July

*/