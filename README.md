# 概要
goとreactを使ったアプリケーションの作成です。
goの学習と共にclean architectureとDDDについても考えなら実装していきます。

## 考え中ポイント
* clean architecture的に行くとdbの関心ごとをusecaseに持ち込めないけどどうやってトランザクションを貼れば良いのか？
  *  beginとかrollbackのインターフェースを生やす方法
  *  もしくは全部をラップしていい感じに実行してくれるものを作る方法
* DIコンテナは何を使うのが良さそうか検討がついていない。
* ORMに依存しないようにするにはgatewaysがmodelに依存しない形でDTOでdriverに渡す必要があるがそこまでやるのか？
  * 別のORMでも遊びたいから余裕が出たタイミングでDTOに変更してみる。  