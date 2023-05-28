# tbaid (Tech Blog Aid)

`tbaid` is a command line tool which supports writing tech blogs.

## Installation

```sh
git clone git@github.com:ktoyod/tbaid.git
cd /path/to/tbaid
make go/install/tbaid
```

You can use `tbaid` command with

```sh
./bin/tbaid ...
```

## Usage

Generate some candidates for the title of your tech blog.

```sh
OPENAI_API_KEY=sk-xxxxx ./bin/tbaid gentitle --keywords Go,Cobra
# Output:
# 1.「Goで爆速CLI開発！Cobraフレームワークの使い方」
# 2.「CobraでGoのCLI開発をシンプルに！効率化の方法とコツ」
# 3.「Go初心者でも簡単！Cobraを使ったCLIアプリケーション開発入門」

OPENAI_API_KEY=sk-xxxxx ./bin/tbaid gentitle --keywords Go,Cobra --number 5 --language 英語
# Output:
# 1. "Building Command-Line Interfaces with Go and Cobra"
# 2. "Mastering CLI Development with Go and Cobra"
# 3. "Simplifying CLI Development with Go and Cobra"
# 4. "Leveling Up Your Command-Line Tooling with Go and Cobra"
# 5. "Creating Powerful CLI Applications Using Go and Cobra"
```

## Environment variables

You should set the following environment variables to use `tbaid` command.

| env variable | required | description |
|--|--|--|
| OPENAI_API_KEY | ✔ | The API Key for OpenAI platform. |
