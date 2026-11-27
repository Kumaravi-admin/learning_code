/*
Comments

• Comments code ko samajhne mein madad karte hain. Compiler unhe ignore kar deta hai.
• Rust mein idiomatic comment // se shuru hota hai aur line ke end tak chalta hai.
• Ek line se bada comment ho toh har line par // lagao.
• Comment code line ke end mein bhi ho sakta hai, par zyada tar code ke upar alag line par likhte hain.
*/

/*
Is lesson mein hum ye cover karenge:

1. Simple comment (//)
2. Multi-line comment
3. Comment code line ke end mein
4. Comment code ke upar alag line par (common style)
*/

// Dono examples ke paas pehle apna-apna fn main() tha. Ek file mein main sirf ek hi ho sakta hai,
// isliye har example ko alag function bana kar main se call kiya hai.

fn main() {
    // 1. Simple comment
    // hello, world

    // 2. Multi-line comment: har line par alag //.
    // So we're doing something complicated here, long enough that we need
    // multiple lines of comments to do it! Whew! Hopefully, this comment will
    // explain what's going on.

    end_of_line_comment();
    above_line_comment();
}

// 3. Comment line ke end mein
fn end_of_line_comment() {
    let lucky_number = 8; // I'm feeling lucky today
}

// 4. Comment code ke upar (zyada common)
fn above_line_comment() {
    // I'm feeling lucky today
    let lucky_number = 7;
}

/*  Terminal output:

(kuch print nahi hota, comments aur variables ke alawa kuch chalta hi nahi)

*/