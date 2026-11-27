# We will cover in this lesson

1. Basic function declaration
2. Function with parameter
3. Multiple parameter function
4. Statements and Expressions
5. Function with return value

---

## 1. Basic function declaration

Functions Rust code mein har jagah dekhe jaate hain. Aapne language ke sabse important functions mein se ek ko pehle hi dekh liya hai: main function, jo kai programs ka entry point hota hai. Aapne `fn` keyword bhi dekha hai, jo aapko naye functions declare karne ki suvidha deta hai.

Rust code function aur variable ke naamon ke liye conventional style ke roop mein _snake case_ ka istemal karta hai, jismein saare letters lowercase hote hain aur underscores words ko alag karte hain. Yahan ek program hai jismein ek example function definition hai:

**Filename: `src/main.rs`**

```rust
fn main() {
    println!("Hello, world!");
    another_function();
}

fn another_function() {
    println!("Another function.");
}
```

Hum Rust mein `fn` likhkar, uske baad ek function name aur parentheses ka set lagakar ek function define karte hain. Curly brackets compiler ko batate hain ki function body kahan shuru aur khatam hoti hai.

Hum apne define kiye gaye kisi bhi function ko uska naam aur uske baad parentheses ka set likhkar call kar sakte hain. Kyunki `another_function` program mein define kiya gaya hai, isliye ise `main` function ke andar se call kiya ja sakta hai. Dhyaan dein ki humne source code mein `another_function` ko `main` function ke baad define kiya hai; hum ise pehle bhi define kar sakte the. Rust ko is baat se koi farak nahi padta ki aap apne functions kahan define karte hain, bas itna zaroori hai ki ve kisi aise scope mein define hon jise caller dekh sake.

Chaliye functions ko aur गहराई (depth) se explore karne ke liye `functions` naam ka ek naya binary project shuru karte hain. `another_function` wale example ko _src/main.rs_ mein rakhein aur ise run karein. Aapko niche diya gaya output dikhna chahiye:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.28s
     Running `target/debug/functions`
Hello, world!
Another function.
```

Lines usi order mein execute hoti hain jis order mein ve `main` function mein dikhai deti hain. Pehle “Hello, world!” message print hota hai, aur phir `another_function` call hota hai aur uska message print hota hai.

---

## 2. Function with parameter

Hum functions ko is tarah define kar sakte hain ki unmein _parameters_ hon, jo special variables hote hain aur function ke signature ka hissa hote hain. Jab kisi function mein parameters hote hain, toh aap un parameters ke liye concrete values de sakte hain. Technically, un concrete values ko _arguments_ kaha jata hai, lekin aam bolchal mein log function ki definition ke variables ya function call karte waqt pass ki gayi concrete values dono ke liye _parameter_ aur _argument_ words ka istemal ek-doosre ki jagah (interchangeably) karte hain.

`another_function` ke is version mein hum ek parameter add karte hain:

**Filename: `src/main.rs`**

```rust
fn main() {
    function_with_parameter(5);
}

fn function_with_parameter(x: i32) {
    println!("The value of x is: {x}");
}
```

Is program ko run karke dekhein; aapko niche diya gaya output milega:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 1.21s
     Running `target/debug/functions`
The value of x is: 5
```

`another_function` ke declaration mein `x` naam ka ek parameter hai. `x` ka type `i32` specify kiya gaya hai. Jab hum `another_function` mein `5` pass karte hain, toh `println!` macro format string mein jahan curly brackets ke andar `x` tha, wahan `5` rakh deta hai.

---

## 3. Multiple parameter function

Function signatures mein, aapko har parameter ka type declare karna _zaroori (must)_ hai. Yeh Rust ke design mein ek jaan-boojhkar liya gaya faisla hai: Function definitions mein type annotations ko zaroori banane ka matlab hai ki compiler ko code mein kahin aur iski zaroorat lagbhag kabhi nahi padti taaki woh samajh sake ki aapka matlab kis type se hai. Agar compiler ko pata hota hai ki function kis type ki ummeed kar raha hai, toh woh zyada madadgar error messages bhi de pata hai.

Jab aap multiple parameters define karte hain, toh parameter declarations ko commas se alag karein, kuch is tarah:

**Filename: `src/main.rs`**

```rust
fn main() {
    print_labeled_measurement(5,'h');
}

fn print_labeled_measurement(value: i32,unit_lable:char) {
    println!("The measuremnt is: {value}{unit_lable}");
}
```

Yeh example `print_labeled_measurement` naam ka ek function banata hai jismein do parameters hain. Pehle parameter ka naam `value` hai aur yeh ek `i32` hai. Doosre ka naam `unit_label` hai aur yeh `char` type ka hai. Phir function aisa text print karta hai jismein `value` aur `unit_label` dono shamil hote hain.

Chaliye is code ko run karke dekhte hain. Apne _functions_ project ki _src/main.rs_ file mein maujood program ko upar diye gaye example se replace karein aur `cargo run` ka istemal karke ise run karein:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/functions`
The measurement is: 5h
```

Kyunki humne function ko `value` ke liye `5` aur `unit_label` ke liye `'h'` ki value ke saath call kiya tha, isliye program ke output mein woh values shamil hain.

---

## 4. Statements and Expressions

Function bodies _statements_ ki ek series se bani hoti hain jo optionally ek _expression_ par khatam hoti hain. Abhi tak humne jitne bhi functions cover kiye hain unmein koi ending expression shamil nahi tha, lekin aapne ek statement ke hisse ke roop mein ek expression dekha hai. Kyunki Rust ek expression-based language hai, isliye is antar ko samajhna bahut zaroori hai. Doosri languages mein aise antar nahi hote, toh chaliye dekhte hain ki statements aur expressions kya hain aur unke antar functions ki bodies ko kaise prabhavit karte hain.

- **Statements** aise instructions hote hain jo koi action perform karte hain aur koi value return nahi karte.
- **Expressions** evaluate hokar ek resultant value dete hain.

Chaliye kuch examples dekhte hain.

Humne darasal statements aur expressions ka istemal pehle hi kar liya hai. `let` keyword ke saath ek variable banana aur use ek value assign karna ek statement hai. Listing 3-1 mein, `let y = 6;` ek statement hai.

**Filename: `src/main.rs`**

```rust
fn main() {
    let y = 6;
    println!("Output of let statement: {y}");
}
```

_Listing 3-1: Ek `main` function declaration jismein ek statement shamil hai_

Function definitions bhi statements hoti hain; upar diya gaya poora example apne aap mein ek statement hai. (Jaisa ki hum jald hi dekhenge, kisi function ko _call_ karna ek statement nahi hai.)

Statements values return nahi karte. Isliye, aap ek `let` statement ko kisi doosre variable ko assign nahi kar sakte, jaisa ki niche diya gaya code karne ki koshish karta hai; aapko ek error milega:

**Filename: `src/main.rs`**

_(This code does not compile!)_

```rust
fn main() {
    // thi code does note complie!
    let x = (let y = 6);
}
```

Jab aap is program ko run karenge, toh aapko kuch is tarah ka error dikhega:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
error: expected expression, found `let` statement
 --> src/main.rs:2:14
  |
2 |     let x = (let y = 6);
  |              ^^^
  |
  = note: only supported directly in conditions of `if` and `while` expressions

warning: unnecessary parentheses around assigned value
 --> src/main.rs:2:13
  |
2 |     let x = (let y = 6);
  |             ^         ^
  |
  = note: `#[warn(unused_parens)]` (part of `#[warn(unused)]`) on by default
help: remove these parentheses
  |
2 -     let x = (let y = 6);
2 +     let x = let y = 6;
  |

warning: `functions` (bin "functions") generated 1 warning
error: could not compile `functions` (bin "functions") due to 1 previous error; 1 warning emitted
```

`let y = 6` statement koi value return nahi karta, isliye `x` ke bind hone ke liye wahan kuch bhi nahi hai. Yeh usse alag hai jo doosri languages, jaise C aur Ruby mein hota hai, jahan assignment khud assignment ki value return karta hai. Un languages mein, aap `x = y = 6` likh sakte hain aur `x` aur `y` dono ki value `6` ho jati hai; Rust mein aisa nahi hota.

Expressions evaluate hokar ek value dete hain aur aap Rust mein jo baaki code likhenge uska zyada-tar hissa yahi hote hain. Ek math operation par vichar karein, jaise `5 + 6`, jo ek expression hai aur evaluate hokar `11` value deta hai. Expressions statements ka hissa ho sakte hain: Listing 3-1 mein, statement `let y = 6;` mein `6` ek expression hai jo evaluate hokar `6` value deta hai. Kisi function ko call karna ek expression hai. Kisi macro ko call karna ek expression hai. Curly brackets ke saath banaya gaya ek naya scope block ek expression hai, udaharan ke liye:

**Filename: `src/main.rs`**

```rust
fn main() {
    let a = {
        let b = 3;
        b + 1
    };
    println!("The value of a is: {a}");
}
```

Yeh expression:

```rust
{
    let b = 3;
    b + 1
}
```

ek block hai jo is case mein evaluate hokar `4` deta hai. Woh value `let` statement ke hisse ke roop mein `y` se bind ho jati hai. `x + 1` line par dhyaan dein jiske ant mein semicolon nahi hai, jo un zyada-tar lines se alag hai jo aapne abhi tak dekhi hain. Expressions ke ant mein semicolons nahi hote. Agar aap kisi expression ke ant mein semicolon laga dete hain, toh aap use ek statement mein badal dete hain, aur phir woh koi value return nahi karega. Agle section mein function return values aur expressions ko explore karte waqt is baat ko dhyaan mein rakhein.

---

## 5. Function with return value

Functions unhe call karne wale code ko values return kar sakte hain. Hum return values ko naam nahi dete, lekin humein ek arrow (`->`) ke baad unka type declare karna zaroori hota hai. Rust mein, function ki return value function body ke block ke aakhiri expression ki value ke barabar (synonymous) hoti hai. Aap `return` keyword ka istemal karke aur ek value specify karke kisi function se jaldi (early) return kar sakte hain, lekin zyada-tar functions aakhiri expression ko implicitly return karte hain. Yahan ek function ka example hai jo ek value return karta hai:

**Filename: `src/main.rs`**

```rust
fn main() {
    let e = five();
    println!("The value of e is: {e}");
}

fn five() -> i32 {
    5
}
```

`five` function mein koi function calls, macros, ya `let` statements bhi nahi hain—bas akela number `5` hai. Rust mein yeh ek bilkul valid function hai. Dhyaan dein ki function ka return type bhi `-> i32` ke roop mein specify kiya gaya hai. Is code ko run karke dekhein; output kuch aisa dikhna chahiye:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.30s
     Running `target/debug/functions`
The value of x is: 5
```

`five` mein jo `5` hai woh function ki return value hai, isliye return type `i32` hai. Chaliye iski aur gehrai se jaanch karte hain. Ismein do important baatein hain: Pehli, `let x = five();` line dikhati hai ki hum ek variable ko initialize karne ke liye function ki return value ka istemal kar rahe hain. Kyunki `five` function `5` return karta hai, isliye woh line niche di gayi line ke barabar hi hai:

```rust
let e = 5;
```

Doosri, `five` function mein koi parameters nahi hain aur yeh return value ka type define karta hai, lekin function ki body mein bina semicolon ke ek akela `5` hai kyunki yeh ek expression hai jiski value hum return karna chahte hain.

Chaliye ek aur example dekhte hain:

**Filename: `src/main.rs`**

```rust
fn main() {
    let f = plus_one(5);
    println!("The value of f is: {f}");
}

fn plus_one(x:i32) -> i32 {
    x + 1
}
```

Is code ko run karne par `The value of x is: 6` print hoga. Lekin kya hoga agar hum `x + 1` wali line ke ant mein ek semicolon laga dein, aur ise ek expression se statement mein badal dein?

**Filename: `src/main.rs`**

_(This code does not compile!)_

```rust
fn main() {
    let f = plus_one(5);
    println!("The value of f is: {f}");
}

fn plus_one(x: i32) -> i32 {
    x + 1;
}
```

Is code ko compile karne par ek error aayega, jo is prakar hai:

```plaintext
$ cargo run
   Compiling functions v0.1.0 (file:///projects/functions)
error[E0308]: mismatched types
 --> src/main.rs:7:24
  |
7 | fn plus_one(x: i32) -> i32 {
  |    --------            ^^^ expected `i32`, found `()`
  |    |
  |    implicitly returns `()` as its body has no tail or `return` expression
8 |     x + 1;
  |          - help: remove this semicolon to return this value

For more information about this error, try `rustc --explain E0308`.
error: could not compile `functions` (bin "functions") due to 1 previous error
```

Main error message, `mismatched types`, is code ki mukhya samasya ko ujagar karta hai. Function `plus_one` ki definition kehti hai ki yeh ek `i32` return karega, lekin statements evaluate hokar koi value nahi dete, jise `()`, yaani unit type ke dwara vyakt kiya jata hai. Isliye, kuch bhi return nahi hota, jo function definition ke vipreet (contradicts) hai aur iske parinam-swaroop error aata hai. Is output mein, Rust is samasya ko theek karne mein madad ke liye ek message deta hai: Yeh semicolon hatane ka sujhav deta hai, jisse yeh error theek ho jayega.
