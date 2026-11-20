

fn main() {
    println!("Hello, world!");
    another_function();
    function_with_parameter(5);
    print_labeled_measurement(5,'h');
    

    let e = five();
    println!("The value of e is: {e}");


    let f = plus_one(5);
    println!("The value of f is: {f}");




    //----------------------------------------------------------------------------------------------------
    // 5. Statements and Expressions
    //----------------------------------------------------------------------------------------------------
    // Function bodies statements ki ek series se bani hoti hain jo optionally ek expression par khatam hoti hain.
    // Abhi tak humne jitne bhi functions cover kiye hain unmein koi ending expression shamil nahi tha, lekin aapne 
    // ek statement ke hisse ke roop mein ek expression dekha hai. Kyunki Rust ek expression-based language hai, 
    // isliye is antar ko samajhna bahut zaroori hai. Doosri languages mein aise antar nahi hote, toh chaliye dekhte 
    // hain ki statements aur expressions kya hain aur unke antar functions ki bodies ko kaise prabhavit karte hain.
    // Statements aise instructions hote hain jo koi action perform karte hain aur koi value return nahi karte.
    // Expressions evaluate hokar ek resultant value dete hain.
    // Chaliye kuch examples dekhte hain.
    // Humne darasal statements aur expressions ka istemal pehle hi kar liya hai. let keyword ke saath ek variable
    //  banana aur use ek value assign karna ek statement hai. Listing 3-1 mein, let y = 6; ek statement hai.
   
    let y = 6;
    println!("Output of let statement: {y}");

    // thi code does note complie!
    // let x = (let y = 6);

    let a = {
        let b = 3;
        b + 1
    };
    println!("The value of a is: {a}");

}

//----------------------------------------------------------------------------------------------------
// 1. Basic function
//----------------------------------------------------------------------------------------------------
// Functions Rust code mein har jagah dekhe jaate hain. 
// Aapne language ke sabse important functions mein se ek ko pehle hi dekh liya hai: main function, 
// jo kai programs ka entry point hota hai. 
// Aapne fn keyword bhi dekha hai, jo aapko naye functions declare karne ki suvidha deta hai.
// Rust code function aur variable ke naamon ke liye conventional style ke roop mein snake case ka istemal karta hai, 
// jismein saare letters lowercase hote hain aur underscores words ko alag karte hain. 
// Yahan ek program hai jismein ek example function definition hai:

fn another_function() {
    println!("Another function.");
}


//----------------------------------------------------------------------------------------------------
// 2. Function with parameter
//----------------------------------------------------------------------------------------------------
// Hum functions ko is tarah define kar sakte hain ki unmein parameters hon, 
// jo special variables hote hain aur function ke signature ka hissa hote hain. 
// Jab kisi function mein parameters hote hain, toh aap un parameters ke liye concrete values de sakte hain. 
// Technically, un concrete values ko arguments kaha jata hai, lekin aam bolchal mein log function 
// ki definition ke variables ya function call karte waqt pass ki gayi concrete values dono ke liye 
// parameter aur argument words ka istemal ek-doosre ki jagah (interchangeably) karte hain.
// another_function ke is version mein hum ek parameter add karte hain:

fn function_with_parameter(x: i32) {
    println!("The value of x is: {x}");
}



//----------------------------------------------------------------------------------------------------
// 3. Function with multiple parameter
//----------------------------------------------------------------------------------------------------
// Function signatures mein, aapko har parameter ka type declare karna zaroori (must) hai. 
// Yeh Rust ke design mein ek jaan-boojhkar liya gaya faisla hai: Function definitions mein 
// type annotations ko zaroori banane ka matlab hai ki compiler ko code mein kahin aur iski 
// zaroorat lagbhag kabhi nahi padti taaki woh samajh sake ki aapka matlab kis type se hai. 
// Agar compiler ko pata hota hai ki function kis type ki ummeed kar raha hai, toh woh zyada 
// madadgar error messages bhi de pata hai.
// Jab aap multiple parameters define karte hain, toh parameter declarations ko commas se alag karein, kuch is tarah:

fn print_labeled_measurement(value: i32,unit_lable:char) {
    println!("The measuremnt is: {value}{unit_lable}");
}

//----------------------------------------------------------------------------------------------------
// 4. Functions with Return values
//----------------------------------------------------------------------------------------------------
// Functions unhe call karne wale code ko values return kar sakte hain. 
// Hum return values ko naam nahi dete, lekin humein ek arrow (->) ke baad unka type declare karna zaroori hota hai. 
// Rust mein, function ki return value function body ke block ke aakhiri expression ki value ke barabar (synonymous) hoti hai. 
// Aap return keyword ka istemal karke aur ek value specify karke kisi function se jaldi (early) return kar sakte hain, 
// lekin zyada-tar functions aakhiri expression ko implicitly return karte hain. 
// Yahan ek function ka example hai jo ek value return karta hai:

fn five() -> i32 {
    5
}


// Doosri, five function mein koi parameters nahi hain aur yeh return value ka type define karta hai, 
// lekin function ki body mein bina semicolon ke ek akela 5 hai kyunki yeh ek expression hai jiski value hum return karna chahte hain.
// Chaliye ek aur example dekhte hain:

fn plus_one(x:i32) -> i32 {
    x + 1
}
// Is code ko run karne par The value of x is: 6 print hoga. Lekin


/*
kya hoga agar hum x + 1 wali line ke ant mein ek semicolon laga dein, 
aur ise ek expression se statement mein badal dein?

[This code does not compile!]

fn plus_one(x: i32) -> i32 {
    x + 1;
}
*/


/*          Terminal output



Hello, world!
Another function.
The value of x is: 5
The measuremnt is: 5h
The value of e is: 5
The value of f is: 6
Output of let statement: 6
The value of a is: 4
*/