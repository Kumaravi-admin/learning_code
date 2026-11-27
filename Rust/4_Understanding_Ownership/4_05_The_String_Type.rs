/*
4.5 The String Type

• String literal ki value compile time par pata hoti hai, wo immutable hai, aur hardcoded hoti hai.
• Har text compile time par known nahi hota (jaise user input), isliye String type chahiye.
• String heap par allocate hoti hai, isliye uska size runtime par badal sakta hai.
• String::from se literal se String banate hain. :: operator from function ko String type ke namespace mein rakhta hai.
• String ko mutate kar sakte hain (mut ke saath), literal ko nahi.
*/


//  Is lesson mein hum ye cover karenge:

//  1. String::from se String banana
//  2. Mutable String (push_str)


fn main() {

    //----------------------------------------------------------------------------------------------------
    // 2. The String Type
    //----------------------------------------------------------------------------------------------------

    // 2.1 String::from: literal se heap wali String banao.
    let st = String::from("hello");
    println!("{st}");

    // 2.2 mut ke saath String badal sakti hai, push_str() end mein text jodta hai.
    let mut message = String::from("hello");
    message.push_str(", world!");
    println!("{message}"); // hello, world!

    // Literal badal kyun nahi sakta, String kyun badal sakti hai?
    // Dono memory ko alag tarah se handle karte hain (agle section mein dekhenge).
}

/*  Terminal output:

hello
hello, world!

*/