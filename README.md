## Alfred
- It is a basic HTTP Multiplexer.
- http.Handler interface is implemented with Custom Server struct.
- **Noticeable code** is in **foundation/server/server.go** file. (160 lines only.)
- Server Struct
  - has a context to hold track id for requests.
  - wraps handlers with provided middlewares.
  - 

## Cloc : 
Note : Vendor directory is excluded.

Language|files|blank|comment|code
:-------|-------:|-------:|-------:|-------:
Go|6|83|0|335
XML|3|0|0|72
make|1|3|0|6
--------|--------|--------|--------|--------
SUM:|10|86|0|413

cloc|github.com/AlDanial/cloc v 1.96  T=0.01 s (830.6 files/s, 41445.5 lines/s)
--- | ---
