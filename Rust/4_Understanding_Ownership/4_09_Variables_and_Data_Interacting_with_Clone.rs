/*
4.9 Clone
• Heap data ki deep copy chahiye toh clone() use karo.
• clone() ka call dikhe toh samajh lo ki kuch bhaari (expensive) ho sakta hai.

*/

// Is lesson mein hum ye cover karenge:

// 1. Clone (heap data ki deep copy)



fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1. Variables and Data Interacting with Clone
    //----------------------------------------------------------------------------------------------------
    // clone() heap data bhi copy karta hai, isliye s4 aur s5 dono valid rehte hain.
    let s4 = String::from("hello");
    let s5 = s4.clone();
    println!("s4 = {s4}, s5 = {s5}");

   
}

/*  Terminal output:


s4 = hello, s5 = hello


*/