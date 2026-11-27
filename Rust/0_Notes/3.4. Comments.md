# We will cover in this lesson

1. Comments kya hote hain
2. Single line comment (`//`)
3. Multi-line comment
4. Code ke end mein comment
5. Code ke upar comment (zyada common style)
6. Documentation comments

---

## 1. Comments kya hote hain

Sabhi programmers apne code ko samajhne mein aasan banane ki koshish karte hain, lekin kabhi-kabhi extra explanation ki zaroorat padti hai. Aise cases mein, programmers apne source code mein _comments_ chhod dete hain jinhe compiler ignore kar deta hai, lekin jo log source code padh rahe hote hain unhe yeh upyogi (useful) lag sakte hain.

---

## 2. Single line comment (`//`)

Yahan ek simple comment hai:

```rust
// hello, world
```

Rust mein, idiomatic comment style do slashes (`//`) se shuru hota hai, aur yeh comment line ke ant (end) tak chalta hai.

---

## 3. Multi-line comment

Aise comments ke liye jo ek single line se zyada bade hote hain, aapko har line par `//` shamil karna hoga, kuch is tarah:

```rust
// So we're doing something complicated here, long enough that we need
// multiple lines of comments to do it! Whew! Hopefully, this comment will
// explain what's going on.
```

---

## 4. Code ke end mein comment

Comments ko code wali lines ke ant mein bhi rakha ja sakta hai:

**Filename: `src/main.rs`**

```rust
fn main() {
    let lucky_number = 8; // I'm feeling lucky today
}
```

---

## 5. Code ke upar comment (zyada common style)

Lekin aap unhe zyada-tar is format mein istemal hote hue dekhenge, jismein comment us code ke upar ek alag line par hota hai jise woh annotate kar raha hai:

**Filename: `src/main.rs`**

```rust
fn main() {
    // I'm feeling lucky today
    let lucky_number = 7;
}
```

---

## 6. Documentation comments

Rust mein ek aur tarah ka comment bhi hota hai, jise _documentation comments_ kehte hain, jis par hum Chapter 14 ke [“Publishing a Crate to Crates.io”](https://doc.rust-lang.org/book/ch14-02-publishing-to-crates-io.html) section mein charcha karenge.
