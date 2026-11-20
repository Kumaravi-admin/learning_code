# The Rust Programming Language

## Variables and Mutability

Jaise ki “Storing Values with Variables” section mein mention kiya gaya tha, **by default, variables immutable hote hain**. Yeh Rust dwara diye jaane wale kai nudges mein se ek hai jo aapko apna code is tarah se likhne ke liye encourage karta hai jo Rust dwara offer kiye jaane wale safety aur easy concurrency ka advantage leta hai. Halanki, aapke paas abhi bhi apne variables ko mutable banane ka option hota hai. Aaiye explore karein ki kaise aur kyun Rust aapko immutability ko favor karne ke liye encourage karta hai aur kyun kabhi-kabhi aap isse opt out karna chahenge.

Jab koi variable immutable hota hai, toh ek baar jab koi value kisi name se bind ho jati hai, aap us value ko change nahi kar sakte. Isko illustrate karne ke liye, `cargo new variables` ka use karke apne projects directory mein `variables` naam ka ek naya project generate karein.

Phir, apne naye `variables` directory mein, `src/main.rs` open karein aur iske code ko niche diye gaye code se replace karein, jo abhi compile nahi hoga:

**Filename: `src/main.rs`**

```rust
fn main() {
    let x = 5;
    println!("The value of x is: {x}");
    x = 6;
    println!("The value of x is: {x}");
}
```

`cargo run` ka use karke program ko save aur run karein. Aapko ek immutability error ke regarding ek error message milna chahiye, jaisa ki is output mein dikhaya gaya hai:

**Output:**

```text
$ cargo run
   Compiling variables v0.1.0 (file:///projects/variables)
error[E0384]: cannot assign twice to immutable variable `x`
 --> src/main.rs:4:5
  |
2 |     let x = 5;
  |         - first assignment to `x`
3 |     println!("The value of x is: {x}");
4 |     x = 6;
  |     ^^^^^ cannot assign twice to immutable variable
  |
help: consider making this binding mutable
  |
2 |     let mut x = 5;
  |         +++

For more information about this error, try `rustc --explain E0384`.
error: could not compile `variables` (bin "variables") due to 1 previous error
```

Yeh example dikhata hai ki compiler aapke programs mein errors find karne mein kaise madad karta hai. Compiler errors frustrating ho sakte hain, lekin sach mein inka matlab sirf yeh hota hai ki aapka program abhi safely woh nahi kar raha hai jo aap chahte hain; inka matlab yeh nahi hai ki aap ek acche programmer nahi hain! Experienced Rustaceans ko bhi abhi compiler errors aate hain.

Aapko error message `cannot assign twice to immutable variable x` isliye mila kyunki aapne immutable `x` variable ko doosri value assign karne ki koshish ki thi.

Yeh important hai ki jab hum kisi aise value ko change karne ki koshish karein jo immutable designate ki gayi hai, toh humein compile-time errors milen, kyunki yeh situation bugs ki taraf le ja sakti hai. Agar hamare code ka ek hissa is assumption par kaam karta hai ki ek value kabhi change nahi hogi aur hamare code ka doosra hissa us value ko change kar deta hai, toh yeh possible hai ki code ka pehla hissa woh na kare jo iske design ke mutabiq karna chahiye tha. Is tarah ke bug ka cause baad mein track down karna difficult ho sakta hai, khaas kar tab jab doosra piece of code value ko sirf kabhi-kabhi change karta hai. Rust compiler yeh guarantee karta hai ki jab aap yeh state karte hain ki ek value change nahi hogi, toh yeh sach mein change nahi hogi, isliye aapko ise khud track karke rakhne ki zaroorat nahi padti hai. Is tarah aapka code samajhna aasaan ho jata hai.

Lekin mutability bohot useful ho sakte hain aur code ko likhne mein zyada convenient bana sakte hain. Halanki variables by default immutable hote hain, aap variable name ke samne `mut` add karke unhein mutable bana sakte hain jaisa ki aapne Chapter 2 mein kiya tha. `mut` add karna code ke future readers ko intent bhi convey karta hai yeh indicate karke ki code ke doosre parts is variable ki value ko change kar rahe honge.

For example, aaiye `src/main.rs` ko change karke yeh karein:

**Filename: `src/main.rs`**

```rust
fn main() {
    let mut x = 5;
    println!("The value of x is: {x}");
    x = 6;
    println!("The value of x is: {x}");
}
```

Jab hum program ko ab run karte hain, toh humein yeh milta hai:

**Output:**

```text
\$ cargo run
   Compiling variables v0.1.0 (file:///projects/variables)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.30s
     Running `target/debug/variables`
The value of x is: 5
The value of x is: 6
```

Humein `x` se bound value ko 5 se 6 change karne ki permission mil jati hai jab `mut` use kiya jata hai. Ultimately, yeh decide karna ki mutability use karni hai ya nahi, aap par depend karta hai aur is baat par depend karta hai ki aapko us particular situation mein kya sabse clearest lagta hai.

---

### Declaring Constants

Immutable variables ki tarah, **constants** aisi values hoti hain jo ek name se bound hoti hain aur change hone ki permission nahi hoti, lekin constants aur variables ke beech kuch differences hote hain:

- **No `mut` Allowed:** Pehle, aapko constants ke sath `mut` use karne ki permission nahi hoti. Constants sirf default by immutable nahi hote—woh hamesha immutable hote hain.
- **`const` Keyword & Type Annotation:** Aap `let` keyword ki jagah `const` keyword ka use karke constants declare karte hain, aur value ka type _must_ annotated hona chahiye. Hum types aur type annotations ko next section, “Data Types” mein cover karenge, isliye abhi details ke baare mein chinta na karein. Bas yeh jaanein ki aapko hamesha type annotate karna hoga.
- **Any Scope:** Constants kisi bhi scope mein declare kiye ja sakte hain, jisme global scope bhi shamil hai, jo unhein aisi values ke liye useful banata hai jiske baare mein code ke kai parts ko pata hona zaroori hota hai.
- **Constant Expression Only:** Last difference yeh hai ki constants ko sirf ek constant expression par set kiya ja sakta hai, na ki kisi aise value ke result par jo sirf runtime par compute ki ja sakti ho.

Yahan ek constant declaration ka example hai:

```rust
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
```

Constant ka naam `THREE_HOURS_IN_SECONDS` hai, aur iski value `60` (ek minute mein seconds ki sankhya) ko `60` (ek ghante mein minutes ki sankhya) se aur phir `3` (is program mein count karne ke ghante) se multiply karne ke result par set ki gayi hai. Constants ke liye Rust ki naming convention yeh hai ki words ke beech mein underscores ke sath saare uppercase letters ka use kiya jaye. Compiler compile time par operations ke ek limited set ko evaluate karne mein saksham hota hai, jo humein is value ko ek aise tareeqe se likhne ki choice deta hai jo samajhne aur verify karne mein aasaan ho, na ki is constant ko value 10,800 par set karne ke. Constants declare karte waqt kin operations ka use kiya ja sakta hai, iske baare mein zyada information ke liye Rust Reference’s section on constant evaluation dekhein.

Constants us poore time ke liye valid hote hain jab tak ek program run hota hai, us scope ke andar jisme unhein declare kiya gaya tha. Yeh property constants ko aapke application domain ki un values ke liye useful banati hai jiske baare mein program ke multiple parts ko pata hona pad sakta hai, jaise ki kisi game ke kisi player ko earn karne ki allowed maximum points, ya light ki speed.

Poore program mein use kiye jaane wale hardcoded values ko constants ke taur par naam dena code ke future maintainers ko us value ka meaning convey karne mein useful hota hai. Yeh code mein sirf ek hi jagah rakhne mein bhi madad karta hai jise future mein hardcoded value update karne ki zaroorat padne par change karne ki zaroorat padegi.

---

### Shadowing

Jaise ki aapne Chapter 2 mein guessing game tutorial mein dekha tha, aap pehle wale variable ke same name ke sath ek naya variable declare kar sakte hain. Rustaceans yeh kehte hain ki pehla variable doosre dwara **shadowed** ho gaya hai, jiska matlab yeh hai ki jab aap variable ke name ka use karenge toh doosra variable wahi hoga jo compiler ko dikhega. Effectively, doosra variable pehle wale ko overshadow kar deta hai, variable name ke saare uses ko apne upar le leta hai jab tak ya toh khud woh shadowed na ho jaye ya scope khatam na ho jaye.

Hum same variable ke name ka use karke aur `let` keyword ka use repeat karke kisi variable ko shadow kar sakte hain, is tarah:

**Filename: `src/main.rs`**

```rust
fn main() {
    let x = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");
}
```
