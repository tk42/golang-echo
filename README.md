# golang-echo
This sample repository is an echo server with Golang that requests/responds gRPC (protoBuf) via Redis.

## Survey
GolangでgRPCアプリケーションを開発するにあたって必要そうな知識を整理する
### gRPC
 - gRPCはGoogle謹製のHTTP/2を利用したRPC(Remote Produce Call)フレームワーク
 - データのシリアライズにはProtocol Buffersを利用する(Protocol Buffers以外も利用可能だがデファクトスタンダード)
 - RPCとしては他にREST-APIなどがあるが，REST-APIのようにパスやメソッドなどを指定する必要がなく注目されている

![gRPC概念図](https://storage.googleapis.com/apps-gcp-tokyo-02.appspot.com/2015/09/grpc_concept_diagram_00.png)
gRPCの全体像([http://www.grpc.io/docs/](http://www.grpc.io/docs/)より転載)

### Protocol Buffers
 - Google製のデータをバイナリ形式にシリアライズするためのフォーマット．JSONの代替を目指しているそう
 - クライアント・サーバー間の通信やデータの永続化などに用いる
 - ```*.proto```ファイルに IDL(Interactive Data Language) でデータ構造を記述する
 - 様々な言語で対応する(Golang/Java/Python etc...)
 - シリアライズ・デシリアライズするためのコードが自動生成される(```protoc```コマンドでコンパイルする)
 - レポジトリは [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf).
 最新バージョンは 3 (proto3)

![Protocol Buffers利用フロー図](https://storage.googleapis.com/apps-gcp-tokyo-02.appspot.com/2015/09/gRPC.png)
Protocol Buffers利用フロー([https://www.apps-gcp.com/grpc-go/](https://www.apps-gcp.com/grpc-go/)より転載)

### Go Protocol Buffer Message API
 - Go Protocol Buffer Message APIとはGolang製アプリが Protocol Bufferを利用するためのもの．最近，仕様にAPIの刷新があったことと関連レポジトリが多いので整理
 - 2020/3/20 APIv2 がGo公式ブログで発表．これによりGoでもReflectionが使えるようになった．
 - APIv2 は後方互換性はないが，APIv1 は今後もずっとサポートを継続するそう
 - Protocol Buffer側のレポジトリは golang/protobuf から [protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go) に移行．Protocol Buffersのメッセージやシリアライズに```protoc-gen-go```コマンドを利用
 - gPRC側のレポジトリの [grpc/grpc-go](https://github.com/grpc/grpc-go) に新たに```protoc-gen-go-grpc```コマンドができた．gPRCのサーバ/クライアントに用いるインターフェース(スタブコード)を生成ためにこのコマンドを利用
 - gRPCを利用したアプリの作成には上記の両方のコマンドが必要

**対応表**

|対象|APIv1|APIv2|
|:--|:--|:--|
|import のパス|github.com/golang/protobuf|google.golang.org/protobuf|
|ソースのリポジトリ|https://github.com/golang/protobuf|https://github.com/protocolbuffers/protobuf-go|
|protoc-gen-go|github.com/golang/protobuf/protoc-gen-go|google.golang.org/protobuf/cmd/protoc-gen-go|
|gRPCスタブコード生成|同上|google.golang.org/grpc/cmd/protoc-gen-go-grpc|

[Protocol Buffers用 Go言語APIの APIv1 と APIv2 の差異](https://qiita.com/kitauji/items/bab05cc8215abe8a6431)より転載

### Third Party Library
 - [gogo/protobuf](https://github.com/gogo/protobuf) は公式の golang/protobuf を補完する非公式のライブラリ．2021/02現在APIv1にのみ対応．(Reflectionに対応していたがAPIv2で公式に対応されたので利用するメリットはない？)
 - [grpc-ecosystem/go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) はinterceptorを集めたライブラリ．リクエストパラメータに依存しない共通処理(具体的には認証・ログ・バリデーション・監視など）をintercept(傍受，途中で捕える)できる機構がgRPCのinterceptor
 - [gRPC command line tool](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) はコマンドラインでgRPCを呼び出すためのコマンド(gRPCはcurl等で手軽に動作確認できないため)
 - [gRPCurl](https://github.com/fullstorydev/grpcurl) は同上のCLIツール．4.5kスター(2021/02時点)．他にもgRPCのクライアントCLIツールはたくさんある模様. [gRPC の CUI のクライアントツールを調べてまとめてみた](https://www.techscore.com/blog/2019/09/26/grpc-cui-client-tools/)

### Reference
 - [gRPCって何？](https://qiita.com/oohira/items/63b5ccb2bf1a913659d6)
 - [いまさらだけどgRPCに入門したので分かりやすくまとめてみた](https://qiita.com/gold-kou/items/a1cc2be6045723e242eb)
 - [Protocol Buffers: ざっくりとした入門](https://qiita.com/kitauji/items/fdbd052c19dad28ab067)
 - [Protocol Buffers用 Go言語APIの APIv1 と APIv2 の差異](https://qiita.com/kitauji/items/bab05cc8215abe8a6431)
 - [Go言語でのgRPCコード生成(2020年10月以降版)](https://note.com/dd_techblog/n/nb8b925d21118)
 - [GoGo Protobufのメリット・デメリット](https://christina04.hatenablog.com/entry/gogo-protobuf-merit-demerit)
 - [go-grpc-middlewareを一通り試してみる](https://qiita.com/Morix1500/items/7a20d76a931af68d860d)
 - [Go + gRPCで開発を始められるDockerfileを作る](https://qiita.com/keitakn/items/434091ff488296951ab6)

## How to use
