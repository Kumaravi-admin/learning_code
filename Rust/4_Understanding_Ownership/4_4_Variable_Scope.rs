/*
4.4 Variable Scope

• Scope = program ka wo hissa jahan tak variable valid hai.
• Variable declare hone ki jagah se valid hota hai, aur current scope ke end tak valid rehta hai.
• Variable s ek string literal hai: uski value program ke text mein hardcoded hoti hai.
*/


// Is lesson mein hum ye cover karenge:

// 1 Variable scope


fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1 Variable Scope
    //----------------------------------------------------------------------------------------------------
    // s yahan se valid hai (declare hone ke baad), aur main ke end tak valid rahega.
    let s = "hello";
    println!("{s}");

    {
        // Is block ke bahar s_var abhi exist hi nahi karta.
        let s_var = "hello"; // s_var yahan se valid hai
        println!("{s_var}");
    } // block khatam, s_var ab valid nahi (isko yahan ke baad use karoge toh compile error aayega)
}

/*  Terminal output:

hello
hello

*/