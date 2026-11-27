/*
3.6 Repetition with Loops

Code ke kisi block ko ek se zyada baar execute karna aksar upyogi hota hai. 
Is kaam ke liye, Rust kai loops pradan karta hai, jo loop body ke andar ke code ko ant tak
chalayenge aur phir turant shuruat se wapas start karenge.
Loops ke saath experiment karne ke liye, chaliye loops naam ka ek naya project banate hain.

Rust mein teen tarah ke loops hote hain: loop, while, aur for. Chaliye har ek ko try karte hain.
*/

/*
Is lesson mein hum ye cover karenge:

1. loop se code repeat karna
2. Loop se value return karna (break ke saath)
3. Loop labels (nested loops mein bahar wala loop todna)
4. while loop (condition wala loop)
5. for loop (collection par loop)
*/

fn main() {

    //----------------------------------------------------------------------------------------------------
    // 1. Repeating Code with loop
    //----------------------------------------------------------------------------------------------------
    // loop ka matlab: code ko baar-baar chalao, jab tak hum khud break na likhein.
    // Ye code uncomment karoge toh "again!" hamesha print hota rahega.
    // Rokne ke liye terminal mein Ctrl + C dabao.
    //
    // loop {
    //     println!("again!");
    // }

    //----------------------------------------------------------------------------------------------------
    // 2. Returning Values from Loops
    //----------------------------------------------------------------------------------------------------
    // loop ek expression hai, isliye break ke saath value dekar usse bahar nikal sakte hain.
    // Wo value let wale variable mein store ho jati hai.
    let mut counter = 0; // mut isliye, kyunki counter badalta rahega

    let result = loop {
        counter += 1; // har baar 1 badhao

        if counter == 10 {
            // loop band karo aur counter * 2 (yaani 20) bahar bhejo
            break counter * 2;
        }
    }; // semicolon yahan zaroori hai, kyunki ye let statement ka end hai

    println!("The result is {result}"); // output: The result is 20

    // Note: break sirf current loop se nikalta hai. return poore function se nikal jata hai.

    //----------------------------------------------------------------------------------------------------
    // 3. Disambiguating with Loop Labels
    //----------------------------------------------------------------------------------------------------
    // Nested loops mein break aur continue sirf sabse andar wale loop par lagte hain.
    // Bahar wale loop ko todna ho toh uska label lagao. Label single quote (') se shuru hota hai.
    let mut count = 0;

    'counting_up: loop {
        // ye bahar wala loop hai, label: 'counting_up
        println!("cont = {count}");
        let mut remaining = 10;

        loop {
            // ye andar wala loop hai, bina label ke
            println!("remaining = {remaining}");

            if remaining == 9 {
                break; // sirf andar wala loop band hoga
            }
            if count == 2 {
                break 'counting_up; // bahar wala (labeled) loop band hoga
            }
            remaining -= 1;
        }

        count += 1; // andar wala loop khatam hone ke baad count badhao
    }

    println!("End count = {count}");

    /*
    Bahar wala loop 'counting_up 0 se 2 tak count up karta hai.
    Andar wala loop 10 se 9 tak count down karta hai.
    Pehla break (bina label) sirf andar wale loop se nikalta hai.
    break 'counting_up bahar wale loop se nikalta hai.

    Book mein diya gaya output:

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
    */

    //----------------------------------------------------------------------------------------------------
    // 4. Streamlining Conditional Loops with while
    //----------------------------------------------------------------------------------------------------
    // while: jab tak condition true hai, loop chalta rahega.
    // Condition false hote hi loop apne aap band ho jata hai (alag se break nahi likhna padta).
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1; // number ko kam karna zaroori hai, warna loop kabhi khatam nahi hoga
    }

    println!("LIFTOFF!!!");

    //----------------------------------------------------------------------------------------------------
    // 5. Looping Through a Collection with for
    //----------------------------------------------------------------------------------------------------
    // Pehle while se array par loop (galti hone ka risk zyada):
    let a = [10, 20, 30, 40, 50];
    let mut index = 0;

    while index < 5 {
        // index ko khud sambhalna padta hai
        println!("the vlaue is: {}", a[index]);
        index += 1;
    }
    // Problem 1: array ki length badli aur 5 badalna bhool gaye, toh program panic karega.
    // Problem 2: slow bhi hai, kyunki har baar Rust check karta hai ki index array ke andar hai ya nahi.

    // Behtar tarika: for loop. Index khud nahi sambhalna padta, aur bounds ki galti ho hi nahi sakti.
    let b = [10, 20, 30, 40, 50];

    for element in b {
        println!("output of for loop -> the value is: {element}");
    }

    // for ke saath Range aur rev():
    // (1..4) matlab 1, 2, 3 (4 shaamil NAHI hota).
    // rev() use ulta kar deta hai,