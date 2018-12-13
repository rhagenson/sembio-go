/*
Package immutable is an implementation of biological sequences that
have immutability. Each operation on an immutable object results in
a new object with the original remaining unchanged. In order to improve memory utilization,
any unmodified parts of the original object are reused. Intermediate version are safe to
take copies of.
*/
package immutable
