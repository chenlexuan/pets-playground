# About The Demo

The demo is an example of using [OpenMined PSI](https://github.com/OpenMined/PSI) library in Go. Adjusted from the intergration test of the library's Go binding.

## Getting Started

Install bazel 6.4.0: (different versions may experience issues)

```sh
sudo apt install bazel-6.4.0
```

Build the demo:

```sh
bazel-6.4.0 build -c opt //:ecdh_psi
```

Run the demo:

```sh
./bazel-bin/ecdh_psi_/ecdh_psi
```

## Explain

The demo proceeds as follow:

1. Server create the setup message from its private input (100 elements): `[0, 2, 4, 6, ..., 198]`
2. Client create its request from its private input (10 elements): `[0, 1, 2, 3, ..., 9]`
3. Server process the request coming from client and response to the client
4. Client compute the intersection according to the response and setup message created in the step 1
