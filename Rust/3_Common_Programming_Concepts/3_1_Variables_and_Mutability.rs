// 3.1 Variables and mutability

/*
We will cover in this lesson:

1. Immutable variables (default) and mut
2. Shadowing
3. Shadowing in an inner scope
4. Shadowing to change the type (mut cannot do this)
*/

fn main() {

    // ---------------------------------------------------------
    // 1. Immutable variables (default) and mut
    // ---------------------------------------------------------
    // In Rust, a variable is IMMUTABLE by default: once set, it cannot change.
    // This code gives a compile error: "cannot assign twice to immutable variable".
    //
    // let x = 5;
    // println!("The value of x is :{x}");
    // x = 6;
    // println!("The value of x is: {x}");
    //
    // Fix: add "mut" (mutable), then the value can change:
    // let mut x = 5;

    // ---------------------------------------------------------
    // 2. Shadowing
    // ---------------------------------------------------------
    // Writing "let x" again creates a NEW variable with the same name.
    // The new x "shadows" (hides) the old x. The old x is not changed, it is just hidden.
    let x = 5;
    let x = x + 1; // new x = 6

    // ---------------------------------------------------------
    // 3. Shadowing in an inner scope
    // ---------------------------------------------------------
    // Inside { } we create another x. It lives only inside this block.
    // When the block ends, that x is gone and the outer x (6) is visible again.
    {
        let x = x * 2; // 6 * 2 = 12
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}"); // back to the outer x, still 6

    // ---------------------------------------------------------
    // 4. Shadowing to change the type
    // ---------------------------------------------------------
    // With "mut" you can change the value, but NOT the type.
    // This gives a compile error (a string cannot become a number):
    //
    // let mut spaces = "     ";
    // spaces = spaces.len();
    //
    // Shadowing creates a new variable, so the type is allowed to change:
    let spaces = "           ";   // type: &str (text)
    let spaces = spaces.len();   // type: usize (number of characters)
    println!("space len: {spaces}");
}

/*  Terminal output:

The value of x in the inner scope is: 12
The value of x is: 6
space len: 11

*/