# Rust FFI for C ~ tutorial

Easily Learn how to integrate C to Rust from lvl 0.

## Getting Started

### Project Structure

```bash
.
|-- Cargo.lock
|-- Cargo.toml
|-- build.rs
|-- c_src
|   |-- CMakeLists.txt
|   `-- lib.c
`-- src
    `-- main.rs

2 directories, 6 files
```

### Flow

Sequentially Progress Through :-

> [!NOTE]
> Of course, once you get a hang of it, you can progress in any manner.

1. Create or Gather required C files
    - *lib.c*
2. Setup linking C with Rust
    - *CMakeLists.txt*
    - *cargo.toml*
    - *build.rs*
3. Now freely work on your Rust project
    - *main.rs*

## Coding Part

### Create or Gather required C files

Here, for the sake of simplicity, we've created a function to add only 2 numbers.

```c
#include <stdio.h>

float addition(float a, float b) {
    
    printf("C: I can add %f and %f\n", a,b);
    printf("C: here's the answer!\n");
    
    return a + b + 5; // +5 for handwritting
}
```

### Setup linking C with Rust

#### `CMakeLists.txt`

Unlike Python and Rust where you can easily call one file into another, C requires a *linker* to do that. `CMake` is that linker.

> [!NOTE]
> Here's a little brief about CMake:
> CMake is obviously more than a linker, it's a build system generator.
> You write a `CMakeLists.txt` file that describes your project and how to build it. This includes things like what source files to compile, what libraries to link against, and what the build targets are.
> You run CMake, which reads the CMakeLists.txt file and generates a build system in the format you choose.
> You use the generated build system to compile and link your project via seperate tools or your IDE.

There will be 4 lines code:
    - 2 integral for every CMake file.
    - 2 to setup linking.

##### Integral Lines

###### `cmake_minimum_required()`

To set minimum version for avoiding compatability issues.

###### `project()`

For describing project details including name, languaged used, version & etc.
We're only gonna focus on essentails.

```cmake
project(NAME LANGUAGE_USED) 
```

- `NAME`: Only relevent in the context of CMake

##### Setup Linking

###### `add_library()`

For Creating a library in C.

```cmake
add_library(NAME LIBRARY_TYPE SOURCE)
```

- `NAME`: Giving a name to library.
- `LIBRARAY_TYPE`: If you're not familiar, there can be 3 types of library
    1. **STATIC**: Integrated at compile-time to the executable. We're going to be using this.
    2. **SHARED**: Loaded at runtime to the executable. Multiple programs can share the same library file which can save memory and disk space, however you will have to ensure it's available on every system where the program runs.
    3. **MODULE**: Similiar to `SHARED` but loading using `dlopen/LoadLibrary`, and are not linked to other targets. Don't worry about these types though, it was only for information sake. At STATIC will generally be used.
- `SOURCE`: Path to file(s) that make up the library.
    1. If multiple files, list them all.
    2. However, that can be hectic, so in such cases we can extract path of all files in a variable `file(GLOB LIB_SOURCES "lib/*.c")` then use that variable `${LIB_SOURCES}`.

> [!IMPORTANT]
> Do remember one thing-- nowhere in your code will you using the name of the `C` file(s) to call upon its functions. You will only be calling the library throughout the program. It will automatically find the function from the program and fetch you the required function.

###### `install()`

For installing the library to the system so your program can use it.

> [!NOTE]
> Library is only installed when you build CMake file using `cmake --build . --target install`.
> Since we will be using Rust for building, therefore it will not actually install this library to the system.

```cmake
install(TARGETS myTarget DESTINATION bin)
```

- **myTarget**: Name of library you're installing. It could also be an executable or a script.
- **DESTINATION**: Path for installing.
    1. We can use `.` as we're not actually installing the library.
    2. `.` refers to *root of installation prefix*.
    3. That term basically refers to the default directory set by system for installing stuff. It's different for every OS type. For unix-based system, it will be `/usr/local`, & for windows-- `C:/Program Files/${PROJECT_NAME}`.

Additionally, `install` can be expanded to be:

```cmake
install(TARGETS myTarget
        RUNTIME DESTINATION bin
        LIBRARY DESTINATION lib
        ARCHIVE DESTINATION lib/static)
```

##### Final Output

```cmake
cmake_minimum_required(VERSION 3.0)

project(Tutorial C)

add_library(ask_c STATIC lib.c)

install(TARGETS ask_c DESTINATION .)
```

---

#### `build.rs`


```rust
extern crate cmake;
use cmake::Config;

fn main() {
    
    let dst = Config::new("c_src").build(); //cmake directory

    println!("cargo:rustc-link-search=native={}", dst.display());
    println!("cargo:rustc-link-lib=static=ask_c"); //library to used
}
```

`println!()` in a *build.rs* file doesn't actually print anything to stdout. It's instead used for sending instructions to Cargo during the build process.
For instance, `println!("cargo:rustc-link-lib=static=ask_c")` tells Cargo to link the ask_c library to associated rust project.

### Now freely work on your Rust project 🚀

*main.rs*

```rust
#[link(name="ask_c", kind="static")] //compiles the library to main program

extern "C" { // specifying what all functions we need from the library
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
```

now perform a `cargo run` and run it🍹
