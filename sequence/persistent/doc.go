/*
Package persistent defines fully persistent biological sequences data
structures.

If you are reading this you either know what persistence is or you
are in for a bit of fun. Persistence is a functional concept stating that,
rather than changing internal state, a data structure preserves its former
state such that it is always accessible. In a concurrent language like Go it
makes a lot of sense to preserve state rather than overwrite it. You are free to
pass a persistent structure around without fear that its state will change in
the future because it cannot change: it can only be deleted after the last
variable binding it goes out of scope.
*/
package persistent
