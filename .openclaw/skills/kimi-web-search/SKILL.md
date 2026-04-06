---
name: kimi-web-search
description: Search the web using Kimi API's built-in $web_search function. Use when you need to search for current information, news, or any web content. The skill handles the tool_calls mechanism to perform web searches via Kimi's native web search capability.
---

# Kimi Web Search

Perform web searches using Kimi API's built-in `$web_search` function through the tool_calls mechanism.

## Prerequisites

- Moonshot API Key must be configured in `~/.openclaw/openclaw.json` at `plugins.moonshot.config.webSearch.apiKey`
- The API Key must have web search permissions

## Usage

Run the search script with your query:

```bash
python scripts/kimi_search.py "your search query"
```

Optional arguments:
- `--count N`: Number of results to return (default: 5)

Example:
```bash
python scripts/kimi_search.py "OpenAI GPT-5 latest news" --count 3
```

## How It Works

1. The script calls Kimi API with the `$web_search` tool declaration
2. Kimi model generates tool_calls with search parameters
3. The script returns the arguments back to Kimi (builtin function execution)
4. Kimi performs the search and returns synthesized results with citations

## Output Format

Results include:
- AI-synthesized answer based on search results
- Citations from the sources found

## Notes

- Uses `kimi-k2.5` model by default (recommended for web search)
- Each search costs ¥0.03 + token consumption
- Search results are included in prompt_tokens
