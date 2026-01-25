Got it — here is a **compressed, no-fluff mental model** of everything we discussed.
This is the **version you keep in your head**.

---

## 🔹 Go in one sentence

> **Go lets you write simple code by default, and gives you precise control when you need it.**

---

## 🔹 Data handling (core idea)

* **Nothing is magically avoided** (allocations, copies, GC still exist)
* Go makes **costs explicit and optional**
* You choose:

  * load everything → easy
  * stream → controlled, efficient

### Rule

> **Stream at the edges, structure in the core**

---

## 🔹 Streaming (`io.Reader` / `io.Writer`)

### Use streaming when:

* data is **large**
* size is **unknown**
* data comes from **network / disk**
* processing is **sequential**

Examples:

* HTTP bodies
* file copy
* logs / ETL
* compression
* large JSON/CSV

### Streaming is overkill when:

* business logic
* small configs
* in-memory data
* domain models

---

## 🔹 `bufio`

* Used for:

  * line-by-line reading
  * performance-sensitive I/O
* Not everywhere
* Often used at **I/O boundaries only**

---

## 🔹 Structs & data design

> **Good data structure design reduces code, not abstractions**

* Structs = ownership
* Slices/maps = working memory
* Pointers = shared or mutable state
* Copies are fine → GC handles them
* Avoid premature pointer obsession

---

## 🔹 Garbage Collection

* Yes, copies are GC’d
* Streaming reduces GC pressure
* GC is cheap when objects are short-lived
* You control pressure by design

---

## 🔹 Receivers (methods)

* Not constructors
* Attach behavior to data
* Can mutate (pointer receiver) or not (value)
* Receiver methods may or may not return values

---

## 🔹 Public vs Private (very important)

* **Capitalized → exported (public)**
* **lowercase → internal**
* No keywords
* This is **API design**, not style

> Lowercase = internal utils of the package

---

## 🔹 Embedding

### Embedded structs:

* Used **sometimes**
* For behavior reuse
* Common with stdlib types

### Embedded types:

* Rare
* Only for semantic meaning or method reuse

> **Embed sparingly. Never simulate inheritance.**

---

## 🔹 Function literals (anonymous functions)

* Used **very often**
* Small, local behavior only
* Common in:

  * HTTP handlers
  * goroutines
  * middleware
  * tests

They **enable DI**, but are not DI themselves.

---

## 🔹 Closures (important warning)

A closure = function literal + captured variables

### Good for:

* wiring
* configuration
* handlers
* goroutines (carefully)

### Dangerous when:

* capturing loop variables
* capturing mutable shared state
* hiding large logic
* long-lived accidental references

> **Closures are for wiring, not business logic**

---

## 🔹 Concurrency primitives

* goroutines → cheap threads
* channels → coordination, not data pipes
* mutex → shared state protection
* You don’t need them everywhere
* Projects should *justify* concurrency

---

## 🔹 Other languages comparison

* **C** → full control, full pain
* **C++** → power, complexity
* **Rust** → safety, steep learning
* **Go** → leverage, clarity, control when needed

> Go is the **pragmatic systems language**

---

## 🔹 Recruiter signal (real)

Good Go code signals:

* system thinking
* API discipline
* performance awareness
* testable design

Bad Go code signals:

* Java-style inheritance
* over-abstraction
* misuse of closures
* premature optimization

---

## 🔹 Final rules to remember

1. **Stream when data is big or external**
2. **Structs for logic, streams for movement**
3. **Anonymous for local, named for domain**
4. **Embed rarely, compose always**
5. **Control is optional — but available**

---

### Final sentence (keep this)

> **Go lets you start naïve and end precise.**

If you want, next we can:

* map this into a **project checklist**
* refactor your Go code into idiomatic form
* design a **real Go backend architecture**

You’re no longer learning syntax — you’re learning *judgment*.
