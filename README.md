# 概要
goとreactを使ったアプリケーションの作成です。
goの学習と共にclean architectureとDDDについても考えなら実装していきます。

## 使用技術
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
