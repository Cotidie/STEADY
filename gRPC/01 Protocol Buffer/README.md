# Protocol Buffer
| https://developers.google.com/protocol-buffers

## Table of Contents
- [Overview](#overview)
  - [Setting](#setting)
- [Basics](#basics)
  - [Type](#type)
  - [Tag (Field Number)](#tag-field-number)

## Overview
![Protocol Buffer](https://i.imgur.com/bBXCEpw.png)

Protocol Buffer is a format of data structure for serializing and deserializing suggested by Google. It has several advantages over JSON, XML etc. Protocol Buffer is commonly used in **gRPC**.
- Data is **strongly typed**
- Can have **comments** over schema
- Stored into **binary**, not string

### Setting
* ```vscode-proto3```: vscode extension for syntax highlighting .proto files
* ```Protoc```: .proto file compiler

## Basics
```proto3
// Decide which protocol buffer syntax to use
syntax = "proto3";

import "google/protobuf.timestamp.proto"

// Enum
enum EyeColor {
    EYE_COLOR_UNSPECIFIED = 0;
    EYE_COLOR_GREEN = 1;
    EYE_COLOR_BROWN = 2;
}

// Define a message schema(ref. Class)
message User {
    // <type> <field> = <field number>
    uint32 id = 1;
    string lname = 2;
    repeated string phones = 3;
}
```
- it won't compile if there is a unpopulated field
  
### Type
Default value of Number types is 0. Put ```repeated``` keyword in the front to indicate an array type.
- **(s)int32/64**
- **uint32/64**
- **(s)fixed32/64**: fixed amount of bytes
- **float, double**
- **bool**: true/false. default to 0
- **string**: UTF-8/ASCII encoded string. default to ""
- **bytes**: raw data bytes. default to NULL
- **enum**: user defined type. default to its first field, with field number 0

### Tag (Field Number)
Field numbers are unique numbers that each field has as a binary identifier.
- range 1 to 15 takes one byte to encode
    => Use for frequently occuring data
- range 16 to 2047 takes two bytes.
- range 19000 to 19999 is reserved by compiler
