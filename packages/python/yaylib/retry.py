# PORTING.md §13: retry policy for transient failures (429, 5xx,
# transport errors). The wire contract:
#
#   - 429 retries on every method — the server explicitly asked.
#   - 5xx / transport errors retry only on safe methods (GET / HEAD /
#     PUT / DELETE / OPTIONS) by default. POST / PATCH retry only when
#     ``retry_on_post`` is True because the Yay! server has no
#     idempotency-key concept.
#   - The response body's ``retry_in`` field, when present, is honored
#     over the computed exponential backoff (seconds).
#   - Backoff is exponential with full jitter, clamped to ``max_delay``.
#   - The zero value (``RetryPolicy(max_attempts=0)``) disables retry.
#
# Delays are expressed in SECONDS (Python idiom — Go uses
# time.Duration, TS uses milliseconds). ``ReconnectPolicy`` (the
# event-stream sibling) has the OPPOSITE zero semantics:
# ``max_attempts == 0`` there means "unlimited".

from __future__ import annotations

import random as _random
from dataclasses import dataclass

_SAFE_METHODS = frozenset({"GET", "HEAD", "PUT", "DELETE", "OPTIONS"})


@dataclass(frozen=True)
class RetryPolicy:
    # Total attempt cap (initial attempt + retries). 0 or 1 disables
    # retry. Default: 3.
    max_attempts: int = 3
    # Base delay (seconds) before the first retry. Default: 0.2s.
    base_delay: float = 0.2
    # Hard ceiling (seconds) on any single sleep. Default: 30s.
    max_delay: float = 30.0
    # Allow retrying POST / PATCH on 5xx / transport errors. Default:
    # False. 429 retries on every method regardless.
    retry_on_post: bool = False


DEFAULT_RETRY_POLICY = RetryPolicy(
    max_attempts=3, base_delay=0.2, max_delay=30.0, retry_on_post=False
)


def method_allows_retry(method: str, policy: RetryPolicy) -> bool:
    """True if ``method`` may retry on 5xx / transport errors under the
    policy. 429 is handled separately (it retries on every method).
    """
    m = method.upper()
    if m in _SAFE_METHODS:
        return True
    if m in ("POST", "PATCH"):
        return policy.retry_on_post
    return False


def should_retry_status(status: int, method: str, policy: RetryPolicy) -> bool:
    if status == 429:
        return True
    if 500 <= status <= 599:
        return method_allows_retry(method, policy)
    return False


def full_jitter_delay(
    attempt: int, policy: RetryPolicy, rand=_random.random
) -> float:
    """Random duration in ``[0, min(max_delay, base_delay * 2^(attempt-1))]``.

    ``attempt`` is the upcoming retry's 1-indexed number (1 = first retry).
    """
    exp = min(
        policy.max_delay,
        policy.base_delay * (2 ** max(0, attempt - 1)),
    )
    return rand() * exp
