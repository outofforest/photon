# photon

Encoding data structures to byte slice in go using C union-like concept.

This is not enforced by the library, due to performance, but you should never ever try to use it with structures
containing pointers.
