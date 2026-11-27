/*
4.8 Scope and Assignment
• Existing variable ko nayi value dene par Rust purani value par turant drop call karke memory free kar deta hai.

*/


// Is lesson mein hum ye cover karenge:

// 1. Scope and assignment (purani value turant drop)

fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1. Scope and Assignment
    //----------------------------------------------------------------------------------------------------
    // Naya value assign karte hi purani "hello" par drop chalta hai aur memory free ho jati hai.
    let mut s3 = String::from("hello"); // heap par "hello"
    s3 = String::from("ahoy"); // "hello" yahin drop, ab s3 "ahoy" ko point karta hai
    println!("{s3}, world!");

}

/*  Terminal output:

ahoy, world!


*/