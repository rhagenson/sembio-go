/*
Package persistent is an implementation of biological sequences that
has partial persistence. Each operation on a persistent object results in
a new object. In order to improve memory utilization, persistent here does
its best to reuse any unmodified parts of the initial object.
*/
package persistent
