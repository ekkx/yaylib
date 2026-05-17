// The pagination cursor is a string in the schema, but the server
// sends it as a bare JSON number on some endpoints and a JSON string
// on others. FromJSON must accept both and yield the string form.

import { PostsResponseFromJSON } from "../src/gen/models/PostsResponse";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

const num = PostsResponseFromJSON({ next_page_value: 585851614, posts: [] });
assert(
  "number wire -> string",
  num.nextPageValue === "585851614",
  `got ${JSON.stringify(num.nextPageValue)}`,
);

const str = PostsResponseFromJSON({ next_page_value: "abc123", posts: [] });
assert(
  "string wire -> string",
  str.nextPageValue === "abc123",
  `got ${JSON.stringify(str.nextPageValue)}`,
);

const nul = PostsResponseFromJSON({ next_page_value: null, posts: [] });
assert(
  "null wire -> undefined",
  nul.nextPageValue === undefined,
  `got ${JSON.stringify(nul.nextPageValue)}`,
);

process.exit(failed === 0 ? 0 : 1);
