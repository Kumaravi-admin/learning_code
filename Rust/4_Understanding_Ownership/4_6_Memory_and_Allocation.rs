/*
4.6 Memory and Allocation

• String literal: contents compile time par pata hote hain, text seedha executable mein hardcoded hota hai, isliye fast hai.
• Runtime par badalne wale ya unknown size ke text ke liye heap par memory chahiye.
• Iska matlab: memory runtime par allocator se request karni padti hai, aur kaam ke baad wapas bhi karni padti hai.
• Request wala part: String::from khud kar leta hai (har language mein lagbhag same).
• Wapas karne wala part alag hai:
    - GC wali languages: GC khud unused memory saaf karta hai.
    - Bina GC wali languages: programmer ko khud free karna padta hai. Bhoolo toh memory waste, jaldi karo toh invalid variable, do baar karo toh bug.
    - Rust: owner variable scope se bahar jaye toh memory apne aap wapas ho jati hai.
• Scope khatam hone par Rust automatically drop function call karta hai (closing curly bracket par).
• C++ mein is pattern ko RAII (Resource Acquisition Is Initialization) kehte hain.
*/


//  Is lesson mein hum ye cover karenge:

//  1. Memory and allocation (String ki memory scope khatam hote hi free hoti hai)


//----------------------------------------------------------------------------------------------------
// 1. Memory and Allocation
//----------------------------------------------------------------------------------------------------
// Ek allocate ke saath ek hi free hona chahiye, aur Rust ye khud karta hai:
//
// {
//     let s = String::from("hello"); // s yahan se valid hai, heap par memory allocate hui
//
//     // s ke saath kaam karein
// }                                  // scope khatam, Rust ne drop call kiya, memory free ho gayi
//                                    // s ab valid nahi hai