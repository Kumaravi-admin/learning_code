/*
Control Flow

Koi condition true hone par kuch code run karne ki kshamta (ability) 
aur jab tak koi condition true hai tab tak kisi code ko baar-baar run karne 
ki kshamta zyada-tar programming languages mein basic building blocks hote hain. 
Sabse aam constructs jo aapko Rust code ke execution ke flow ko control karne dete hain, 
ve hain if expressions aur loops.

*/

/*
Is lesson mein hum ye cover karenge:

1. if expressions
2. else if se multiple conditions handle karna
3. let statement mein if ka use
*/

fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1. if Expressions
    //----------------------------------------------------------------------------------------------------
    // if expression tabhi code ka block chalata hai jab condition true ho.
    // Condition false ho toh wo block skip ho jata hai (aur else wala block chalta hai, agar likha ho).
    let number = 3;

    // number < 5 true hai (3 < 5), isliye pehla block chalega.
    // Condition ke aas-paas ( ) bracket zaroori nahi hain, par { } zaroori hain.
    if number < 5 {
        println!("condition was true ");
    } else {
        println!("condition was false");
    }

    // Condition hamesha bool (true ya false) honi chahiye.
    // Rust number ko apne aap bool nahi banata (C ya JavaScript ki tarah nahi).
    // Ye compile NAHI hoga:   if number { ... }
    // Isliye hum comparison saaf likhte hain:
    if number != 0 {
        println!("number was something other than zero");
    }

    //----------------------------------------------------------------------------------------------------
    // 2. Handling Multiple Conditions with else if
    //----------------------------------------------------------------------------------------------------
    // Ek se zyada conditions check karni ho toh else if use karo.
    // Rust upar se neeche check karta hai aur sirf PEHLA true wala block chalata hai.
    // Baaki sab skip ho jate hain, chahe wo bhi true hon.
    let n = 6;

    if n % 4 == 0 {
        // 6 % 4 = 2, 0 nahi hai, isliye ye false hai
        println!("Number is divisible by 4");
    } else if n % 3 == 0 {
        // 6 % 3 = 0, isliye ye true hai: ye block chalega aur baaki skip honge
        println!("number is dibisible by 3");
    } else if n % 2 == 0 {
        // 6, 2 se bhi divisible hai, par is line tak program pahunchta hi nahi
        println!("number is divisible by 2");
    } else {
        // tabhi chalta hai jab upar ki KOI bhi condition true na ho
        println!("number is not divisible by 4,3, or 2");
    }

    //----------------------------------------------------------------------------------------------------
    // 3. Using if in a let Statement
    //----------------------------------------------------------------------------------------------------
    // Rust mein if ek EXPRESSION hai: ye ek value wapas deta hai.
    // Isliye hum isse let ke right side mein likh kar value choose kar sakte hain.
    let condition = true;

    // condition true hai, isliye pehle block ki value (5) a mein store hogi.
    // Dhyan do: 5 ya 6 ke baad semicolon nahi hai. Bina semicolon wali value hi block wapas deta hai.
    let a = if condition { 5 } else { 6 };

    println!("The value of a is: {a}");

    // Dono branches ka type SAME hona chahiye, kyunki a ka sirf ek type hota hai.
    // Ye compile NAHI hoga (number aur text mix hain):
    // let b = if condition { 5 } else { "six" };
}

/*  Terminal output:

condition was true 
number was something other than zero
number is dibisible by 3
The value of a is: 5

*/