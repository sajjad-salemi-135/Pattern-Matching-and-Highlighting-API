# Gin-Gonic Pattern Matching and Highlighting API

## Overview

This repository provides a simple web server using the Gin-Gonic framework. The server exposes a single POST endpoint that accepts a text input and a regex pattern, processes the input to find and highlight matching patterns, and returns the results.

## Features

- **Pattern Matching**: Uses regular expressions to find words in the input text that match the provided pattern.
- **Concurrency**: Utilizes goroutines and sync mechanisms to process input efficiently.
- **Highlighting**: Highlights the matched patterns in the input text by enclosing them in HTML `<b>` tags.

## Prerequisites

- Go 1.16 or higher
- Gin-Gonic framework
