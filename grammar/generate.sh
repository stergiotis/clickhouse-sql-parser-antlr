#!/usr/bin/env bash

antlr4 -Dlanguage=Go -visitor -no-listener -package grammar -o . *.g4
