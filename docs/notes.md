# some notes #


// Circuit Breaker

// Circuit breakers automatically degrade service functions in response to a failures, preventing larger or cascading failures by eliminating recurring errors and providing reasonable error responses.
// Like the electrical switch from which this pattern derives its name, Breaker has two possible states: closed and open.
// In the closed state everything is functioning normally. All requests received from the client by the breaker are forwarded unchanged, and all responses are forwarded back to the client.
// In the open state, the breaker doesn’t forward requests, instead it fails fast by responding with an informative error message.
// The breaker should track the errors returned the underlying function; if the number of consecutive errors exceeds a defined threshold, the breaker trips and its state switches to open.
// Let's implement a circuit breaker that accepts a function, and an unsigned integer representing the number of consecutive failures allowed before the circuit automatically opens.
// This exercise can be completed in any language, below we provide some example signatures in popular languages. Feel free to work in this window (and make any changes you want), or in a text editor/IDE.

// Go
// Example breaker function accepts a function and returns a wrapped function with the same signature
//  func Breaker(circuit func(context.Context) (string, error), failureThreshold uint) func(context.Context) (string, error) {}

// Typescript
// Breaker function accepts a function and threshold, and returns a wrapped function with the same signature
// function breaker(circuit: () -> string, failureThreshold: int): () -> string {}

// Python
// Breaker function accepts a function and threshold, and returns a wrapped function with the same signature
// def breaker(circuit: Callable, failureThreshold: int): Callable

// Java
// Breaker function accepts a function and threshold, and returns a wrapped function with the same signature
// public Callable<String> breaker(Callable<String> circuit, failureThreshold: int) {}