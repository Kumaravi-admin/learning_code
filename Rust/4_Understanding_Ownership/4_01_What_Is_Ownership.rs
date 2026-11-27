/*

4.1 Ownership Kya Hai?
Ownership un rules ka set hai jo batate hain ki ek Rust program memory ko kaise manage karta hai.
Har program ko chalte waqt computer ki memory manage karni hoti hai.
Kuch languages mein garbage collection hoti hai jo program chalte waqt unused memory dhoondh kar saaf karti hai; doosri languages mein programmer ko khud explicitly memory allocate aur free karni padti hai.
Rust teesra tareeka apnata hai: Memory ko ownership ke system aur un rules ke zariye manage kiya jaata hai jinhe compiler check karta hai. Agar koi bhi rule tootta hai, toh program compile nahi hoga. Ownership ke kisi bhi feature se aapka program chalte waqt slow nahi hoga.
Kyunki ownership kaafi programmers ke liye naya concept hai, iski aadat padne mein thoda waqt lagta hai. Achhi baat yeh hai ki jaise-jaise aap Rust aur ownership ke rules mein experienced hote jaate hain, aapko safe aur efficient code likhna utna hi aasaan lagta hai. Lage rahiye!
Jab aap ownership samajh lenge, toh aapko un features ko samajhne ki mazboot foundation mil jaayegi jo Rust ko unique banate hain. Is chapter mein aap ownership ko aise examples ke zariye seekhenge jo ek bahut common data structure—strings—par focus karte hain.

*/
