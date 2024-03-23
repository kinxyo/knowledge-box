
cmake_minimum_required(VERSION 3.0)

project(Tutorial C) # project(NAME LANGUAGE) 

#[[ 
    NAME -- Only relevant to CMake
]]

add_library(ask_c STATIC lib.c) # add_library(NAME LIBRARY_TYPE ENTRY)

#[[ 
    add_library() -- creating a library.
    NAME -- giving a name to the library.
    LIBRARY_TYPE -- STATIC | SHARED | MODULE ~ Visit README to learn more.
    SOURCE -- path to all file(s) from which the library is created. If multiple files, list them all.
]]

install(TARGETS ask_c DESTINATION .) # install(TARGETS myTarget DESTINATION bin )

#[[

  myTarget -- name of executable, library, or a script
  DESTINATION -- bin ~ directory relative to the installation prefix.

  additionally, `install` can be expanded to be:
    
    install(TARGETS myTarget
        RUNTIME DESTINATION bin
        LIBRARY DESTINATION lib
        ARCHIVE DESTINATION lib/static)

]]


# -------------------

#[[

1. STATIC: A static library is a collection of object files that can be linked into an application at compile-time. It becomes part of the executable file. The advantage of static libraries is that you can be sure that all your function calls to a library will work because the library code is already part of the executable. However, they can make the executable file size quite large.

2. SHARED: A shared library is a library that is loaded at runtime when the executable file that depends on it is run. The advantage of shared libraries is that multiple programs can share the same library file which can save memory and disk space. However, you need to ensure that the library is available on every system where the program runs.

3. MODULE: Module libraries are plugins that are meant to be loaded at runtime using dlopen/LoadLibrary or similar functionality. Unlike SHARED libraries, they are not linked into other targets.


]]