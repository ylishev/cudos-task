# Cudos-Task Tool Documentation
Short setup and documentation info.

## Build instructions
For easy build, the tool is made available via [cudos-task.dockerfile](cudos-task.dockerfile) and thus has the pre-requirement of docker locally installed.

```bash
git clone git@github.com:ylishev/cudos-task.git 
cd cudos-task

# To build the tool, while you are still in the tool's source directory, you have to run:

docker build -f cudos-task.dockerfile -t cudos-task .

# To start the tool with --help
 
docker run -it --rm --name cudos-task-container cudos-task --help
```

## Linter & Tests

Linter, tests and coverage is available also via [test-cudos-task.dockerfile](test-cudos-task.dockerfile):
```bash
docker build -f test-cudos-task.dockerfile -t testing-cudos-task .
docker run -it --rm --name testing-cudos-task-container testing-cudos-task
```
#### WARN: Running linting and tests above on Mac with ARM might be very slow, and it is recommended to build the images with native architecture, for example:

```bash
docker buildx build --platform linux/arm64 -f test-cudos-task.dockerfile -t testing-cudos-task .
docker run -it --rm --platform linux/arm64 --name testing-cudos-task-container testing-cudos-task
```

Running only the linter (again, recommended with the proper --platform option) might be done locally, in the tool's source directory:
```bash
docker run --rm --platform linux/arm64 -v $(pwd):/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run ./...
```

## Tool usage
```bash
docker run -it --rm -v /Users/yuliyan/developers/workspace/cudos-node/cudos-data:/cudos-data \
--name cudos-task-container cudos-task withdraw-rewards --node https://rpc.testnet.cudos.org:443 \
--keyring-dir /cudos-data --chain-id cudos-testnet-public-4 \
--from test --to-address cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq \
--interval 30s
```

#### Important
<pre>
This tool relies on accounts created (or imported) with <b>cudos-node</b> and have being tested
with <b>--keyring-backend test</b>! The tool needs access to the keys, stored in the keyring backend.

For that reason the path to --keyring-dir <b>must</b> be provided and mounted via docker's <b>-v</b> option.
In my case, keyring-dir is located inside cudos-node tool at:
<i>/Users/yuliyan/developers/workspace/cudos-node/cudos-data</i> and it is exposed by <b>docker</b> as /cudos-data.
That is the reason why <b>--keyring-dir</b> is pointing at <b>/cudos-data</b> location.
</pre>

Aside from the note above, **--from**, **--to-address** and **--interval** flags could be tweaked.
**--interval** (duration syntax, i.e. 30s, 2m, etc.) is used to run the schedule.