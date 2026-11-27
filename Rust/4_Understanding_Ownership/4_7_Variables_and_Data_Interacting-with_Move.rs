/*
4.7 Variables and Data Interacting with Move

• Integer: simple fixed size value hai, dono 5 stack par push hote hain, isliye x se y mein copy ho jata hai.
• String: teen parts ki hoti hai (pointer, length, capacity), ye stack par rehte hain. Asli text heap par rehta hai.
• let s2 = s1; stack ka pointer, length, capacity copy karta hai. Heap ka data copy nahi hota.
• Dono scope se bahar jaate toh ek hi memory do baar free hoti (double free, memory safety bug).
• Isliye Rust s1 ko invalid kar deta hai. Is ko move kehte hain (shallow copy + purana variable invalid).
• Rust kabhi bhi automatically heap data ki deep copy nahi banata, isliye automatic copy sasti hoti hai.
*/


//  Is lesson mein hum ye cover karenge:

//  1. Move (integer copy vs String move)


fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1. Variables and Data Interacting with Move
    //----------------------------------------------------------------------------------------------------

    // 1.1 Integer: value copy hoti hai, x aur y dono valid rehte hain.
    let x = 5;
    let y = x;
    println!("{y}");

    // 1.2 String: copy nahi, MOVE hota hai.
    //
    //   s1 (stack)              heap
    //   ptr  ───────────►  [ h e l l o ]
    //   len = 5
    //   cap = 5
    //
    // let s2 = s1; ke baad stack wala hissa s2 mein chala jata hai, heap ka "hello" wahi ek hi rehta hai.
    // s1 ab invalid hai, isliye sirf s2 hi memory free karega.
    let s1 = String::from("hello");
    let s2 = s1;
    println!("{s2}, world!");

    // s1 ab use nahi kar sakte. Uncomment karoge toh compile error aayega:
    // println!("{s1}, world!"); // error[E0382]: borrow of moved value: `s1`

    // Heap data ki asli copy chahiye toh clone() use karo (aage dekhenge).
}

/*  Terminal output:

5
hello, world!

*/