# About The Demo

The demo is an example about [Spiral PIR Scheme](https://eprint.iacr.org/2022/368), using the [Blyss SDK](https://github.com/blyssprivacy/sdk).

## Getting Started

Run the demo:

```sh
cargo run
```

## Explain

The demo proceeds as follow:

1. Server init server state, and update the server state with the key-value pairs below:

   ```json
   {
        "key-1": "value-1",
        "key-2": "value-2",
        "key-3": "value-3",
   }
   ```

2. Client get params from server and generate query to retrieve `key-1` value privately
3. Server process the query and response without knowing the key-value pair retrieved
4. Client handle the response to get the `value-1`

## Parameters Note

- `n`: plaintext dimension
- `nu_1`, `nu_2`: database configuration ($v_1$, $v_2$ in the paper)
  - `nu_1`: the first dimension, determines the number of rounds of query expansion and the noise growth in the query expansion scales exponentially with $v_1$
  - `nu_2`: the number of subsequent dimension, determines the number of rounds of folding needed during query processing
- `p`: the plaintext modulus, a power of two with maximum value $2^{30}$
- `q1`, `q2`: the reduced moduli
- `t_gsw`, `t_conv`, `t_exp_left`, `t_exp_right`: decomposition dimensions
