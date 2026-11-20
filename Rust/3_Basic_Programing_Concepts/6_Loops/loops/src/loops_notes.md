# We will cover in this lesson

1. loop se code repeat karna
2. Loop se value return karna (break ke saath)
3. Loop labels (nested loops mein bahar wala loop todna)
4. while loop (condition wala loop)
5. for loop (collection par loop)

# Repetition with Loops

Code ke ek block ko ek se zyada baar execute karna aksar useful hota hai.

Is task ke liye, Rust kai loops provide karta hai, jo loop body ke andar ke code ko shuru se aakhir tak run karenge aur phir turant shuru se dobara start karenge.

Loops ke saath experiment karne ke liye, aaiye `loops` naam ka ek naya project banayein. Rust mein teen tarah ke loops hain: `loop`, `while`, aur `for`.

Aaiye har ek ko try karein.

## Repeating Code with `loop`

`loop` keyword Rust ko batata hai ki code ke ek block ko baar-baar hamesha ke liye execute kare ya tab tak jab tak aap explicitly use rukne ke liye na kahein.

Ek example ke roop mein, apne `loops` directory mein `src/main.rs` file ko aisa dikhne ke liye badlein:

**Filename:** `src/main.rs`

```rust
fn main() {
    loop {
        println!("again!");
    }
}
```

Jab hum is program ko run karte hain, toh humein screen par `again!` baar-baar lagatar print hota hua dikhega jab tak hum program ko manually stop nahi karte.

Zyada tar terminals program ko interrupt karne ke liye keyboard shortcut `Ctrl-C` support karte hain jo ek continual loop mein phans gaya hai.

Ise try karein:

```shell
$ cargo run
   Compiling loops v0.1.0 (file:///projects/loops)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.08s
     Running `target/debug/loops`
again!
again!
again!
again!
^Cagain!
```

Symbol `^C` us jagah ko represent karta hai jahan aapne `Ctrl-C` press kiya.

Ho sakta hai ki aapko `^C` ke baad `again!` word print hota dikhe ya na dikhe, yeh is baat par depend karta hai ki interrupt signal milne par code loop mein kahan tha.

Khushkismati se, Rust code ka use karke loop se bahar nikalne ka ek tareeka bhi provide karta hai. Aap loop ke andar `break` keyword rakh kar program ko bata sakte hain ki loop execute karna kab band karna hai.

Yaad karein ki humne aisa Chapter 2 ke guessing game ke “Quitting After a Correct Guess” section mein kiya tha taaki jab user sahi number guess karke game jeet jaye toh program exit ho jaye.

Humne guessing game mein `continue` ka bhi use kiya tha, jo ek loop mein program ko batata hai ki loop ke is iteration mein bache hue kisi bhi code ko skip kare aur agle iteration par chala jaye.

## Returning Values from Loops

`loop` ka ek use kisi aise operation ko retry karna hai jiske baare mein aapko pata hai ki woh fail ho sakta hai, jaise check karna ki kisi thread ne apna kaam poora kiya hai ya nahi.

Aapko us operation ka result loop se bahar baaki code mein pass karne ki zaroorat bhi pad sakti hai.

Aisa karne ke liye, aap jis value ko return karwana chahte hain use `break` expression ke baad add kar sakte hain jiska use aap loop ko rokne ke liye karte hain; woh value loop se bahar return kar di jayegi taaki aap use use kar sakein, jaisa ki yahan dikhaya gaya hai:

```rust
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {result}");
}
```

Loop se pehle, hum `counter` naam ka ek variable declare karte hain aur use `0` se initialize karte hain.

Phir hum `result` naam ka ek variable declare karte hain jo loop se return hone wali value ko hold karega.

Loop ke har iteration par, hum `counter` variable mein `1` add karte hain, aur phir check karte hain ki kya `counter`, `10` ke barabar hai.

Jab aisa hota hai, toh hum `counter * 2` value ke saath `break` keyword ka use karte hain.

Loop ke baad, hum us statement ko end karne ke liye ek semicolon ka use karte hain jo `result` ko value assign karta hai.

Finally, hum `result` mein maujood value ko print karte hain, jo is case mein `20` hai.

Aap loop ke andar se `return` bhi kar sakte hain. Jabki `break` sirf current loop se bahar nikalta hai, `return` hamesha current function se hi bahar nikal deta hai.

## Disambiguating with Loop Labels

Agar aapke paas loops ke andar loops hain, toh `break` aur `continue` us point par sabse andar wale loop par apply hote hain.

Aap optionally ek loop par ek loop label specify kar sakte hain jiska use aap phir `break` ya `continue` ke saath yeh batane ke liye kar sakte hain ki yeh keywords innermost loop ke bajaye labeled loop par apply hon.

Loop labels ek single quote ke saath shuru hone chahiye. Yahan do nested loops ke saath ek example diya gaya hai:

```rust
fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");
}
```

Bahar wale loop ka label `'counting_up` hai, aur yeh `0` se `2` tak count up karega.

Andar wala loop bina label ke `10` se `9` tak count down karta hai.

Pehla `break` jo koi label specify nahi karta, woh sirf inner loop se exit karega. `break 'counting_up;` statement outer loop se exit karega.

Yeh code print karta hai:

```shell
$ cargo run
   Compiling loops v0.1.0 (file:///projects/loops)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.58s
     Running `target/debug/loops`
count = 0
remaining = 10
remaining = 9
count = 1
remaining = 10
remaining = 9
count = 2
remaining = 10
End count = 2
```

## Streamlining Conditional Loops with `while`

Kisi program ke liye aksar loop ke andar kisi condition ko evaluate karna useful hota hai.

Jab tak condition true rehti hai, loop chalta rehta hai. Jab condition true nahi rehti, toh program `break` call karta hai aur loop ko rok deta hai.

Is tarah ka behavior `loop`, `if`, `else`, aur `break` ke combination ka use karke implement karna possible hai; agar aap chahein toh abhi ek program mein ise try kar sakte hain.

Halanki, yeh pattern itna common hai ki Rust mein iske liye ek built-in language construct hai, jise `while` loop kehte hain.

Listing 3-3 mein, hum `while` ka use karke program ko teen baar loop karate hain, har baar countdown karte hue, aur phir loop ke baad ek message print karke exit karte hain.

**Filename:** `src/main.rs`

```rust
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1;
    }

    println!("LIFTOFF!!!");
}
```

_Listing 3-3: Using a while loop to run code while a condition evaluates to true_

Yeh construct bahut saare nesting ko khatam kar deta hai jo `loop`, `if`, `else`, aur `break` ka use karne par zaroori hota, aur yeh zyada clear hai.

Jab tak condition true evaluate hoti hai, code chalta hai; warna yeh loop se exit kar jaata hai.

## Looping Through a Collection with `for`

Aap kisi collection ke elements, jaise ki ek array, par loop chalane ke liye `while` construct ka use choose kar sakte hain.

Example ke liye, Listing 3-4 mein diya gaya loop array `a` ke har element ko print karta hai.

**Filename:** `src/main.rs`

```rust
fn main() {
    let a = [10, 20, 30, 40, 50];
    let mut index = 0;

    while index < 5 {
        println!("the value is: {}", a[index]);

        index += 1;
    }
}
```

_Listing 3-4: Looping through each element of a collection using a while loop_

Yahan, code array ke elements ke through count up karta hai.

Yeh index `0` se shuru hota hai, aur phir tab tak loop karta hai jab tak yeh array ke final index tak nahi pahunch jaata (yani jab `index < 5` ab true nahi rehta).

Is code ko run karne par array ka har element print hoga:

```shell
$ cargo run
   Compiling loops v0.1.0 (file:///projects/loops)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.32s
     Running `target/debug/loops`
the value is: 10
the value is: 20
the value is: 30
the value is: 40
the value is: 50
```

Saari paanch array values terminal mein dikhti hain, jaisa ki expect kiya gaya tha.

Bhale hi index kisi point par `5` ki value tak pahunch jayega, loop array se chhathi value fetch karne ki koshish karne se pehle hi execute hona band kar deta hai.

Halanki, yeh approach error-prone hai; agar index value ya test condition galat ho toh hum program ko panic karwa sakte hain.

Example ke liye, agar aapne `a` array ki definition ko char elements mein badal diya lekin condition ko `while index < 4` mein update karna bhool gaye, toh code panic kar jayega.

Yeh slow bhi hai, kyunki compiler loop ke through har iteration par is baat ka conditional check perform karne ke liye runtime code add karta hai ki index array ke bounds ke andar hai ya nahi.

Ek zyada concise alternative ke roop mein, aap ek `for` loop ka use kar sakte hain aur collection mein har item ke liye kuch code execute kar sakte hain.

Ek `for` loop Listing 3-5 mein diye gaye code jaisa dikhta hai.

**Filename:** `src/main.rs`

```rust
fn main() {
    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {element}");
    }
}
```

_Listing 3-5: Looping through each element of a collection using a for loop_

Jab hum is code ko run karte hain, toh humein wahi output milega jo Listing 3-4 mein mila tha.

Isse bhi zaroori baat yeh hai ki humne ab code ki safety badha di hai aur un bugs ke chance ko khatam kar diya hai jo array ke end se aage nikal jaane ya utni door tak na jaane aur kuch items ko miss kar dene se ho sakte the.

`for` loops se generate kiya gaya machine code bhi zyada efficient ho sakta hai kyunki har iteration par index ko array ki length ke saath compare karne ki zaroorat nahi hoti.

`for` loop ka use karne se, aapko kisi aur code ko badalne ki yaad rakhne ki zaroorat nahi padegi agar aap array mein values ki sankhya badalte hain, jaisa ki aap Listing 3-4 mein use kiye gaye method ke saath karte.

`for` loops ki safety aur conciseness unhe Rust mein sabse zyada use kiya jaane wala loop construct banati hai.

Yahan tak ki un situations mein bhi jahan aap kisi code ko ek certain number of times run karna chahte hain, jaisa ki Listing 3-3 mein `while` loop ka use karne wale countdown example mein tha, zyada tar Rustaceans ek `for` loop ka use karenge.

Aisa karne ka tareeka standard library dwara provide kiye gaye `Range` ka use karna hai, jo ek number se shuru hokar doosre number se pehle khatam hone wale saare numbers ko sequence mein generate karta hai.

Yahan bataya gaya hai ki ek `for` loop aur ek aur method jise humne abhi tak discuss nahi kiya hai, `rev`, ka use karke range ko reverse karne par countdown kaisa dikhega:

**Filename:** `src/main.rs`

```rust
fn main() {
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
```

Yeh code thoda zyada accha hai, hai na?
