#[link(name="ask_c", kind="static")]

extern "C" {
    fn addition(a: f32, b: f32) -> f32;
}

fn main() {
    
    //clearning the whiteboard
    let _ = std::process::Command::new("clear").status();

    //thinking of some numbers
    let a = 5f32;
    let b = 3f32;
    
    //telling C to add these numbers
    println!("Add these numbers: {a} & {b}");
    let result = unsafe {
        addition(a, b)
    };

    //le C:
    println!("C: {} !!!", result);
    
    //outro
    println!("\nNarrator: And the actual answer was {}.", a+b);
    println!("\nC☕");
}