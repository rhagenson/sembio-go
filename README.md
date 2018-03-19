# BiGr -- **B**io**I**nformatics **Gr**ammar

## Overview

This directory represents a collection of the minimum interactions with biological data. It will be implemented as a set of Go interfaces and function to abstract away functionality from data layout on-disk. That is to say, there is no reason a simple DNA sequence must be represented as a string, but if a user needs to know what nucleotide is in the 1st or 100th position the DNA sequence representation must support that lookup as well as `Transcribe(dna DNASequence) RNASequence` function should be possible with the representation.

## Why does `<interface>` not have `<method>`?

The interfaces are defined to be as small as possible with internal functionality being paramount. In a basic example, the `Sequence` interface knows nothing about how to transcribe/translate itself because the necessary translation table is not internal to the `Sequence`, rather there are functions `Transcribe()` and `Translate()` that define this behavior. The result is that to translate a sequence it is not `Sequence.Translate() -> <TranslatedSequence>` but rather `Translate(Sequence) -> <TranslatedSequence>`. This being said, by moving the transcription/translation functionality outside the `Sequence` interface this allows dedicated `Transcriber` and `Translator` interfaces to be designed as well as the `Transcribe()` and `Translate()` functions to be reused for each new data structure that is transcribable and translatable.

It may be that `<interface>` would have `<method>` in any other library. However, if the data structures that should implement `<interface>` would require information beyond their internal state or including the method would promote defining an otherwise unnecessary private variable than there is no good reason to bloat the interface unnecessarily. It is better to define non-internal functionality as a non-internal method that can be shared between structures and improved in one central location.

## Current Activity

Currently, I (Ryan H.) am collecting standard formats to determine the minimum operations that each biological data type must support (no matter its on-disk structure). My reasoning for stating this is that I want a more mathematical, but quick-to-process manner to work with biological data while being equally capable to switch out on-disk representations to optimize processing for different operations (slow DNA Sequence lookup, but fast DNA translation to Protein). By defining how biological data types interact internally (functions) and with the world around them (interfaces) I can switch out how data is structured more easily.

## Major Challenge

The current major challenge I foresee is defining interfaces that in some way state the validity of a data structure. That is: How do I define that a DNA sequence cannot contain the letter Z? Or must at least contain the four nucleotides, if not the IUPAC ambiguous codes? The best course is a `Validate() bool` method which returns internal validity, but this can be falsified by always returning `true` no matter the validity, which I do not want to allow. The next possibility is functions such as `ValidateDNA(d DNA) bool` which would allow me to check the validity of a type by asking for the proper attributes of its internal state. This latter is likely best, but would require enough methods in each interface to validate the data structure.

## Projected version 1 state

The version 1 state should define the operations across the base types of biological work, i.e., DNA, RNA, Protein, and others that are found in other bioinformatics libraries such as Phylip and PDB.

## Projected version 2 additions

The version 2 state should define data/information lossiness and/or reversibility between types. That is to say that converting DNA to RNA has no loss of information because the operation is reversible, however RNA to Protein loses information because the reverse operation is possible, but may not result in the original RNA sequence. If I can abstract out the measurement of this lossiness that would be great, but I doubt it is possible to do that simply.

## What about `biogo`?

`biogo` does somethings really well, however there are somethings I do not like about it that are too fundamental to change without breaking a lot of existing code. Therefore, I am making a second Bioinformatics Go library that matches what I want. 
