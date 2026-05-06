# Review Journal

The review surface for `lattice-mob-forms-vault` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 163, lane `ship`
- `stress`: `sync drift`, score 116, lane `watch`
- `edge`: `local state`, score 145, lane `ship`
- `recovery`: `conflict cost`, score 196, lane `ship`
- `stale`: `form pressure`, score 211, lane `ship`

## Note

The repository should be understandable without pretending it is larger than it is.
