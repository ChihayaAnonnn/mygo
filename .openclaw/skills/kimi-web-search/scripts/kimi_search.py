#!/usr/bin/env python3
"""
Kimi Web Search - Perform web searches using Kimi API's built-in $web_search function.

Usage:
    python kimi_search.py "your search query" [--count N]

Example:
    python kimi_search.py "OpenAI GPT-5 latest news" --count 3
"""

import argparse
import json
import os
import sys
from typing import Optional

try:
    from openai import OpenAI
except ImportError:
    print("Error: openai package not installed. Run: pip install openai")
    sys.exit(1)


def get_api_key() -> str:
    """Get Moonshot API Key from OpenClaw config."""
    config_path = os.path.expanduser("~/.openclaw/openclaw.json")
    
    if not os.path.exists(config_path):
        raise FileNotFoundError(f"OpenClaw config not found: {config_path}")
    
    with open(config_path, 'r') as f:
        config = json.load(f)
    
    api_key = config.get("plugins", {}).get("moonshot", {}).get("config", {}).get("webSearch", {}).get("apiKey")
    
    if not api_key:
        # Fallback to models config
        api_key = config.get("models", {}).get("providers", {}).get("moonshot", {}).get("apiKey")
    
    if not api_key:
        raise ValueError("Moonshot API Key not found in OpenClaw config")
    
    return api_key


def search_web(query: str, count: int = 5) -> str:
    """
    Perform web search using Kimi's $web_search builtin function.
    
    Args:
        query: Search query string
        count: Number of results to request (informational, actual results depend on Kimi)
    
    Returns:
        Synthesized search results with citations
    """
    api_key = get_api_key()
    
    client = OpenAI(
        api_key=api_key,
        base_url="https://api.moonshot.cn/v1"
    )
    
    # Define the $web_search tool
    tools = [{
        "type": "builtin_function",
        "function": {
            "name": "$web_search"
        }
    }]
    
    messages = [
        {
            "role": "system",
            "content": "You are a helpful assistant with web search capabilities."
        },
        {
            "role": "user",
            "content": f"Please search for: {query}"
        }
    ]
    
    # First call: Get tool_calls from model
    # Note: $web_search requires temperature=1
    response = client.chat.completions.create(
        model="kimi-k2.5",
        messages=messages,
        tools=tools,
        temperature=1
    )
    
    message = response.choices[0].message
    finish_reason = response.choices[0].finish_reason
    
    if finish_reason != "tool_calls":
        # Model didn't call tools, return direct response
        return message.content or "No response from model."
    
    # Add assistant message with tool_calls to conversation
    messages.append({
        "role": "assistant",
        "content": message.content or "",
        "tool_calls": [
            {
                "id": tc.id,
                "type": tc.type,
                "function": {
                    "name": tc.function.name,
                    "arguments": tc.function.arguments
                }
            }
            for tc in message.tool_calls
        ]
    })
    
    # Add tool response for each tool_call
    # For $web_search, we just return the arguments back to Kimi
    # Kimi will execute the search internally
    for tool_call in message.tool_calls:
        messages.append({
            "role": "tool",
            "tool_call_id": tool_call.id,
            "content": tool_call.function.arguments  # Return arguments as-is for builtin function
        })
    
    # Second call: Get final response with search results
    final_response = client.chat.completions.create(
        model="kimi-k2.5",
        messages=messages,
        tools=tools,
        temperature=1
    )
    
    final_message = final_response.choices[0].message
    
    # Extract citations if present in the response
    result = final_message.content or "No results found."
    
    return result


def main():
    parser = argparse.ArgumentParser(
        description="Search the web using Kimi API's built-in $web_search function"
    )
    parser.add_argument("query", help="Search query string")
    parser.add_argument(
        "--count",
        type=int,
        default=5,
        help="Number of results to request (default: 5)"
    )
    
    args = parser.parse_args()
    
    try:
        result = search_web(args.query, args.count)
        print(result)
    except FileNotFoundError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error performing search: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
