# Counting Bits

> NeetCode 150 · Bit Manipulation · Blind 75 — popcounts for 0..n in one pass: i's count is i>>1's count plus i's own lowest bit. Go, stdlib only, tested.

## Problem

Produce the set-bit count of every integer from 0 to n — in O(n) total, not O(n log n) of per-number counting.

## Approach

The DP rides the binary representation: shifting i right by one drops only its lowest bit, so `count(i) = count(i >> 1) + (i & 1)` — and i>>1 < i is always already computed in an ascending fill. One addition per number, answers memoized in the output array itself. The sibling recurrence `count(i) = count(i & (i-1)) + 1` (Kernighan's step from dsa-265 turned into a subproblem link) works identically. The stdlib cross-check over 4096 values in the tests keeps the claim honest.

**Complexity:** O(n) time, O(n) output.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
