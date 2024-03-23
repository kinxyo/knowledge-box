use std::io::Write;

extern crate cmake;
use cmake::Config;

fn main() {
    let dst = Config::new("c_src").build(); //cmake directory

    let data = format!("{}\n", dst.display());
    let mut file = std::fs::OpenOptions::new().create(true).append(true).open("LOG").unwrap();
    writeln!(file, "{}", data).unwrap();


    println!("cargo:rustc-link-search=native={}", dst.display());
    println!("cargo:rustc-link-lib=static=ask_c"); //library name given in cmake file
}