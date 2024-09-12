# About The Demo

The demo is an example about using [FedAvg](https://arxiv.org/abs/1602.05629) algorithm to train a CNN image classifiers model with 10 clients/organizations, each client/organization has their own data. The code is adjusted from the flower official [tutorial](https://flower.ai/docs/framework/main/en/tutorial-series-get-started-with-flower-pytorch.html).

## Basic Requirements

Install the tools below:

- [Poetry](https://python-poetry.org/docs/#installation) for dependency management
- [pyenv](https://github.com/pyenv/pyenv) for python version management

## Getting Started

Install Python `3.10.11` (used for the demo):

```sh
pyenv install 3.10.11
pyenv local 3.10.11
```

Spawns a shell within the virtual environment:

```sh
poetry shell
```

Install the projects dependencies:

```sh
poetry install
```

Run the module:

```sh
poetry run python -m fedavg
```
