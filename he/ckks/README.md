# About The Demo

The demo is an example about [CKKS](https://eprint.iacr.org/2016/421.pdf) Scheme, using the [lattigo](https://github.com/tuneinsight/lattigo) library.

## Getting Started

Install dependencies:

```sh
go mod tidy && go mod vendor
```

Run the demos located in the `cmd/` directory:

```sh
go run cmd/addition/main.go
```

## Parameters Note

- `N`: the ring dimension (security increases with `N` and performance decreases with `N`)
- `Q`: the ciphertext modulus, product of a chain of small coprime moduli `qi` that verify `qi = 1 mod 2N` in order to enable both the RNS and NTT representation
- `P`: the auxiliary modulus used in key switching, the `LogP` default value is `[61]*max(1, floor(sqrt(#qi)))`
- `Xe`: the distribution of the error
- `Xs`: the distribution of the secret
- `scale`: the plaintext scale
- `NthRoot`: the `N`-th root of unity in the polynomial ring

Parameters Selection (FYI):

- Define the security level required, for example `128`, refer to the [HE Standard](https://homomorphicencryption.org/standard/) to get recommended `N` and `logQ` (the actual `logPQ` used should be equal to or smaller than the recommended `logQ` to offer the security defined)
- Fine-tuning `Q` and `P` (depending on the application):
  - the number of `qi` should be `1 + depth` (1 plus the depth of application circuits), and `qi` are chosen to be of size 30 to 60 bits for the best performance
  - `q0` must be large enough to store the result of the computation
  - `q1, ...` should be usually as close as possible to `2^{LogDefaultScale}`
  - the number of `pj` can be `sqrt(#qi)` (Selecting the number involves a tradeoff between homomorphic capacity and the complexity of the key-switching)
  - `|pj| >= max(|qi|)`
- The `LogDefaultScale` is usually the same as `q1, ...`
