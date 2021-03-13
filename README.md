
# 淫夢で分かる迫真Vue.js部
初投稿です、オッスお願いしまーす！
これを見てるホモにVue.jsをぶち込んでやるぜー！
:::info
　これはホモのみしか受け付けていません！
 あと、素人が書いてます！
 間違ってたら許してください、何でもしますから！（何でもするとは言っていない）
:::
<br>

## 一覧

[TOC]

<br>

## Vue.js?　なんすか、ビュージェーエスって？
1. Evan You(中国人)によって作られたJavaScriptフレームワーク(jsがちょー凄くなったやつ)
2. ↑アジアで大人気、違うフレームワークのReactとかAngularは米国で人気らしい
3. 正直ReactでもVue.jsでもフロントエンドは開発出来る、メリット,デメリットはあるけど、好きな方でいいと思う。
4. HTML,CSSは基礎知ってると良い、生jsも基礎をやっていた方が分かりやすいし、それを生かした物なので、まずは生jsとか**基礎からやろう**　いきなりvue.jsはヤバい
6. しかし、生jsと書き方が違うし、コードを書く勝手が違う、概念も違うので**全く新しい言語を覚える感じ**で挑む
<br>
## 知っておきたい事
vue.jsを使うには、
1. html,css,jsを自前で用意して、外部からhtmlのscriptタグにvue.jsのCDNを読み込ませる(jQueryみたいな)やり方
2. vue-cliというターミナルを使ってコマンドを入力し、必要なファイルを構築して貰って書くやり方

の２種類が存在する。

この２種類は出来る事の差があるし、**書き方が違う。**　ネットで調べて、どっちを使って説明しているのかごっちゃになっているので、注意が必要(これで１回つまづいた)
今回は実際に使われるvue-cliを使ったやり方で説明する。
<br>
## 環境構築(コードを書く準備)
HomeBrew,Nodebrew(node.js),VueCLIをインストール(下記URLを参照)
https://zenn.dev/naitoki/articles/0470c2427de5e54ccc76
※Node.jsをそのまま入れずにNodebrewを入れる理由は、バージョン管理がしやすいから。

yarnのインストール
https://qiita.com/akitxxx/items/c97ff951ca31298f3f24

任意:VSCodeの拡張機能として,Material Icon Theme,Atom One Dark,Vue.js Extension Pack,zenkaku　がオススメ。

Windows10はPowerShell使うんじゃなくて、WSL2(windowsでubuntuとかlinuxが使えるやつ)でやった方がいい。[（スクリプト云々の問題）](https://qiita.com/Targityen/items/3d2e0b5b0b7b04963750)

(Dockerで環境を汚さないでやる方法もあるが、hihumikanはやり方を知らないので説明しない)

これらのインストールし終えたら、コマンドを打って環境構築(コードを書く準備)をして貰う。
自分の適当な場所（cdして好きなディレクトリ）で
```
$ vue create 好きな名前
```
そうすると、色々環境構築する上で色々質問される
```
Vue CLI v4.5.11
? Please pick a preset:
  Default ([Vue 2] babel, eslint)
  Default (Vue 3 Preview) ([Vue 3] babel, eslint)
  Manually select features
```
こだわりなければDefault(vue2)でいい、ホモの場合はManuallyを選択、
```
? Check the features needed for your project:
 ◉ Choose Vue version
 ◉ Babel
 ◯ TypeScript
 ◯ Progressive Web App (PWA) Support
 ◯ Router
 ◯ Vuex
 ◯ CSS Pre-processors
 ◉ Linter / Formatter
 ◯ Unit Testing
 ◯ E2E Testing

```
と表示されているが、VuexとかRouterとかは使う時にこことかで選択する。この要素は後から身につける物なので知らなくてもいい。後からも追加可能だし。
ホモはESLintが嫌いなので、ここでLinter / Formatterのインストールを除外したりしてる。

後はEnter連打でインストールが始まる。

インストールが終わると
```
success Saved lockfile.
✨  Done in 5.68s.
⚓  Running completion hooks...

📄  Generating README.md...

🎉  Successfully created project aa.
👉  Get started with the following commands:

 $ cd aa
 $ yarn serve
```
と表示される。これで必要なファイルの準備が出来た。

```
 $ cd aa
 $ yarn serve
```
をすると、
```
 DONE  Compiled successfully in 7223ms                                                                                                                                           20:29:10


  App running at:
  - Local:   http://localhost:8081/
  - Network: http://192.168.0.13:8081/

  Note that the development build is not optimized.
  To create a production build, run yarn build.
```
と表示される。

http://localhost:8081/ にアクセス(別端末は http://192.168.0.13:8081/ )すると、
![](https://i.imgur.com/YhYkort.png)
が表示される。

これの何が良いかというと、文字とか変えてすぐに変更が見たい時に`yarn serve`してブラウザを立ち上げておくだけで、ブラウザで変更がすぐに確認できるから。開発の時は`yarn serve`しておこう。

因みに、毎回`yarn serve`してターミナルをずっと立ち上げておかなければならないのと、終了させる時はターミナルをそのまま閉じるか、control+cで出来る。

また続きは書きます。

###### tags: `Templates` `Documentation`
