// Comments

// Sabhi programmers apne code ko samajhne mein aasan banane ki koshish karte hain, 
// lekin kabhi-kabhi extra explanation ki zaroorat padti hai. 
// Aise cases mein, programmers apne source code mein comments chhod dete hain jinhe compiler ignore kar deta hai, 
// lekin jo log source code padh rahe hote hain unhe yeh upyogi (useful) lag sakte hain.

// 1. Yahan ek simple comment hai:

// hello, world



// Rust mein, idiomatic comment style do slashes (//) se shuru hota hai, aur yeh comment line ke ant (end) tak chalta hai. 
// Aise comments ke liye jo ek single line se zyada bade hote hain, aapko har line par // shamil karna hoga, kuch is tarah:

// So we're doing something complicated here, long enough that we need
// multiple lines of comments to do it! Whew! Hopefully, this comment will
// explain what's going on.


// 2. Comments ko code wali lines ke ant mein bhi rakha ja sakta hai:

fn main() {
    let lucky_number = 8; // I'm feeling lucky today
}



// 3. Lekin aap unhe zyada-tar is format mein istemal hote hue dekhenge, 
// jismein comment us code ke upar ek alag line par hota hai jise woh annotate kar raha hai:

fn main() {
    // I'm feeling lucky today
    let lucky_number = 7;
}