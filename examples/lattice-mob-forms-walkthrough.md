# Lattice Mob Forms Vault Walkthrough

This note is the quickest way to read the extra review model in `lattice-mob-forms-vault`.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 163 | ship |
| stress | sync drift | 116 | watch |
| edge | local state | 145 | ship |
| recovery | conflict cost | 196 | ship |
| stale | form pressure | 211 | ship |

Start with `stale` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

`stale` is the optimistic case; use it to make sure the scoring path still rewards strong signal.
