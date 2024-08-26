# Clickhouse SQL Parser

This is a Go [Clickhouse](https://clickhouse.com) SQL Parser built with [Antlr4](https://www.antlr.org).

The original antlr4 grammar files are from the [official clickhouse repository](https://github.com/ClickHouse/ClickHouse/tree/master/utils/antlr), although the maintainers of clickhouse have stated that this is unsupported, but this is still the easiest way to quickly build a clickhouse sql parser with as robust as possible.

The original antlr4 grammar files contain a little bit of C++ code, but this can be replaced with Go easily, you can see the current grammar file [here](./grammar/).

## Installation

Before using this library, you should know and agree with the following points:

1. **This library has not undergone any security testing, and you need to take full responsibility for any potential risks, including data loss.**
2. **We are currently in the early stages of development and may change some APIs at any time without prior notice.**

Now, you can install this library with the following command.

```shell
go get github.com/hungtcs/clickhouse-sql-parser@master
```

## Roadmap

- [ ] Cover Antlr4 CST to AST
- [ ] Implement AST Visitor
- [ ] AST Serialize to SQL
