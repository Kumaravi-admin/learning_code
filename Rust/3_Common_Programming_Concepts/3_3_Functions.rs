/*
Functions

• Function naam aur variable naam snake_case mein likhte hain (lowercase, words ke beech underscore).
• fn keyword se function banate hain. main function program ka entry point hai.
• Parameter function signature ka hissa hai, aur har parameter ka type likhna zaroori hai.
• Parameter = definition mein variable, argument = call karte waqt di gayi value.
• Return value ka type arrow (->) ke baad likhte hain.
• Function ka aakhiri expression (bina semicolon) hi return value hoti hai. Jaldi nikalne ke liye return keyword bhi hai.

Statements aur Expressions:
• Rust expression-based language hai, isliye ye farq zaroori hai.
• Statement: koi kaam karta hai, value return nahi karta (jaise let y = 6;).
• Expression: evaluate hokar value deta hai (jaise 5 + 6, ya { } block).
*/

/*
Is lesson mein hum ye cover karenge:

1. Basic function
2. Function with parameter
3. Function with multiple parameters
4. Functions with return values
5. Statements and expressions
*/

fn main() {
    println!("Hello, world!");

    // 1-4: neeche likhe functions ko yahan se call kar rahe hain.
    another_function(); // 1
    function_with_parameter(5); // 2: 5 argument hai
    print_labeled_measurement(5, 'h'); // 3: do arguments

    let e = five(); // 4: function ki return value e mein aati hai
    println!("The value of e is: {e}");

    let f = plus_one(5); // 4: 5 + 1 = 6
    println!("The value of f is: {f}");

    //----------------------------------------------------------------------------------------------------
    // 5. Statements and Expressions
    //----------------------------------------------------------------------------------------------------
    // let y = 6; ek statement hai: kaam karta hai, par value return nahi karta.
    let y = 6;
    println!("Output of let statement: {y}");

    // Statement ki value nahi hoti, isliye ye compile NAHI hota:
    // let x = (let y = 6); // error: expected expression, found `let` statement

    // { } block ek expression hai, aur uski value aakhiri line (bina semicolon) hoti hai.
    let a = {
        let b = 3;
        b + 1 // semicolon nahi hai, isliye 4 block ki value ban kar a mein jata hai
    };
    println!("The value of a is: {a}");
}

//----------------------------------------------------------------------------------------------------
// 1. Basic function
//----------------------------------------------------------------------------------------------------
// fn ke baad naam (snake_case), phir () aur { } mein body.
// Function main ke upar ya neeche kahin bhi likh sakte ho.
fn another_function() {
    println!("Another function.");
}

//----------------------------------------------------------------------------------------------------
// 2. Function with parameter
//----------------------------------------------------------------------------------------------------
// x parameter hai, aur uska type (i32) likhna zaroori hai.
fn function_with_parameter(x: i32) {
    println!("The value of x is: {x}");
}

//----------------------------------------------------------------------------------------------------
// 3. Function with multiple parameter
//----------------------------------------------------------------------------------------------------
// Multiple parameters comma se alag karte hain, har ek ka type alag likhna padta hai.
fn print_labeled_measurement(value: i32, unit_lable: char) {
    println!("The measuremnt is: {value}{unit_lable}");
}

//----------------------------------------------------------------------------------------------------
// 4. Functions with Return values
//----------------------------------------------------------------------------------------------------
// Return type -> ke baad likhte hain. Return value ko naam nahi dete.
// Body mein bina semicolon ka 5 ek expression hai, wahi return hota hai.
fn five() -> i32 {
    5
}

fn plus_one(x: i32) -> i32 {
    x + 1 // semicolon nahi, isliye ye expression hai aur value return hoti hai
}

// Semicolon laga diya toh x + 1; statement ban jata hai, koi value nahi milti, aur compile NAHI hota:
//
// fn plus_one(x: i32) -> i32 {
//     x + 1;
// }
// error[E0308]: mismatched types (i32 chahiye tha, () mila)

/*  Terminal output:

Hello, world!
Another function.
The value of x is: 5
The measuremnt is: 5h
The value of e is: 5
The value of f is: 6
Output of let statement: 6
The value of a is: 4

*/