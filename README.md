# goでのテスト

* テスト<br>
go test

* テストの詳細<br>
go test -v

* 関数名ごとにテスト<br>
go test -run 関数名

* 関数名ごとに詳細テスト<br>
go test -v -run 関数名

* ベンチマークの結果<br>
go test -bench . -benchmem

* カバレッジ(テストの網羅率)を計測
go test -cover