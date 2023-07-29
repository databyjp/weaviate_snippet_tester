## WeaviateWatcher

Test your Weaviate code snippets here

### Python

#### Setup

Set up a virtual environment as you prefer - e.g. use `venv` or `conda`
Activate your virtual environment (e.g. `source .venv/bin/activate`)
Install packages in  virtual environment (e.g. `pip install weaviate-client notebook jupyterlab`)

#### Execution

Load the Jupyter notebook, edit and run the code.

### TS

#### Setup

`npm install typescript ts-node weaviate-ts-client`
Check that `"type": "module"` is in the `package.json` file
Check `tsconfig.json` for `target` and `module` settings ("esnext")

#### Execution

Create your script and run the code.
(e.g. see `test_ts_snippet.ts`)

```
node --loader=ts-node/esm <FILE>.ts
```
