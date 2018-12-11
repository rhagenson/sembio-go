// Copyright 2017 Ryan Hagenson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package bio is a library representing Bioinformatics solutions.

This library defines biological data types abstractly by their functionality (interfaces) and
how these abstract types relate to one another (functions) while saying nothing about how the
concrete data is structured. The advantage of this approach is that concrete data can be
restructured for efficiency while still providing a compile-time check that it is still
a proper functional representation of its abstract type.
*/
package bio
