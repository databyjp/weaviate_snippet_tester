## WeaviateWatcher

Test your Weaviate code snippets here

### Global setup

Some code examples will be best suited for running on a local Weaviate instance, and some on a cloud instance with data pre-populated.

#### Local instance

The local instance can be spun up using the `docker-compose.yml` file in the root of this repo.

Then you can access it at `http://localhost:8080` without authentication.

#### Cloud instance (edu-demo)

The cloud instance is available at `https://edu-demo.weaviate.network` and requires authentication. You can use the read-only key (`learn-weaviate`) as shown in the code examples for most cases.

If you require write access, speak to the Education/Documentation team, and we can provide you OIDC access for your Weaviate account.

### Python

#### Setup

Install Python (recommend using Pyenv)
Set up a virtual environment as you prefer - e.g. use `venv` or `conda`
Activate your virtual environment (e.g. `source .venv/bin/activate`)
Install packages in  virtual environment (e.g. `pip install weaviate-client notebook jupyterlab`)

#### Execution

Load the Jupyter notebook, edit and run the code.
(e.g. see `example_py.ipynb`)

### TS

#### Setup

Install Node + NPM
Install required libraries: `npm install typescript ts-node weaviate-ts-client`
Check that `"type": "module"` is in the `package.json` file
Check `tsconfig.json` for `target` and `module` settings ("esnext")

#### Execution

Create your script and run the code.
(e.g. see `example_ts.ts`)

```
node --loader=ts-node/esm <FILE>.ts
```

### Go

#### Setup

Install Go

#### Execution

Run the code (e.g. `example_go.go`)):
```
go run <FILE>.go
```
