/*
4.10 Copy (stack-only data)
• Integer jaise types poori tarah stack par hote hain, isliye copy sasti hai. Deep aur shallow copy mein koi farq nahi.
• Aise types par Copy trait hota hai: assign karne par move nahi, copy hota hai, purana variable valid rehta hai.
• Jis type ya uske kisi part ne Drop trait implement kiya ho, usme Copy nahi laga sakte (compile error).
• Copy wale types: saare integer (u32), bool, saare float (f64), char, aur tuple (sirf tab jab andar ke sab types Copy hon).
• (i32, i32) Copy hai, (i32, String) Copy nahi hai.
• Jis cheez ko allocation chahiye ya jo resource hai, wo Copy nahi ho sakti.
*/

//  Is lesson mein hum ye cover karenge:

//  1. Copy (integers move nahi hote)

fn main() {


    //----------------------------------------------------------------------------------------------------
    // 1. Copy (stack-only data)
    //----------------------------------------------------------------------------------------------------
    // Integer par Copy trait hai: x2 move nahi hota, copy hota hai, aur clone() bhi nahi chahiye.
    let x2 = 5;
    let y2 = x2;
    println!("x2 = {x2}, y2 = {y2}");
}

/*  Terminal output:


x2 = 5, y2 = 5

*/