// Package geo contains a collection of 2-D types and other mathematical functions
// geared towards games.
//
// NumGen
//  The functions that deal with NumGen/VecGen are useful for when you want to generate
//  random or specific series of numbers with some constraints. An example might be a
//  particle system where, instead of taking a Vec for the initial particle velocity,
//  it takes a VecGen supplied by the user. The particle system calls this VecGen whenever
//  a new particle is created giving the user total control over behavior of the initial
//  velocities.
//
//
// Vec
//  2-D Vector with many useful functions for working with it by value or reference.
//
// Rect
//  Defines a rectangular area. Like Vec it also has many functions for values and references.
//
// Miscellaneous
//  ...
//
// TODO:
//  - circle
//  - camera
//  - noise
package geo
