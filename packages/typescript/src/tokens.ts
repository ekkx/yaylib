// PORTING.md §4: Tokens is the immutable snapshot of OAuth credentials.
// Mutate by going through SetTokens, LoginWithEmail, LoadSession, or the
// internal 401 refresh path — never by mutating a returned snapshot.

export interface Tokens {
  readonly access: string;
  readonly refresh: string;
}

export function emptyTokens(): Tokens {
  return { access: "", refresh: "" };
}
