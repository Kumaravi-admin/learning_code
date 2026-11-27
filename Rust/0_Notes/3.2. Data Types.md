# We will cover in this lesson

- Data types ka overview aur Rust ka statically typed hona
- Type annotations kab zaroori hote hain (`parse()` ka example)
- Scalar types: integers, floating-point, Booleans, characters
- Integer types (signed/unsigned, `isize`/`usize`) aur integer literals
- Integer overflow (debug vs release mode) aur `wrapping_*`, `checked_*`, `overflowing_*`, `saturating_*` methods
- Floating-point types (`f32`, `f64`)
- Numeric operations (addition, subtraction, multiplication, division, remainder)
- Boolean type (`bool`)
- Character type (`char`) aur Unicode scalar values
- Compound types: tuples aur arrays
- Tuple: destructuring, index se access, aur unit type `()`
- Array: declare karna, type likhna, `[3; 5]` syntax
- Array element access aur invalid index par panic (memory safety)

---

## Data Types

Rust mein har ek value ka ek khaas data type hota hai, jo Rust ko batata hai ki kis tarah ka data specify kiya gaya hai taaki woh jaan sake ki uss data ke saath kaise kaam karna hai. Hum do data type subsets dekhenge: scalar aur compound.

Dhyan rakhein ki Rust ek statically typed language hai, jiska matlab hai ki ise compile time par sabhi variables ke types pata hone chahiye. Compiler aam taur par value aur hum use kaise istemal karte hain, iske aadhar par infer (andaza) laga sakta hai ki hum kaunsa type use karna chahte hain. Aise cases mein jahan kai types sambhav hain, jaise jab humne Chapter 2 ke “Comparing the Guess to the Secret Number” section mein parse ka use karke ek String ko numeric type mein convert kiya tha, toh humein ek type annotation add karna zaroori hota hai, kuch is tarah:

```rust
let guess: u32 = "42".parse().expect("Not a number!");
println!("{}",guess);
```

Agar hum upar dikhaye gaye code mein `: u32` type annotation add nahi karte hain, toh Rust niche diya gaya error dikhayega, jiska matlab hai ki compiler ko humse aur jaankari chahiye taaki woh jaan sake ki hum kaunsa type use karna chahte hain:

```plaintext
$ cargo build
   Compiling no_type_annotations v0.1.0 (file:///projects/no_type_annotations)
error[E0284]: type annotations needed
 --> src/main.rs:2:9
  |
2 |     let guess = "42".parse().expect("Not a number!");
  |         ^^^^^        ----- type must be known at this point
  |
  = note: cannot satisfy `<_ as FromStr>::Err == _`
help: consider giving `guess` an explicit type
  |
2 |     let guess: /* Type */ = "42".parse().expect("Not a number!");
  |              ++++++++++++

For more information about this error, try `rustc --explain E0284`.
error: could not compile `no_type_annotations` (bin "no_type_annotations") due to 1 previous error
```

Aap doosre data types ke liye alag-alag type annotations dekhenge.

---

### Scalar Types

Ek scalar type ek single value ko represent karta hai. Rust mein chaar primary scalar types hain: integers, floating-point numbers, Booleans, aur characters. Aap inhe doosri programming languages se pehchante honge. Chaliye dekhte hain ki yeh Rust mein kaise kaam karte hain.

#### Integer Types

Ek integer bina fractional (dashamlav) component wala number hota hai. Humne Chapter 2 mein ek integer type use kiya tha, `u32` type. Yeh type declaration batata hai ki isse judi value ek unsigned integer honi chahiye (signed integer types `u` ke bajaye `i` se shuru hote hain) jo 32 bits ki jagah leta hai. Table 3-1 Rust mein built-in integer types dikhati hai. Hum integer value ka type declare karne ke liye inmein se kisi bhi variant ka istemal kar sakte hain.

**Table 3-1: Rust mein Integer Types**

| Length                 | Signed  | Unsigned |
| ---------------------- | ------- | -------- |
| 8-bit                  | `i8`    | `u8`     |
| 16-bit                 | `i16`   | `u16`    |
| 32-bit                 | `i32`   | `u32`    |
| 64-bit                 | `i64`   | `u64`    |
| 128-bit                | `i128`  | `u128`   |
| Architecture-dependent | `isize` | `usize`  |

Har variant ya toh signed ho sakta hai ya unsigned aur uska ek explicit size hota hai. Signed aur unsigned is baat ko darshate hain ki kya number ka negative hona sambhav hai—doosre shabdon mein, kya number ke saath sign hona zaroori hai (signed) ya phir yeh hamesha positive hi rahega aur isliye bina sign ke represent kiya ja sakta hai (unsigned). Yeh kagaz par numbers likhne jaisa hai: Jab sign mayne rakhta hai, toh number ko plus sign ya minus sign ke saath dikhaya jaata hai; halanki, jab yeh maan lena safe ho ki number positive hai, toh use bina kisi sign ke dikhaya jaata hai.

Signed numbers ko two’s complement representation ka use karke store kiya jaata hai. Har signed variant −(2ⁿ⁻¹) se lekar 2ⁿ⁻¹ − 1 tak (inclusive) ke numbers store kar sakta hai, jahan _n_ un bits ki sankhya hai jo woh variant use karta hai. Toh, ek `i8` −(2⁷) se 2⁷ − 1 tak ke numbers store kar sakta hai, jo −128 se 127 ke barabar hota hai. Unsigned variants 0 se 2ⁿ − 1 tak ke numbers store kar sakte hain, isliye ek `u8` 0 se 2⁸ − 1 tak ke numbers store kar sakta hai, jo 0 se 255 ke barabar hota hai.

Iske alawa, `isize` aur `usize` types aapke computer ke architecture par nirbhar karte hain jis par aapka program chal raha hai: agar aap 64-bit architecture par hain toh 64 bits aur agar aap 32-bit architecture par hain toh 32 bits.

Aap Table 3-2 mein dikhaye gaye kisi bhi roop mein integer literals likh sakte hain. Dhyan dein ki jo number literals multiple numeric types ke ho sakte hain, unmein type batane ke liye type suffix lagane ki anumati hoti hai, jaise `57u8`. Number literals ko padhne mein aasan banane ke liye `_` ka istemal visual separator ke roop mein bhi kiya ja sakta hai, jaise `1_000`, jiski value bilkul wahi hogi jaise aapne `1000` likha ho.

**Table 3-2: Rust mein Integer Literals**

| Number literals  | Example       |
| ---------------- | ------------- |
| Decimal          | `98_222`      |
| Hex              | `0xff`        |
| Octal            | `0o77`        |
| Binary           | `0b1111_0000` |
| Byte (`u8` only) | `b'A'`        |

Toh aapko kaise pata chalega ki kis type ka integer use karna hai? Agar aap sure nahi hain, toh shuruat karne ke liye Rust ke defaults aam taur par acche hote hain: Integer types default roop se `i32` hote hain. Woh primary sthiti jismein aap `isize` ya `usize` ka use karenge, woh tab hoti hai jab aap kisi collection ki indexing kar rahe hon.

#### Integer Overflow

Maan lijiye aapke paas `u8` type ka ek variable hai jo 0 se 255 ke beech ki values rakh sakta hai. Agar aap variable ko uss range ke bahar ki value par badalne ki koshish karte hain, jaise 256, toh integer overflow hoga, jiske parinam-swaroop do behaviors mein se ek ho sakta hai.

Jab aap debug mode mein compile kar rahe hote hain, toh Rust mein integer overflow ke liye checks shamil hote hain jo is behavior ke hone par aapke program ko runtime par panic kar dete hain. Jab koi program error ke saath exit hota hai toh Rust _panicking_ shabd ka istemal karta hai; hum Chapter 9 ke “Unrecoverable Errors with `panic!`” section mein panics par aur gehrai se charcha karenge.

Jab aap `--release` flag ke saath release mode mein compile kar rahe hote hain, toh Rust mein integer overflow ke woh checks shamil nahi hote jo panics ka kaaran bante hain. Iske bajaye, agar overflow hota hai, toh Rust _two’s complement wrapping_ karta hai. Sanchep mein, type jitni maximum value rakh sakta hai usse badi values "wrap around" hokar uss minimum value par aa jaati hain jo woh type rakh sakta hai. `u8` ke maamle mein, value 256 ban jaati hai 0, value 257 ban jaati hai 1, aur aise hi aage chalta hai. Program panic nahi karega, lekin variable mein ek aisi value hogi jo shayad woh nahi hai jiski aap umeed kar rahe the. Integer overflow ke wrapping behavior par bharosa karna ek error maana jaata hai.

Overflow ki sambhavna ko explicitly handle karne ke liye, aap primitive numeric types ke liye standard library dwara diye gaye methods ki in families ka istemal kar sakte hain:

- `wrapping_*` methods, jaise `wrapping_add`, ke saath sabhi compilation modes mein wrap karein.
- `checked_*` methods ke saath agar overflow ho toh `None` value return karein.
- `overflowing_*` methods ke saath value aur ek Boolean return karein jo yeh batata hai ki overflow hua tha ya nahi.
- `saturating_*` methods ke saath value ki minimum ya maximum values par saturate (rok) dein.

#### Floating-Point Types

Rust mein floating-point numbers ke liye bhi do primitive types hain, jo decimal points (dashamlav) wale numbers hote hain. Rust ke floating-point types `f32` aur `f64` hain, jinka size kramashah 32 bits aur 64 bits hota hai. Default type `f64` hai kyunki modern CPUs par iski speed lagbhag `f32` ke barabar hi hoti hai lekin yeh zyada precision (sateekta) dene mein saksham hai. Sabhi floating-point types signed hote hain.

Yahan ek example hai jo floating-point numbers ko kaam karte hue dikhata hai:

**Filename: `src/main.rs`**

```rust
fn main() {
    let x = 2.0;   // f64
    let y: f32 = 3.0; // f32
    println!("Output of x: {}",x);
    println!("Output of y: {}",y);
}
```

Floating-point numbers ko IEEE-754 standard ke anusaar represent kiya jaata hai.

#### Numeric Operations

Rust sabhi number types ke liye un basic mathematical operations ko support karta hai jinki aap umeed karte hain: addition, subtraction, multiplication, division, aur remainder. Integer division zero ki taraf nearest integer tak truncate (kat) ho jaata hai. Niche diya gaya code dikhata hai ki aap ek `let` statement mein har numeric operation ka istemal kaise karenge:

\*\*Filename:
