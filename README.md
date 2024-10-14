# tiny-errors with stacktrace

## 仕様

- 標準ライブラリの error 互換

  - https://github.com/golang/go/tree/go1.23.2/src/errors

- 満たすべきインターフェイス

```
type error interface {
	Error() string
}
```

- 標準 error にあるメソッドは実装

  - errors.Is、 errors.As を実装
  - errors.Join を実装

- スタックトレースを利用可能にする

## Test

`make test`

## 参考

- https://github.com/pkg/errors
- https://github.com/golang/xerrors
