# About The Demo

## Basic Requirements

Install the tools below:

- [Poetry](https://python-poetry.org/docs/#installation) for dependency management
- [pyenv](https://github.com/pyenv/pyenv) for python version management

## Getting Started

Install Python `3.11.8` (used for the demo; 3.9+ should work fine):

```sh
pyenv install 3.11.8
pyenv local 3.11.8
```

Spawns a shell within the virtual environment:

```sh
poetry shell
```

Install the projects dependencies:

```sh
poetry install
```

Run the module (in 3 terminals respectively):

```sh
# Party 0
poetry run python -m bgw.elderly -M3 -I0 --no-log
```

```sh
# Party 1
poetry run python -m bgw.elderly -M3 -I1 --no-log
```

```sh
# Party 2
poetry run python -m bgw.elderly -M3 -I2 --no-log
```
