# We will cover in this lesson

1. if Expressions
2. Handling Multiple Conditions with else if
3. Using if in a let Statement

---

## Control Flow

Kisi condition ke true hone par kuch code run karne ki ability aur kisi condition ke true rehte hue kuch code ko baar-baar run karne ki ability, zyada tar programming languages ke basic building blocks hote hain. Rust code ke execution flow ko control karne wale sabse common constructs `if` expressions aur loops hain.

---

## 1. if Expressions

Ek `if` expression aapko aapke code ko conditions ke basis par branch karne deta hai. Aap ek condition dete hain aur phir kehte hain, “Agar yeh condition poori hoti hai, toh yeh block of code run karo. Agar condition poori nahi hoti hai, toh yeh block of code run mat karo.”

Apne projects directory mein `branches` naam ka ek naya project banayein taaki `if` expression ko explore kiya ja sake. _src/main.rs_ file mein, neeche diya gaya code input karein:

**Filename: `src/main.rs`**

```rust
fn main() {
    let number = 3;

    if number < 5 {
        println!("condition was true ");
    } else {
        println!("condition was false");
    }
}
```

Saare `if` expressions `if` keyword se shuru hote hain, jiske baad ek condition aati hai. Is case mein, condition check karti hai ki kya variable `number` ki value 5 se kam hai ya nahi. Hum woh block of code rakhte hain jo condition true hone par execute karna hai, condition ke turant baad curly brackets ke andar. `if` expressions mein conditions se jude code blocks ko kabhi-kabhi _arms_ bhi kaha jaata hai, bilkul waise hi jaise `match` expressions ke arms hote hain jinhe humne Chapter 2 ke “Comparing the Guess to the Secret Number” section mein discuss kiya tha.

Optionally, hum ek `else` expression bhi include kar sakte hain, jo humne yahan choose kiya hai, taaki program ko ek alternative block of code diya ja sake jo condition ke false evaluate hone par execute ho. Agar aap `else` expression nahi dete hain aur condition false hoti hai, toh program bas `if` block ko skip kar dega aur agle code par badh jayega.

Is code ko run karke dekhein; aapko neeche diya gaya output dikhna chahiye:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
condition was true
```

Aaiye `number` ki value ko aisi value mein badal kar dekhte hain jo condition ko false bana de, aur dekhte hain kya hota hai:

```rust
let number = 7;
```

Program ko dobara run karein, aur output dekhein:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
condition was false
```

Yeh bhi dhyan dene layak hai ki is code mein condition ek `bool` honi hi chahiye. Agar condition `bool` nahi hai, toh humein ek error milega. Example ke liye, neeche diye gaye code ko run karke dekhein:

**Filename: `src/main.rs`**

```rust
fn main() {
    let number = 3;

    if number {
        println!("number was three");
    }
}
```

Is baar `if` condition `3` ki value par evaluate hoti hai, aur Rust ek error throw karta hai:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
error[E0308]: mismatched types
 --> src/main.rs:4:8
  |
4 |     if number {
  |        ^^^^^^ expected `bool`, found integer

For more information about this error, try `rustc --explain E0308`.
error: could not compile `branches` (bin "branches") due to 1 previous error
```

Yeh error batata hai ki Rust ek `bool` expect kar raha tha lekin use ek integer mila. Ruby aur JavaScript jaisi languages ke vipreet, Rust automatically non-Boolean types ko Boolean mein convert karne ki koshish nahi karega. Aapko hamesha explicitly `if` ko uski condition ke roop mein ek Boolean dena hoga. Agar hum chahte hain ki `if` code block tabhi run ho jab koi number 0 ke barabar na ho, toh hum `if` expression ko is tarah badal sakte hain:

**Filename: `src/main.rs`**

```rust
fn main() {
    let number = 3;

    if number !=0 {
        println!("number was something other than zero");
    }
}
```

Is code ko run karne par `number was something other than zero` print hoga.

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
number was something other than zero
```

---

## 2. Handling Multiple Conditions with else if

Aap `if` aur `else` ko ek `else if` expression mein combine karke multiple conditions ka use kar sakte hain. Example ke liye:

**Filename: `src/main.rs`**

```rust
fn main() {
    let n = 6;

    if n % 4 == 0 {
        println!("Number is divisible by 4");
    } else if n % 3 == 0 {
        println!("number is dibisible by 3");
    } else if n % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4,3, or 2");
    }
}
```

Is program ke paas chaar possible paths hain jo yeh le sakta hai. Ise run karne ke baad, aapko neeche diya gaya output dikhna chahiye:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
number is dibisible by 3
```

Jab yeh program execute hota hai, toh yeh har `if` expression ko sequence mein check karta hai aur pehla body execute karta hai jiske liye condition true evaluate hoti hai. Dhyan dein ki bhale hi 6, 2 se divisible hai, humein output `number is divisible by 2` nahi dikhta, aur na hi humein `else` block se `number is not divisible by 4, 3, or 2` text dikhta hai. Aisa isliye hai kyunki Rust sirf pehli true condition ke liye block execute karta hai, aur ek baar jab use ek mil jaati hai, toh woh baaki conditions ko check hi nahi karta.

Bahut zyada `else if` expressions ka use karna aapke code ko clutter kar sakta hai, isliye agar aapke paas ek se zyada hain, toh aap apne code ko refactor karna chah sakte hain. Chapter 6 ek powerful Rust branching construct ke baare mein batata hai jise `match` kehte hain, jo in cases ke liye hai.

---

## 3. Using if in a let Statement

Kyunki `if` ek expression hai, hum ise `let` statement ke right side par use kar sakte hain taaki outcome ko ek variable mein assign kiya ja sake, jaisa ki Listing 3-2 mein dikhaya gaya hai.

**Filename: `src/main.rs`**

```rust
fn main() {
    let condition = true;
    let a = if condition { 5 } else { 6 };

    println!("The value of a is: {a}");
}
```

_Listing 3-2: Assigning the result of an `if` expression to a variable_

`number` variable ek value se bind ho jayega jo `if` expression ke outcome par based hogi. Is code ko run karein aur dekhein kya hota hai:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.30s
     Running `target/debug/branches`
The value of a is: 5
```

Yaad rakhein ki code ke blocks unke andar ke last expression par evaluate hote hain, aur numbers khud bhi expressions hote hain. Is case mein, poore `if` expression ki value is baat par depend karti hai ki code ka kaunsa block execute hota hai. Iska matlab hai ki `if` ke har arm se result ke roop mein aane wali values ka type same hona chahiye; Listing 3-2 mein, `if` arm aur `else` arm dono ke results `i32` integers the. Agar types mismatched hain, jaisa ki neeche diye gaye example mein hai, toh humein ek error milega:

**Filename: `src/main.rs`**

```rust
fn main() {
    let condition = true;

    let number = if condition { 5 } else { "six" };

    println!("The value of number is: {number}");
}
```

Jab hum is code ko compile karne ki koshish karte hain, toh humein ek error milega. `if` aur `else` arms mein aise value types hain jo incompatible hain, aur Rust bilkul sahi jagah batata hai ki program mein problem kahan hai:

```shell
$ cargo run
   Compiling branches v0.1.0 (file:///projects/branches)
error[E0308]: `if` and `else` have incompatible types
 --> src/main.rs:4:44
  |
4 |     let number = if condition { 5 } else { "six" };
  |                                 -          ^^^^^ expected integer, found `&str`
  |                                 |
  |                                 expected because of this

For more information about this error, try `rustc --explain E0308`.
error: could not compile `branches` (bin "branches") due to 1 previous error
```

`if` block mein expression ek integer par evaluate hota hai, aur `else` block mein expression ek string par evaluate hota hai. Yeh kaam nahi karega kyunki variables ka ek single type hona zaroori hai, aur Rust ko compile time par definitively pata hona chahiye ki `number` variable ka type kya hai. `number` ka type jaanne se compiler har jagah jahan hum `number` use karte hain, wahan type ke valid hone ko verify kar paata hai. Rust aisa nahi kar paata agar `number` ka type sirf runtime par decide hota; compiler zyada complex hota aur code ke baare mein kam guarantees de paata agar use kisi bhi variable ke liye multiple hypothetical types ka hisab rakhna padta.
