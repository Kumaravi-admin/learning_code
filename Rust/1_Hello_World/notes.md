# The Rust Programming Language

## Hello, World!

Ab jabki aapne Rust install kar liya hai, toh yeh aapka pehla Rust program likhne ka samay hai. Nayi language seekhte waqt yeh ek tradition hai ki ek chhota sa program likha jaye jo screen par `Hello, world!` print kare, toh hum bhi yahi karenge!

> **Note:** Yeh book ye maankar chalti hai ki aapko command line ki basic jaankari hai. Rust aapke editing, tooling, ya code kahan rehta hai isko lekar koi specific demand nahi karta, isliye agar aap command line ki bajaye kisi IDE ka use karna pasand karte hain, toh bejhijhak apna favourite IDE use karein. Aajkal bahut saare IDEs mein kuch had tak Rust ka support hota hai; details ke liye IDE ki documentation check karein. Rust team `rust-analyzer` ke zariye behtareen IDE support enable karne par focus kar rahi hai. Jyada details ke liye Appendix D dekhein.

---

### Project Directory Setup

Aap apne Rust code ko store karne ke liye ek directory banakar shuruat karenge. Rust ke liye isse koi farq nahi padta ki aapka code kahan rehta hai, lekin is book ke exercises aur projects ke liye, hum yeh suggest karte hain ki aap apne home directory mein ek `projects` directory banayein aur apne saare projects wahin rakhein.

Ek terminal open karein aur `projects` directory ke andar ek `projects` directory aur “Hello, world!” project ke liye ek directory banane ke liye yeh commands enter karein.

**Linux, macOS, aur Windows par PowerShell ke liye:**

```bash
mkdir ~/projects
cd ~/projects
mkdir hello_world
cd hello_world
```

**Windows CMD ke liye:**

```cmd
mkdir "%USERPROFILE%\projects"
cd /d "%USERPROFILE%\projects"
mkdir hello_world
cd hello_world
```

---

### Rust Program Basics

Aage, ek naya source file banayein aur iska naam `main.rs` rakhein. Rust files hamesha `.rs` extension ke sath khatam hoti hain. Agar aap apne filename mein ek se zyada word use kar rahe hain, toh convention yeh hai ki unhe alag karne ke liye underscore ka use karein. Jaise ki, `helloworld.rs` ke bajaye `hello_world.rs` use karein.

Abhi jo `main.rs` file aapne banayi hai use open karein aur Listing 1-1 ka code enter karein.

**Filename: `main.rs`**

```rust
fn main() {
    println!("Hello, world!");
}
```

_Listing 1-1: Ek program jo Hello, world! print karta hai._

File ko save karein aur apne terminal window par `~/projects/hello_world` directory mein wapas jayein.

**Linux ya macOS par, file ko compile aur run karne ke liye:**

```bash
rustc main.rs
./main
```

**Output:**

```text
Hello, world!
```

**Windows par (`./main` ki bajaye `.\main` command enter karein):**

```cmd
rustc main.rs
.\main
```

**Output:**

```text
Hello, world!
```

Aapke operating system se koi farq nahi padta, string `Hello, world!` terminal par print ho jani chahiye. Agar aapko yeh output nahi dikhti, toh madad pane ke liye Installation section ke “Troubleshooting” part ko refer karein.

Agar `Hello, world!` print ho gaya, toh badhaai ho! Aapne officially ek Rust program likh liya hai. Isse aap ek Rust programmer ban gaye hain—welcome!

---

### The Anatomy of a Rust Program

Chaliye is “Hello, world!” program ko detail mein review karte hain. Yahan puzzle ka pehla tukda hai:

```rust
fn main() {

}
```

Yeh lines `main` naam ke ek function ko define karti hain. `main` function special hota hai: Yeh har executable Rust program mein chalne wala hamesha pehla code hota hai. Yahan, pehli line ek aise function ko declare karti hai jiska naam `main` hai jiska koi parameters nahi hai aur jo kuch bhi return nahi karta. Agar parameters hote, toh woh parentheses (`()`) ke andar jaate.

Function body `{}` ke andar wrapped hoti hai. Rust ko sabhi function bodies ke aas-paas curly brackets ki zaroorat hoti hai. Opening curly bracket ko function declaration ke same line par rakhna aur beech mein ek space add karna ek achhi style hai.

> **Note:** Agar aap Rust projects mein ek standard style par stick rehna chahte hain, toh aap apne code ko kisi particular style mein format karne ke liye `rustfmt` naam ke ek automatic formatter tool ka use kar sakte hain (`rustfmt` ke baare mein zyada Appendix D mein hai). Rust team ne is tool ko standard Rust distribution ke sath include kiya hai, jaise `rustc` hai, toh yeh aapke computer par pehle se installed hona chahiye!

`main` function ki body mein yeh code hota hai:

```rust
println!("Hello, world!");
```

Yeh line is chhote se program mein saara kaam karti hai: Yeh screen par text print karti hai. Yahan dhyaan dene layak teen important details hain:

1. **Macro Call (`println!`):** Pehle, `println!` ek Rust macro ko call karta hai. Agar isne iske bajaye kisi function ko call kiya hota, toh ise `println` (bina `!` ke) ke taur par likha jata. Rust macros code likhne ka ek tarika hain jo Rust syntax ko extend karne ke liye code generate karte hain, aur hum inhein Chapter 20 mein aur zyada detail mein discuss karenge. Abhi ke liye, aapko sirf yeh janne ki zaroorat hai ki `!` use karne ka matlab hai ki aap ek normal function ki bajaye ek macro call kar rahe hain aur macros hamesha functions jaise rules ko follow nahi karte hain.
2. **String Argument:** Doosra, aapko `"Hello, world!"` string dikhti hai. Hum is string ko `println!` ko ek argument ke taur par pass karte hain, aur string screen par print ho jati hai.
3. **Semicolon (`;`):** Teesra, hum line ko semicolon (`;`) ke sath end karte hain, jo yeh batata hai ki yeh expression khatam ho gaya hai, aur agla shuru hone ke liye taiyar hai. Rust code ki zyadatar lines semicolon ke sath end hoti hain.

---

### Compilation and Execution

Aapne abhi ek naya banaya hua program run kiya hai, toh aaiye process ke har step ko examine karte hain.

Rust program run karne se pehle, aapko `rustc` command enter karke aur apne source file ka naam pass karke Rust compiler ka use karke ise compile karna hoga, jaise yeh:

```bash
rustc main.rs
```

Agar aapka C ya C++ background hai, toh aap notice karenge ki yeh `gcc` ya `clang` jaisa hai. Successfully compile hone ke baad, Rust ek binary executable output karta hai.

**Linux, macOS, aur Windows par PowerShell par (`ls` command):**

```bash
ls
```

**Output:**

```text
main  main.rs
```

Linux aur macOS par, aapko do files dikhengi. Windows par PowerShell ke sath, aapko wahi teen files dikhengi jo aap CMD ka use karke dekhenge.

**Windows par CMD ke sath (`dir` command):**

```cmd
dir /B %= /B option kehta hai ki sirf file names dikhayein =%
```

**Output:**

```text
main.exe
main.pdb
main.rs
```

Yeh `.rs` extension ke sath source code file, executable file (Windows par `main.exe`, lekin baaki sabhi platforms par `main`), aur, Windows use karte waqt, `.pdb` extension ke sath debugging information contain karne wali ek file dikhata hai.

Yahan se, aap `main` ya `main.exe` file run karte hain, jaise yeh:

```bash
./main # ya Windows par .\main
```

Agar aapka `main.rs` aapka “Hello, world!” program hai, toh yeh line aapke terminal par `Hello, world!` print karti hai.

Agar aap kisi dynamic language se zyada familiar hain, jaise Ruby, Python, ya JavaScript, toh aap kisi program ko alag-alag steps mein compile aur run karne ke aadi nahi ho sakte hain. Rust ek **ahead-of-time compiled language** hai, iska matlab hai ki aap ek program compile kar sakte hain aur executable kisi aur ko de sakte hain, aur woh bina Rust install kiye bhi ise run kar sakte hain. Agar aap kisi ko `.rb`, `.py`, ya `.js` file dete hain, toh unhe ek Ruby, Python, ya JavaScript implementation install rakhni padegi (respective). Lekin un languages mein, aapko apne program ko compile aur run karne ke liye sirf ek command ki zaroorat hoti hai. Language design mein har cheez ek trade-off hoti hai.

Simple programs ke liye sirf `rustc` ke sath compile karna theek hai, lekin jaise-jaise aapka project badhega, aap saari options ko manage karna chahenge aur apne code ko share karna aasaan banana chahenge. Aage, hum aapko **Cargo** tool se introduce karenge, jo aapko real-world Rust programs likhne mein madad karega.
