# ⚡ My own GO-based library 

[![Go Version](https://img.shields.io/github/go-mod/go-version/iamkaro/GO-LIB)](https://go.dev/) 
[![GitHub Stars](https://img.shields.io/github/stars/iamkaro/GO-LIB?style=social)](https://github.com/iamkaro/GO-LIB/stargazers)


*Developed in 2021* 
<br>


GO-LIB is a Modern Go Developer's Utility Kit. 
A high-performance toolkit that extends the standard library with essential features. 
GO-LIB provides modules for **advanced concurrency**, **robust networking**, **comprehensive OS tools**, data manipulation using **generics**, and **native localization support**.
<br>



## ✨ Highlights & Key Features

* **Go Idiomatic Concurrency:** Specialized packages like `sync/once` and `sync/value` offer safer, cleaner ways to manage Goroutines and shared state.
* **Modern Generics:** Full support for **Go Generics** in data structures (`list`) and mathematical operations (`number`).
* **Localization Support:** Out-of-the-box utilities for managing **Persian/Solar time** (`time/solar`).
* **Comprehensive OS Toolkit:** Advanced file system management (`os/fs`), CLI argument parsing (`os/flag`), and interactive terminal menus (`os/menu`).
* **Minimal Dependencies:** Focused on core Go functionality to keep the library lean and fast.


## 🛠️ Installation

Simply use `go get` to add the module to your project's dependencies:

```bash
go get github.com/iamkaro/GO-LIB 
```


## 📚 Documentation

### 1. Concurrency & Synchronization ⚙️

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`sync/once`** | Provides robust **thread-safe execution** mechanisms for tasks, including singleton initialization (`NewSingle`) and high-performance **Goroutine Thread Pools** (`NewThread`). | `NewSingle()`, `single.Run()`, `single.TryRun()`, `NewThread()`, `thread.Run()`, `thread.RunNow()`, `thread.Start()`, `thread.Empty()`, `thread.Stop()` |
| **`sync/value`** | A generic, atomic wrapper for any value type, ensuring **safe concurrent read/write** operations without explicit locking. | `New()`, `v.Set()`, `v.Get()` |
| **`sync/self`** | Implements a simple, safe signal mechanism for managing the state of a resource that needs to be **closed only once** and checked repeatedly. | `New()`, `it.Closed()`, `it.NotClosed()`, `it.Close()` |

> **How to Import:** To import any package in this section, use the main path `github.com/iamkaro/GO-LIB` followed by the sub-path (e.g., **`sync/once`**). **Example: `"github.com/iamkaro/GO-LIB/sync/once"`**

### 2. Operating System & I/O 📁

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`os/fs`** | A comprehensive file system utility kit offering advanced features like recursive **folder scanning**, state checks (`Exist`, `State`), file movement, and safe I/O handlers. | `IsFile()`, `IsFolder()`, `Exist()`, `NotExist()`, `State()`, `Scan()`, `Glob()`, `Move()`, `CreateFolder()`, `Delete()`, `DeleteFull()`, `FileWriter()`, `FileReader()`, `NewPath()` |
| **`os/term`** | Low-level utilities for direct terminal interaction, including **cursor manipulation** (`Move*`, `Goto*`), screen clearing, style application, and direct character input capture. | `GetChar()`, `GetWord()`, `GetLine()`, `MoveUP()` (and all others), `GotoX()`, `GotoXY()`, `Clear()`, `NewStyle()`, `DisableCtrlC()` |
| **`os/menu`** | Facilitates the creation of simple, interactive **Terminal User Interface (TUI) menus** for command-line applications by easily adding, starting, and printing command options. | `AddCommand()`, `Start()`, `Print()` |
| **`os/flag`** | A **generic command-line argument parser** that simplifies flag definition, parsing, and retrieval (`Get[Type]`), ensuring type safety for CLI inputs. | `Get[Type]()`, `Parse()`, `ParseFrom()`, `Print()`, `Reset()` |

> **How to Import:** To import packages in this section, use the main path `github.com/iamkaro/GO-LIB` followed by the sub-path (e.g., **`os/fs`**). **Example: `"github.com/iamkaro/GO-LIB/os/fs"`**

### 3. Networking & Communication 🌐

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`browser`** | A feature-rich HTTP client wrapper for making GET/POST requests, managing **cookies**, handling file downloads, and configuring maximum concurrent requests. | `New()`, `NewWithoutCookies()`, `Get()`, `Post()`, `Download()`, `MaxRequests()` (Config) |
| **`httpx`** | A streamlined HTTP server toolkit that simplifies backend development by offering easy **routing management** (`AddRoute`, `SetRoutes`) and server lifecycle control. | `NewServer()`, `server.AddRoute()`, `server.SetRoutes()`, `server.RemoveRoute()`, `server.Listen()`, `server.Close()` |
| **`iprange`** | Tools for creating, parsing (`ParseRange`), and managing objects representing **ranges and pools of IP addresses**, highly effective for access control lists. | `New()`, `ParseRange()`, `NewPool()` |
| **`websocket`** | Provides robust primitives for easily setting up a **WebSocket Server** (`NewServer`) and managing client connections (`NewClient`), including send/receive methods. | `NewServer()`, `NewClient()`, `client.Send()`, `client.Receive()`, `client.Close()` |
| **`notification/tg`** | A simplified wrapper for the Telegram Bot API, focusing on easy **sending of messages and files** to channels or users via a configured bot token. | `SetBot()` (Config), `Send()`, `SendByFile()` |

> **How to Import:** Packages in this section are imported directly from the module root `github.com/iamkaro/GO-LIB` (e.g., **`browser`**). **Example: `"github.com/iamkaro/GO-LIB/browser"`**

### 4. Utilities & Data 🔧

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`storage`** | Provides a fixed-size, thread-safe **in-memory byte buffer** for efficient storage, retrieval, and counting of raw binary data. | `NewBytes()`, `storage.IsFull()`, `storage.IsNotFull()`, `storage.Save()`, `storage.SaveByte()`, `storage.Get()`, `storage.GetAndReset()`, `storage.Count()` |
| **`logger`** | A flexible logging utility supporting standard console output, logging to files (`ToFile`), and separation of concerns through named loggers and sub-loggers. | `ToFile()` (Config), `Log()`, `Logf()`, `Panic()`, `Panicf()`, `Error()`, `NotError()`, `NewLogger()`, `SubLogger()` |
| **`check`** | Contains common **validation helpers** and a unique **Fluent Conditional** syntax (`IF().Else()`) for writing concise, readable, and highly expressive validation logic. | `Create()`, `Check()`, `Value()`, `IF().Else()` (Fluent) |
| **`list`** | Provides essential generic array and slice utility methods (`Add`, `Remove`, `Foreach`, `Len`), allowing type-safe manipulation of lists. | `NewArray[T]()`, `Add()`, `Remove()`, `Get()`, `Len()`, `Foreach()` |
| **`number`** | Offers generic and safe mathematical utility functions, including **rounding**, power calculation, and thread-safe finding of minimum and maximum values (`Min[T]`, `Max[T]`). | `Round()`, `Pow()`, `Min[T]()`, `Max[T]()`, `min.Value()`, `max.Value()`, `min.Out()`, `max.Out()` |
| **`text`** | A powerful toolkit for complex string manipulation, including **fixed-length formatting**, generic parsing of strings to any type (`Parse[Type]`), and substring operations. | `FixLength()`, `Parse[Type]()`, `ParseBool()`, `StringOf()`, `ToString()`, `ParseComplex()`, `SubString()` |

> **How to Import:** Packages in this section are also imported directly from the module root `github.com/iamkaro/GO-LIB` (e.g., **`logger`**). **Example: `"github.com/iamkaro/GO-LIB/logger"`**

### 5. Encoding & Serialization 🔐

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`encoding/json`** | Provides custom, streamlined functions for encoding (`Encode`, `BEncode`) and decoding (`Decode`, `BDecode`) data to and from the **JSON format**. | `Encode()`, `Decode()`, `BEncode()`, `BDecode()` |
| **`encoding/base64`** | Simple, direct utility functions for **Base64 encoding** (`Encode`) and decoding (`Decode`) of data. | `Encode()`, `Decode()` |

> **How to Import:** To import packages in this section, use the main path `github.com/iamkaro/GO-LIB` followed by the sub-path (e.g., **`encoding/json`**). **Example: `"github.com/iamkaro/GO-LIB/encoding/json"`**

### 6. Time & Localization 📅

| Package | Purpose | Key Functions/Methods |
| :--- | :--- | :--- |
| **`time`** | Extended time utilities, providing **detailed date/time parsing and formatting across any timezone structure**. The package includes dedicated structures for specific timezones, such as **Iran Standard Time (IST)**, as implementation examples. | `Iran.TimeNow()`, `Iran.Parse()`, `Iran.ParseString()`, `Iran.ParseUnix()`, `Iran.ParseUnixNano()` |
| **`time/solar`** | Specialized functions for converting between **Gregorian and Persian (Solar/Jalali) calendars**, including leap year calculation and Unix timestamp conversion. | `ToSolar()`, `LeapYear()`, `ToUnix()` |

> **How to Import:** To import packages in this section, use the main path `github.com/iamkaro/GO-LIB` followed by the sub-path (e.g., **`time/solar`**). **Example: `"github.com/iamkaro/GO-LIB/time/solar"`**


## 🤝 Contributing to GO-LIB

Your contributions are highly valued! Whether it's reporting a bug, proposing a new feature, or submitting code, every effort helps make GO-LIB better.

1.  **Report Bugs:** Please open an [Issue](https://github.com/iamkaro/GO-LIB/issues) if you encounter any problems.
2.  **Suggest Features:** Discuss new ideas on the [Issues](https://github.com/iamkaro/GO-LIB/issues) page.
3.  **Submit Code:** Fork the repository and submit a [Pull Request](https://github.com/iamkaro/GO-LIB/pulls) with your changes.


