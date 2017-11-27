# BiGr -- **B**io**I**nformatics **Gr**ammar

## Overview

This directory represents a collection of the minimum interactions between biological data. It will be implemented as a set of Go interfaces to abstract away functionality from data layout on-disk. That is to say, there is no reason a simple DNA sequence must be represented as a string, but if a user needs to know what nucleotide is in the 1st or 100th position the DNA sequence must support that lookup. 

## Current Acitivity

Currently, I (Ryan H.) am collecting standard formats to determine the minimum opertions that each biological data type must support (no matter its on-disk sturcture). 

## Major Challenge

The current major challenge I foresee is defining interfaces that in some way state the validity of data structure. That is: How do I define that a DNA sequence cannot contain the letter Z? Or must at least contain the four nucleotides, if not the IUPAC ambiguous codes?

## Projected version 1 state

The version 1 state should define the operations across the base types of biological work, i.e., DNA, RNA, Protein, and others often built off of in other bioinformatics libraries such as Phylip and PDB.

## Projected version 2 additions

The version 2 state should define data/information lossiness and/or reversibility between types. That is to say that converting DNA to RNA has no loss of information because the operation is reversible, however RNA to Protein loses information because the reverse operation is possible, but may not result in the original RNA sequence.
