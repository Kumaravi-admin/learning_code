/// Data Types


fn main() {

let guess: u32 = "42".parse().expect("Not a number!");
println!("{}",guess);


// Floating-Point Types
// Rust ke paas floating-point numbers ke liye bhi do primitive types hain, jo decimal points wale numbers hote hain. Rust ke floating-point types f32 aur f64 hain, jo size mein क्रमशः 32 bits aur 64 bits hain. Default type f64 hai kyunki modern CPUs par, yeh f32 ke lagbhag barabar speed par chalta hai lekin zyada precision ke kabil hota hai. Sabhi floating-point types signed hote hain.
// Yahan ek example hai jo floating-point numbers ko action mein dikhata hai:

let x = 2.0;   // f64
let y: f32 = 3.0; // f32
println!("Output of x: {}",x);
println!("Output of y: {}",y);


// Numeric Operations
// Rust sabhi number types ke liye basic mathematical operations ko support karta hai jo aap expect karenge: addition, subtraction, multiplication, division, aur remainder. Integer division zero ki taraf nearest integer par truncate hota hai. Niche diya gaya code yeh dikhata hai ki aap let statement mein har numeric operation ka kaise use karenge:

// addition
let sum = 5 + 10;
println!("Addition: {}",sum);

// subtraction
let differance = 98.5 - 4.3;
println!("Subtraction: {}",differance);


// multiplication
let product = 4 *30;
println!("Multiplication: {}",product);


// division
let quatient = 56.7 / 32.2;
let truncated = -5 / 3; // Results in -1

println!("Division quatient: {}", quatient);
println!("Division truncated: {}", truncated);


// reminder
let remainder = 43 % 5;
println!("Remainder: {}", remainder);


// The Boolean Type
// Jaise ki zyadatar doosri programming languages mein hota hai, Rust mein ek Boolean type ki do possible values hoti hain: true aur false. Booleans size mein ek byte ke hote hain. Rust mein Boolean type ko bool ka use karke specify kiya jata hai. For example:

let t = true;

let f: bool = false;   // with explicity type annotation

println!("Result of t: {}",t);
println!("Result of f: {}",f);


// The character type
// Rust ka char type language ka sabse primitive alphabetic type hai. char values declare karne ke kuch examples yeh hain:

let c = 'z';
let z: char = 'z';
let heart_eyed_cat = '😻';
println!("Result of c: {}",c);
println!("Result of z: {}",z);
println!("Result of heart eyed cat: {}", heart_eyed_cat);

// The Tuple type
// Ek tuple alag-alag types ke kai values ko ek compound type mein group karne ka ek general tareeqa hai. Tuples ki ek fixed length hoti hai: ek baar declare hone ke baad, yeh size mein grow ya shrink nahi ho sakte hain.
// Hum parentheses ke andar comma-separated list likh kar ek tuple banate hain. Tuple mein har position ka ek type hota hai, aur tuple mein alag-alag values ke types ka same hona zaroori nahi hai. Humne is example mein optional type annotations add kiye hain:

let tup = (500,6.4,1);

// Variable tup poore tuple se bind hota hai kyunki ek tuple ko ek single compound element maana jata hai. Tuple se individual values bahar nikalne ke liye, hum pattern matching ka use karke tuple value ko destructure kar sakte hain, jaise ki yeh:

let (x,y,z) = tup;
println!("The value of x is: {x}");
println!("The value of y is: {y}");
println!("The value of z is: {z}");


// Hum ek period (.) ke baad us value ke index ka use karke bhi direct tuple element ko access kar sakte hain jise hum access karna chahte hain. For example:
let x: (i32,f64,u8) = (500,6.4,1);

let five_hundred = x.0;
let six_point_four = x.1;
let one = x.2;

println!("Result of five hundered: {five_hundred}");
println!("Result of five six point four: {six_point_four}");
println!("Result of five one: {one}");


// The Array Type
// Multiple values ka collection rakhne ka ek aur tareeqa array ke sath hai. Tuple ke mukable, array ke har element ka type same hona chahiye. Kuch doosri languages ke arrays ke mukable, Rust mein arrays ki ek fixed length hoti hai.
// Hum square brackets ke andar comma-separated list ke roop mein array ki values likhte hain:

let a = [1,2,3,4,5];

println!("Output of a index 3: {}",a[3]);  // index 3 ko print kar rahe hain

// Arrays tab upyogi hote hain jab aap chahte hain ki aapka data heap ke bajaye stack par allocate ho (hum Chapter 4 mein stack aur heap par aur charcha karenge), bilkul un anya types ki tarah jo humne ab tak dekhe hain, ya jab aap yeh sunishchit karna chahte hain ki aapke paas hamesha elements ki ek fixed sankhya rahe. Halanki, ek array vector type jitna flexible nahi hota. Vector ek similar collection type hai jo standard library dwara pradan kiya jaata hai jise size mein badhne ya ghatne ki anumati hoti hai kyunki iska content heap par rehta hai. Agar aap sure nahi hain ki array use karna hai ya vector, toh zyada sambhavna hai ki aapko vector use karna chahiye. Chapter 8 mein vectors par aur vistar se charcha ki gayi hai.
// Halanki, arrays tab zyada upyogi hote hain jab aap jaante hain ki elements ki sankhya ko badalne ki zaroorat nahi padegi. Udaharan ke liye, agar aap kisi program mein mahino ke naamon ka istemal kar rahe hain, toh aap shayad vector ke bajaye ek array ka istemal karenge kyunki aap jaante hain ki ismein hamesha 12 elements honge:
let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];
              

println!("Output of moths 6th index value: {}", months[6]); // 6th index ki value ko print kar rahe hain
// Aap har element ke type, ek semicolon, aur phir array mein elements ki sankhya ke saath square brackets ka istemal karke ek array ka type likhte hain, kuch is tarah:

let a: [i32;5] = [1, 2, 3, 4, 5];
} // end of main function
