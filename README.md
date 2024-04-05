# 概要
goとreactを使ったアプリケーションの作成です。
## 目的
* goの習得
  * effective goなどを読みつつgoの言語思想的なものを理解する。
    * できれば他の言語との比較をして言語化する
* clearn architectureの理解を深くする
  * レイヤーの役割と依存関係の反転のさせ方(DIP)はわかるがアプリケーションの作成レベルで起きる問題を整理したい。
    * golangの言語仕様の問題で発生するもの
    * ライブラリの問題で発生するもの
* DDDの学習
  *  CQRSとかは体験したことがあるが集約単位でリポジトリに入れてみるとかやったことのないものを試してみたい。
* ゆくゆくはマイクロサービス化して境界を考える(一旦保留)

## 使用技術
* go (1.22.0)
  * echo(4.11.4),gorm(1.25.8)
* ローカル環境
  * k3s,skaffold,helm
    * k8sの勉強兼ねてdocker composeではなく、k3sでクラスターを作りskaffoldでk8sリソースを構築しています。
* DIコンテナ
  * 考え中

## 考え中ポイント
* clean architecture的に行くとdbの関心ごとをusecaseに持ち込めないけどどうやってトランザクションを貼れば良いのか？
  *  beginとかrollbackのインターフェースを生やす方法
  *  もしくは全部をラップしていい感じに実行してくれるものを作る方法
* ORMに依存しないようにするにはgatewaysがmodelに依存しない形でDTOでdriverに渡す必要があるがそこまでやるのか？
  * 別のORMでも遊びたいから余裕が出たタイミングでDTOに変更してみる。  
